//
// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//

package configurator

import (
	"fmt"
	"os"

	"vars"

	"github.com/my10c/packages-go/is"
	"github.com/my10c/packages-go/print"

	"github.com/BurntSushi/toml"
	"github.com/akamensky/argparse"
)

type (

	// for logging
	LogConfig struct {
		LogsDir			string
		LogFile			string
		LogMaxSize		int
		LogMaxBackups	int
		LogMaxAge		int
		LogEnable		bool
	}

	Config struct {
		AuthValues		Auth
		TeamsValues		TeamsConfig
		LogValues		LogConfig
		// given from the command line
		TeamsMessage	[]string
		// from command line or default is used
		ConfigFile		string
		MsgEmoji		string
		LockFile		string
		LockPID			int
		Quite			bool
	}

	// Configuration for Teams
	TeamsConfig struct {
		Token		string
		User		string
		Channel		string
		UserEmoji	string
		MsgEmoji	string
		Lock		bool
		LockFile	string
	}

	Auth struct {
		AllowUsers	[]string
		AllowMods	[]string
	}

	tomlConfig struct {
		Auth		Auth			`toml:"auth"`
		Teams		TeamsConfig		`toml:"teams"`
		LogConfig	LogConfig		`toml:"logconfig"`
	}
)

// function to initialize the configuration
func Configurator() *Config {
	// the rest of the values will be filled from the given configuration file
	return &Config{}
}

func (c *Config) InitializeArgs(p *print.Print, i *is.Is) {
	parser := argparse.NewParser(vars.MyProgname, vars.MyDescription)
	configFile := parser.String("c", "configFile",
		&argparse.Options{
			Required: false,
			Help:		"Configuration file to be use",
			Default:	vars.ConfigFile,
		})

	slackMessage := parser.StringList("m", "message",
		&argparse.Options{
			Required: false,
			Help:		"Message to be sent between double quotes or single quotes, required",
		})

	slackEmoji := parser.String("e", "emoji",
		&argparse.Options{
			Required:	false,
			Help:		"Emoji to use.",
		})

	quietFlag := parser.Flag("q", "quiet",
		&argparse.Options{
			Required:	false,
			Help:		"quiet mode",
			Default:	vars.Quiet,
		})

	showVersion := parser.Flag("v", "version",
		&argparse.Options{
			Required: false,
			Help:		"Show version",
		})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *showVersion {
		p.ClearScreen()
		p.PrintYellow(vars.MyProgname + " version: " + vars.MyVersion + "\n")
		os.Exit(0)
	}

	if len(*slackMessage) == 0 {
		p.PrintRed("The flag -m/--message is required\n")
		os.Exit(1)
	} 

	if _, ok, _ := i.IsExist(*configFile, "file"); !ok {
		p.PrintRed("Configuration file " + *configFile + " does not exist\n")
		os.Exit(1)
	}

	if *quietFlag {
		c.Quite = true
	}

	c.ConfigFile	= *configFile
	c.TeamsMessage = *slackMessage
	c.MsgEmoji		= *slackEmoji
}

// function to add the values to the Config object from the configuration file
func (c *Config) InitializeConfigs(p *print.Print) {
	var configValues tomlConfig
	
	// set default value and then overwrite if exist in the configuration file
	// set to default for lock logs
	c.TeamsValues.Lock = vars.Lock
	c.TeamsValues.LockFile = vars.LockFile
	c.LogValues.LogEnable = vars.LogEnable
	c.LogValues.LogFile = vars.LogFile

	if _, err := toml.DecodeFile(c.ConfigFile, &configValues); err != nil {
		p.PrintRed("Error reading the configuration file\n")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(configValues.Teams.Token) == 0 ||
		len(configValues.Teams.User) == 0 ||
		len(configValues.Teams.Channel) == 0 {
		p.PrintRed("Error reading the configuration file, some value are missing or is empty\n")
		p.PrintBlue("Make sure token, user and channel are set\n")
		p.PrintBlue("Aborting...\n")
		os.Exit(1)
	}

	c.AuthValues = configValues.Auth
	c.LogValues = configValues.LogConfig
	c.TeamsValues = configValues.Teams
	c.LockFile = vars.LockFile
}
