/*

Copyright (c) 2018 - 2022 PhotoPrism UG. All rights reserved.

    This program is free software: you can redistribute it and/or modify
    it under Version 3 of the GNU Affero General Public License (the "AGPL"):
    <https://docs.photoprism.app/license/agpl>

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    The AGPL is supplemented by our Trademark and Brand Guidelines,
    which describe how our Brand Assets may be used:
    <https://photoprism.app/trademark>

Feel free to send an email to hello@photoprism.app if you have questions,
want to support our work, or just want to say hello.

Additional information can be found in our Developer Guide:
<https://docs.photoprism.app/developer-guide/>

*/
package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/commands"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/event"
)

var version = "development"
var log = event.Log

const appName = "PhotoPrism"
const appEdition = "PhotoPrism® CE"
const appDescription = "AI-Powered Photos App. Visit https://docs.photoprism.app/ to learn more."
const appCopyright = "(c) 2018-2022 PhotoPrism UG. All rights reserved."

// Metadata contains build specific information.
var Metadata = map[string]interface{}{
	"Name":        appName,
	"Edition":     appEdition,
	"Description": appDescription,
	"Version":     version,
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()

	app := cli.NewApp()
	app.Usage = appEdition
	app.Description = appDescription
	app.Version = version
	app.Copyright = appCopyright
	app.EnableBashCompletion = true
	app.Flags = config.Flags.Cli()
	app.Commands = commands.PhotoPrism
	app.Metadata = Metadata

	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
}
