package commands

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/pkg/report"
)

// ShowOptionsCommand configures the command name, flags, and action.
var ShowOptionsCommand = cli.Command{
	Name:    "options",
	Aliases: []string{"flags"},
	Usage:   "Displays supported config flags and variable names",
	Flags:   report.CliFlags,
	Action:  showOptionsAction,
}

var faceOptionsInfo = `!!! info ""    
    To [recognize faces](https://docs.photoprism.app/user-guide/organize/people/), PhotoPrism first extracts crops from your images using a
    [library](https://github.com/esimov/pigo) based on [pixel intensity comparisons](https://arxiv.org/pdf/1305.4537.pdf).
    These are then fed into TensorFlow to compute [512-dimensional vectors](https://www.cv-foundation.org/openaccess/content_cvpr_2015/papers/Schroff_FaceNet_A_Unified_2015_CVPR_paper.pdf)
    for characterization. In the final step, the [DBSCAN algorithm](https://en.wikipedia.org/wiki/DBSCAN)
    attempts to cluster these so-called face embeddings, so they can be matched to persons with just a few clicks.
    A reasonable range for the similarity distance between face embeddings is between 0.60 and 0.70, with a higher
    value being more aggressive and leading to larger clusters with more false positives.
    To cluster a smaller number of faces, you can reduce the core to 3 or 2 similar faces.

We recommend that only advanced users change these parameters:`

// showOptionsAction shows environment variable command-line parameter names.
func showOptionsAction(ctx *cli.Context) error {
	conf := config.NewConfig(ctx)
	conf.SetLogLevel(logrus.FatalLevel)

	rows, cols := config.Flags.Report()

	// CSV Export?
	if ctx.Bool("csv") || ctx.Bool("tsv") {
		result, err := report.Render(rows, cols, report.CliFormat(ctx))

		fmt.Println(result)

		return err
	}

	type Section struct {
		Start   string
		Caption string
		Info    string
	}

	s := []Section{
		{Start: "PHOTOPRISM_ADMIN_PASSWORD", Caption: "Authentication"},
		{Start: "PHOTOPRISM_LOG_LEVEL", Caption: "Logging"},
		{Start: "PHOTOPRISM_CONFIG_PATH", Caption: "Storage"},
		{Start: "PHOTOPRISM_WORKERS", Caption: "Index Workers"},
		{Start: "PHOTOPRISM_READONLY", Caption: "Feature Flags"},
		{Start: "PHOTOPRISM_DEFAULT_LOCALE", Caption: "Customization"},
		{Start: "PHOTOPRISM_CDN_URL", Caption: "Site Information"},
		{Start: "PHOTOPRISM_HTTP_PORT", Caption: "Web Server"},
		{Start: "PHOTOPRISM_DATABASE_DRIVER", Caption: "Database Connection"},
		{Start: "PHOTOPRISM_DARKTABLE_BIN", Caption: "File Converters"},
		{Start: "PHOTOPRISM_DOWNLOAD_TOKEN", Caption: "Security Tokens"},
		{Start: "PHOTOPRISM_THUMB_COLOR", Caption: "Image Quality"},
		{Start: "PHOTOPRISM_FACE_SIZE", Caption: "Face Recognition",
			Info: faceOptionsInfo},
		{Start: "PHOTOPRISM_PID_FILENAME", Caption: "Daemon Mode",
			Info: "If you start the server as a *daemon* in the background, you can additionally specify a filename for the log and the process ID:"},
	}

	j := 0

	for i, sec := range s {
		fmt.Printf("### %s ###\n\n", sec.Caption)
		if sec.Info != "" && ctx.Bool("md") {
			fmt.Printf("%s\n\n", sec.Info)
		}

		secRows := make([][]string, 0, len(rows))

		for {
			row := rows[j]

			if len(row) < 1 {
				continue
			}

			if i < len(s)-1 {
				if s[i+1].Start == row[0] {
					break
				}
			}

			secRows = append(secRows, row)
			j++

			if j >= len(rows) {
				break
			}
		}

		result, err := report.Render(secRows, cols, report.CliFormat(ctx))

		if err != nil {
			return err
		}

		fmt.Println(result)

		if j >= len(rows) {
			break
		}
	}

	return nil
}
