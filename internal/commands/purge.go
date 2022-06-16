package commands

import (
	"context"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize/english"

	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/service"
	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
)

// PurgeCommand registers the index cli command.
var PurgeCommand = cli.Command{
	Name:   "purge",
	Usage:  "Updates missing files, photo counts, and album covers",
	Flags:  purgeFlags,
	Action: purgeAction,
}

var purgeFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "hard",
		Usage: "permanently remove from index",
	},
	cli.BoolFlag{
		Name:  "dry",
		Usage: "dry run, don't actually remove anything",
	},
}

// purgeAction removes missing files from search results
func purgeAction(ctx *cli.Context) error {
	start := time.Now()

	conf := config.NewConfig(ctx)
	service.SetConfig(conf)

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := conf.Init(); err != nil {
		return err
	}

	conf.InitDb()

	// get cli first argument
	subPath := strings.TrimSpace(ctx.Args().First())

	if subPath == "" {
		log.Infof("purge: removing missing files in %s", clean.Log(filepath.Base(conf.OriginalsPath())))
	} else {
		log.Infof("purge: removing missing files in %s", clean.Log(fs.RelName(filepath.Join(conf.OriginalsPath(), subPath), filepath.Dir(conf.OriginalsPath()))))
	}

	if conf.ReadOnly() {
		log.Infof("config: read-only mode enabled")
	}

	w := service.Purge()

	opt := photoprism.PurgeOptions{
		Path: subPath,
		Dry:  ctx.Bool("dry"),
		Hard: ctx.Bool("hard"),
	}

	if files, photos, err := w.Start(opt); err != nil {
		return err
	} else {
		log.Infof("purged %s and %s in %s", english.Plural(len(files), "file", "files"), english.Plural(len(photos), "photo", "photos"), time.Since(start))
	}

	conf.Shutdown()

	return nil
}
