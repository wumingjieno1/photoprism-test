/*

Package commands provides photoprism CLI (sub-)commands.

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
package commands

import (
	"os"
	"syscall"

	"github.com/sevlyar/go-daemon"
	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/pkg/fs"
)

var log = event.Log

// PhotoPrism contains the photoprism CLI (sub-)commands.
var PhotoPrism = []cli.Command{
	StartCommand,
	StopCommand,
	StatusCommand,
	IndexCommand,
	ImportCommand,
	CopyCommand,
	FacesCommand,
	PlacesCommand,
	PurgeCommand,
	CleanUpCommand,
	OptimizeCommand,
	MomentsCommand,
	ConvertCommand,
	ThumbsCommand,
	MigrationsCommand,
	BackupCommand,
	RestoreCommand,
	ResetCommand,
	PasswdCommand,
	UsersCommand,
	ShowCommand,
	VersionCommand,
	ShowConfigCommand,
}

// childAlreadyRunning tests if a .pid file at filePath is a running process.
// it returns the pid value and the running status (true or false).
func childAlreadyRunning(filePath string) (pid int, running bool) {
	if !fs.FileExists(filePath) {
		return pid, false
	}

	pid, err := daemon.ReadPidFile(filePath)
	if err != nil {
		return pid, false
	}

	process, err := os.FindProcess(int(pid))
	if err != nil {
		return pid, false
	}

	return pid, process.Signal(syscall.Signal(0)) == nil
}
