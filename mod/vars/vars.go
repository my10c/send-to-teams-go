//
// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//

package vars

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"
)

var (
	MyVersion   = "0.0.1"
	now         = time.Now()
	MyProgname  = path.Base(os.Args[0])
	myAuthor    = "Luc Suryo"
	myCopyright = "Copyright 2023 - " + strconv.Itoa(now.Year()) + " ©Badassops LLC"
	myLicense   = "License 3-Clause BSD, https://opensource.org/licenses/BSD-3-Clause ♥"
	myEmail     = "<luc@badassops.com>"
	MyInfo      = fmt.Sprintf("%s (version %s)\n%s\n%s\nWritten by %s %s\n",
		MyProgname, MyVersion, myCopyright, myLicense, myAuthor, myEmail)
	MyDescription = "Simple script send a message to a Microsoft Teams channel"

	// Default configuration file
	ConfigFile = "/usr/local/etc/send-to-teams/config.ini"

	// default values for logs
	LogEnable	bool = false
	// Logs		Log

	// default emoji
	Emoji = "🚨"

	// default value for lock
	Lock		bool = false
	LockFile	= "/tmp/" + MyProgname + ".lock"

	// default quiet
	Quiet		bool = false

	// we sets these under variable
	// default values
	LogsDir       = fmt.Sprintf("/var/log/%s", MyProgname)
	LogFile       = fmt.Sprintf("%s.log", MyProgname)
	LogMaxSize    = 128 // megabytes
	LogMaxBackups = 14  // 14 files
	LogMaxAge     = 14  // 14 days
)
