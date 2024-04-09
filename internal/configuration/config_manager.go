/*
 * Copyright (C) 2024 by Jason Figge
 */

package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	GoosLinux   = "linux"
	GoosDarwin  = "darwin"
	GoosWindows = "windows"
)

type ConfigManager struct {
	configFile string
	help       bool
	verbose    bool
	version    bool
	config     *Configuration
}

type Configuration struct{}

func NewConfigManager() (*ConfigManager, bool) {
	cm := ConfigManager{}
	cm.setDefaults()
	cm.parseCommandLine()
	cm.load()
	return &cm, true
}

func (cm *ConfigManager) setDefaults() {
	cm.configFile = "config.yaml"
	currentUser, err := user.Current()
	if err != nil {
		return
	}
	switch runtime.GOOS {
	case GoosLinux:
		cm.configFile = fmt.Sprintf("/home/%s/.monitor/config.yaml", currentUser.Username)
	case GoosDarwin:
		cm.configFile = fmt.Sprintf("/Users/%s/.monitor/config.yaml", currentUser.Username)
	case GoosWindows:
		cm.configFile = fmt.Sprintf("C:\\Users\\%s\\.monitor\\config.yaml", currentUser.Username)
	}
}

func (cm *ConfigManager) parseCommandLine() {
	success := true
	for index := 1; index < len(os.Args); index++ {
		switch os.Args[index] {
		case "-h", "--help":
			cm.help = true
		case "-V", "--version":
			cm.version = true
		case "-v", "--verbose":
			cm.verbose = true
		case "-c", "--config":
			index++
			cm.configFile, success = parameter(index)

		default:
			if strings.HasPrefix(os.Args[index], "-") {
				fmt.Printf("  Error - unknown paramters (%s) at position %d\n", os.Args[index], index)
			} else {
				fmt.Printf("  Error - unexpected argument (%s) as position %d\n", os.Args[index], index)
			}
			success = false
		}
		if !success {
			cm.help = true
		}
	}
}

func parameter(index int) (string, bool) {
	if index < len(os.Args) && !strings.HasPrefix(os.Args[index], "-") {
		return os.Args[index], true
	}
	fmt.Printf("  Error - paramreter %s requires a value\n", os.Args[index-1])
	return "", false
}

func (cm *ConfigManager) load() bool {
	if fi, err := os.Stat(cm.configFile); os.IsNotExist(err) {
		fmt.Printf("config file (%s) cannot be read: file not found\n", cm.configFile)
		return false
	} else if fi.IsDir() {
		fmt.Printf("config file (%s) cannot be read: file is a directory\n", cm.configFile)
		return false
	}
	bs, err := os.ReadFile(cm.configFile)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Printf("config file (%s) cannot be read: permission denied\n", cm.configFile)
		} else {
			fmt.Printf("config file (%s) cannot be read: %v\n", cm.configFile, err)
		}
		return false
	}

	cm.config = &Configuration{}
	if strings.HasSuffix(cm.configFile, "yaml") || strings.HasSuffix(cm.configFile, "yml") {
		err = yaml.Unmarshal(bs, cm.config)
	} else if strings.HasSuffix(cm.configFile, "json") {
		err = json.Unmarshal(bs, cm.config)
	} else {
		fmt.Printf("config file (%s) has unknown extension\n", cm.configFile)
		return false
	}
	if err != nil {
		fmt.Printf("config file (%s) cannot be parsed: %v\n", cm.configFile, err)
		return false
	}
	return true

}

func (cm *ConfigManager) HelpFlag() bool {
	return cm.help
}

func (cm *ConfigManager) VerboseFlag() bool {
	return cm.verbose
}

func (cm *ConfigManager) VersionFlag() bool {
	return cm.version
}
