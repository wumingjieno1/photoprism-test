package photoprism

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/ffmpeg"

	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
)

// ToAvc converts a single video file to MPEG-4 AVC.
func (c *Convert) ToAvc(f *MediaFile, encoder ffmpeg.AvcEncoder, noMutex, force bool) (file *MediaFile, err error) {
	if f == nil {
		return nil, fmt.Errorf("convert: file is nil - possible bug")
	}

	if !f.Exists() {
		return nil, fmt.Errorf("convert: %s not found", f.RootRelName())
	}

	avcName := fs.VideoAVC.FindFirst(f.FileName(), []string{c.conf.SidecarPath(), fs.HiddenPath}, c.conf.OriginalsPath(), false)

	mediaFile, err := NewMediaFile(avcName)

	if err == nil && mediaFile.IsVideo() {
		return mediaFile, nil
	}

	if !c.conf.SidecarWritable() {
		return nil, fmt.Errorf("convert: transcoding disabled in read-only mode (%s)", f.RootRelName())
	}

	if c.conf.DisableFFmpeg() {
		return nil, fmt.Errorf("convert: ffmpeg is disabled for transcoding %s", f.RootRelName())
	}

	fileName := f.RelName(c.conf.OriginalsPath())
	avcName = fs.FileName(f.FileName(), c.conf.SidecarPath(), c.conf.OriginalsPath(), fs.ExtAVC)

	cmd, useMutex, err := c.AvcConvertCommand(f, avcName, encoder)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Make sure only one convert command runs at a time.
	if useMutex && !noMutex {
		c.cmdMutex.Lock()
		defer c.cmdMutex.Unlock()
	}

	if fs.FileExists(avcName) {
		avcFile, avcErr := NewMediaFile(avcName)
		if avcErr != nil {
			return avcFile, avcErr
		} else if !force || !avcFile.InSidecar() {
			return avcFile, nil
		} else if err = avcFile.Remove(); err != nil {
			return avcFile, fmt.Errorf("convert: failed removing %s (%s)", clean.Log(avcFile.RootRelName()), err)
		} else {
			log.Infof("convert: replacing %s", clean.Log(avcFile.RootRelName()))
		}
	}

	// Fetch command output.
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Env = []string{fmt.Sprintf("HOME=%s", c.conf.CmdCachePath())}

	event.Publish("index.converting", event.Data{
		"fileType": f.FileType(),
		"fileName": fileName,
		"baseName": filepath.Base(fileName),
		"xmpName":  "",
	})

	log.Infof("%s: transcoding %s to %s", encoder, fileName, fs.VideoAVC)

	// Log exact command for debugging in trace mode.
	log.Trace(cmd.String())

	// Run convert command.
	start := time.Now()
	if err = cmd.Run(); err != nil {
		if stderr.String() != "" {
			err = errors.New(stderr.String())
		}

		// Log ffmpeg output for debugging.
		if err.Error() != "" {
			log.Debug(err)
		}

		// Log filename and transcoding time.
		log.Warnf("%s: failed transcoding %s [%s]", encoder, fileName, time.Since(start))

		// Remove broken video file.
		if !fs.FileExists(avcName) {
			// Do nothing.
		} else if err = os.Remove(avcName); err != nil {
			return nil, fmt.Errorf("convert: failed removing %s (%s)", clean.Log(RootRelName(avcName)), err)
		}

		// Try again using software encoder.
		if encoder != ffmpeg.SoftwareEncoder {
			return c.ToAvc(f, ffmpeg.SoftwareEncoder, true, false)
		} else {
			return nil, err
		}
	}

	// Log transcoding time.
	log.Infof("%s: created %s [%s]", encoder, filepath.Base(avcName), time.Since(start))

	return NewMediaFile(avcName)
}

// AvcConvertCommand returns the command for converting video files to MPEG-4 AVC.
func (c *Convert) AvcConvertCommand(f *MediaFile, avcName string, encoder ffmpeg.AvcEncoder) (result *exec.Cmd, useMutex bool, err error) {
	fileName := f.FileName()
	bitrate := c.AvcBitrate(f)
	ffmpegBin := c.conf.FFmpegBin()

	switch {
	case fileName == "":
		return nil, false, fmt.Errorf("convert: %s video filename is empty - possible bug", f.FileType())
	case bitrate == "":
		return nil, false, fmt.Errorf("convert: transcoding bitrate is empty - possible bug")
	case ffmpegBin == "":
		return nil, false, fmt.Errorf("convert: ffmpeg must be installed to transcode %s to avc", clean.Log(f.BaseName()))
	case !f.IsAnimated():
		return nil, false, fmt.Errorf("convert: file type %s of %s cannot be transcoded to avc", f.FileType(), clean.Log(f.BaseName()))
	}

	return ffmpeg.AvcConvertCommand(fileName, avcName, ffmpegBin, c.AvcBitrate(f), encoder)
}

// AvcBitrate returns the ideal AVC encoding bitrate in megabits per second.
func (c *Convert) AvcBitrate(f *MediaFile) string {
	const defaultBitrate = "8M"

	if f == nil {
		return defaultBitrate
	}

	limit := c.conf.FFmpegBitrate()
	quality := 12

	bitrate := int(math.Ceil(float64(f.Width()*f.Height()*quality) / 1000000))

	if bitrate <= 0 {
		return defaultBitrate
	} else if bitrate > limit {
		bitrate = limit
	}

	return fmt.Sprintf("%dM", bitrate)
}
