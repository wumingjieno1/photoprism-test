package ffmpeg

import (
	"fmt"
	"os/exec"

	"github.com/photoprism/photoprism/pkg/fs"
)

// AvcConvertCommand returns the command for converting video files to MPEG-4 AVC.
func AvcConvertCommand(fileName, avcName, ffmpegBin, bitrate string, encoder AvcEncoder) (result *exec.Cmd, useMutex bool, err error) {
	if fileName == "" {
		return nil, false, fmt.Errorf("empty input filename")
	} else if avcName == "" {
		return nil, false, fmt.Errorf("empty output filename")
	}

	// Don't transcode more than one video at the same time.
	useMutex = true

	// Animated GIF?
	if fs.FileType(fileName) == fs.ImageGIF {
		result = exec.Command(
			ffmpegBin,
			"-i", fileName,
			"-movflags", "faststart",
			"-pix_fmt", "yuv420p",
			"-vf", "scale=trunc(iw/2)*2:trunc(ih/2)*2",
			"-f", "mp4",
			"-y",
			avcName,
		)

		return result, useMutex, nil
	}

	// Display encoder info.
	if encoder != SoftwareEncoder {
		log.Infof("convert: ffmpeg encoder %s selected", string(encoder))
	}

	switch encoder {
	case IntelEncoder:
		// ffmpeg -hide_banner -h encoder=h264_qsv
		format := "format=rgb32"
		result = exec.Command(
			ffmpegBin,
			"-qsv_device", "/dev/dri/renderD128",
			"-i", fileName,
			"-c:a", "aac",
			"-vf", format,
			"-c:v", string(encoder),
			"-vsync", "vfr",
			"-r", "30",
			"-b:v", bitrate,
			"-bitrate", bitrate,
			"-f", "mp4",
			"-y",
			avcName,
		)

	case AppleEncoder:
		// ffmpeg -hide_banner -h encoder=h264_videotoolbox
		format := "format=yuv420p"
		result = exec.Command(
			ffmpegBin,
			"-i", fileName,
			"-c:v", string(encoder),
			"-c:a", "aac",
			"-vf", format,
			"-profile", "high",
			"-level", "51",
			"-vsync", "vfr",
			"-r", "30",
			"-b:v", bitrate,
			"-f", "mp4",
			"-y",
			avcName,
		)

	case NvidiaEncoder:
		// ffmpeg -hide_banner -h encoder=h264_nvenc
		result = exec.Command(
			ffmpegBin,
			"-r", "30",
			"-i", fileName,
			"-pix_fmt", "yuv420p",
			"-c:v", string(encoder),
			"-c:a", "aac",
			"-preset", "15",
			"-pixel_format", "yuv420p",
			"-gpu", "any",
			"-vf", "format=yuv420p",
			"-rc:v", "constqp",
			"-cq", "0",
			"-tune", "2",
			"-b:v", bitrate,
			"-profile:v", "1",
			"-level:v", "41",
			"-coder:v", "1",
			"-f", "mp4",
			"-y",
			avcName,
		)

	case Video4LinuxEncoder:
		// ffmpeg -hide_banner -h encoder=h264_v4l2m2m
		format := "format=yuv420p"
		result = exec.Command(
			ffmpegBin,
			"-i", fileName,
			"-c:v", string(encoder),
			"-c:a", "aac",
			"-vf", format,
			"-num_output_buffers", "72",
			"-num_capture_buffers", "64",
			"-max_muxing_queue_size", "1024",
			"-crf", "23",
			"-vsync", "vfr",
			"-r", "30",
			"-b:v", bitrate,
			"-f", "mp4",
			"-y",
			avcName,
		)

	default:
		format := "format=yuv420p"
		result = exec.Command(
			ffmpegBin,
			"-i", fileName,
			"-c:v", string(encoder),
			"-c:a", "aac",
			"-vf", format,
			"-max_muxing_queue_size", "1024",
			"-crf", "23",
			"-vsync", "vfr",
			"-r", "30",
			"-b:v", bitrate,
			"-f", "mp4",
			"-y",
			avcName,
		)
	}

	return result, useMutex, nil
}
