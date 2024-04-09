/*
 * Copyright (C) 2024 by Jason Figge
 */

package main

import (
	"fmt"
	"os"
	"runtime"

	"us.figge.monitor/internal"
)

// Version information, populated by the build process
var (
	Version     string // this variable is defined in Makefile
	Commit      string // this variable is defined in Makefile
	Branch      string // this variable is defined in Makefile
	BuildNumber string //nolint:revive // this variable is defined in Makefile
)

func main() {
	app, ok := internal.NewApplication()
	if !ok {
		return
	}
	config := app.Configuration()

	if config.HelpFlag() {
		help()
	} else if config.VersionFlag() {
		version(config.VerboseFlag())
	} else {
		app.Start()
	}
}

func help() {
	fmt.Printf("6502 monitor\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("  -c, --config      Specify the monitor configuration file\n")
	fmt.Printf("  -h, --help        Display this message.\n")
	fmt.Printf("  -v, --verbose     Verbose mode.  Prints progress debug messages.\n")
	fmt.Printf("  -V, --version     Display version information.\n")
}

func version(verbose bool) {
	if verbose {
		fmt.Printf(
			"%s verison %s %s/%s, build %s, commit %s, branch %s\n",
			os.Args[0], Version, runtime.GOOS, runtime.GOARCH, BuildNumber, Commit, Branch,
		)
	} else {
		fmt.Printf(
			"%s verison %s %s/%s, build %s, commit %s\n",
			os.Args[0], Version, runtime.GOOS, runtime.GOARCH, BuildNumber, Commit,
		)
	}
}
