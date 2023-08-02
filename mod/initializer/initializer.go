//
// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//

package initializer

import (
	"configurator"
	"vars"
)

// initialize the system/variable/template
func Init(c *configurator.Config) {

	// set to default if not set in the configuration file
	if c.LogValues.LogsDir == "" {
		c.LogValues.LogsDir = vars.LogsDir
	}
	if c.LogValues.LogFile == "" {
		c.LogValues.LogFile = vars.LogFile
	}
	if c.LogValues.LogMaxSize == 0 {
		c.LogValues.LogMaxSize = vars.LogMaxSize
	}
	if c.LogValues.LogMaxBackups == 0 {
		c.LogValues.LogMaxBackups = vars.LogMaxBackups
	}
	if c.LogValues.LogMaxAge == 0 {
		c.LogValues.LogMaxAge = vars.LogMaxAge
	}

    // see if emoji was give on the commandline
    if len(c.MsgEmoji) != 0 {
        c.TeamsValues.MsgEmoji = c.MsgEmoji
    }
}
