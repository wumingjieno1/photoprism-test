/*

Package auto provides workers for background indexing and import operations.

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
package auto

import (
	"time"

	"github.com/photoprism/photoprism/internal/config"

	"github.com/photoprism/photoprism/internal/event"
)

var log = event.Log

var stop = make(chan bool, 1)

// Wait starts waiting for indexing & importing opportunities.
func Start(conf *config.Config) {
	// Don't start ticker if both are disabled.
	if conf.AutoIndex().Seconds() <= 0 && conf.AutoImport().Seconds() <= 0 {
		return
	}

	ticker := time.NewTicker(time.Minute)

	go func() {
		for {
			select {
			case <-stop:
				ticker.Stop()
				return
			case <-ticker.C:
				if mustIndex(conf.AutoIndex()) {
					log.Debugf("auto-index: starting")
					ResetIndex()
					if err := Index(); err != nil {
						log.Errorf("auto-index: %s", err)
					}
				} else if mustImport(conf.AutoImport()) {
					log.Debugf("auto-import: starting")
					ResetImport()
					if err := Import(); err != nil {
						log.Errorf("auto-import: %s", err)
					}
				}
			}
		}
	}()
}

// Stop stops waiting for indexing & importing opportunities.
func Stop() {
	stop <- true
}
