// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
//

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	// "context"

	// local
	"configurator"
	"initializer"
	"logs"
	// "vars"

	// on github
	"github.com/my10c/packages-go/is"
	"github.com/my10c/packages-go/lock"
	"github.com/my10c/packages-go/print"
	"github.com/my10c/packages-go/spinner"

	// teams sdk
	// msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	// graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
)

func main() {
	LockPid := os.Getpid()
	progName, _ := os.Executable()
	progBase := filepath.Base(progName)

	i := is.New()
	p := print.New()
	s := spinner.New(10)

	// get given parameters and set the values
	config := configurator.Configurator()
	config.InitializeArgs(p, i)
	config.InitializeConfigs(p)

	// initialize the value to default if not defined
	initializer.Init(config)

	// quite mode no spinner
	if !config.Quite {
		go s.Run()
	}

	// initialize the logger system is it was set to true
	if config.LogValues.LogEnable {
		LogConfig := &logs.LogConfig{
			LogsDir:		config.LogValues.LogsDir,
			LogFile:		config.LogValues.LogFile,
			LogMaxSize:		config.LogValues.LogMaxSize,
			LogMaxBackups:	config.LogValues.LogMaxBackups,
			LogMaxAge:		config.LogValues.LogMaxAge,
		}

		logs.InitLogs(LogConfig)
		logs.Log("System all clear", "INFO")
	}

	// prevent a race
	time.Sleep(1 * time.Second)
	if !config.Quite {
		s.Stop()
	}

	if config.TeamsValues.Lock {
		// create the lock file to prevent an other script is running/started if lock was set
		lockPtr := lock.New(config.LockFile)
		// check lock file; lock file should not exist
		config.LockPID = LockPid
		if _, fileExist, _ := i.IsExist(config.LockFile, "file"); fileExist {
	 		lockPid, _ := lockPtr.LockGetPid()
			if progRunning, _ := i.IsRunning(progBase, lockPid); progRunning {
	 			p.PrintRed(fmt.Sprintf("\nError there is already a process %s running, aborting...\n", progBase))
				os.Exit(0)
			}
		}
		// save to create new or overwrite the lock file
		if err := lockPtr.LockIt(LockPid); err != nil {
			p.PrintRed(fmt.Sprintf("\nError creating the lock file, error %s, aborting..\n", err.Error()))
			os.Exit(0)
		}
	}

	// prepare the message
	TeamsMsg := fmt.Sprintf("%s %s\n",config.TeamsValues.MsgEmoji, strings.Join(config.TeamsMessage, " "))

	p.PrintBlue(fmt.Sprintf("\n %s \n", TeamsMsg))

	// send the message
	// _, _, err := TeamsAPI.PostMessage(config.TeamsValues.Channel,
	// 				Teams.MsgOptionText(TeamsMsg, false),
	// 				Teams.MsgOptionPostMessageParameters(TeamsMsgOptions),)
	// if err !=nil {
	// 		p.PrintRed(fmt.Sprintf("\nError sending the message, error %s..\n", err.Error()))
	// }

	if !config.Quite {
		p.TheEnd()
		fmt.Printf("\t%s\n", p.PrintLine(print.Purple, 50))
	}

	if config.TeamsValues.Lock {
		os.Remove(config.LockFile)
	}

	if config.LogValues.LogEnable {
		logs.Log("System Normal shutdown", "INFO")
	}
	os.Exit(0)
}
