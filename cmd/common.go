package cmd

import (
	"fmt"
	"os"

	"github.com/schaermu/gopress/conf"
)

const (
	// SUCCESS is an alias of exit status code to ease read.
	SUCCESS int = 0
	// FAILURE is an alias of exit status code to ease read.
	FAILURE int = 1
)

// osExit is a copy of `os.Exit` to ease the "exit status" test.
// See: https://stackoverflow.com/a/40801733/8367711
var osExit = os.Exit

// EchoStdErrIfError is an STDERR wrapper and returns 0(zero) or 1.
// It does nothing if the error is nil and returns 0.
func EchoStdErrIfError(err error) int {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		return FAILURE
	}

	return SUCCESS
}

// TConfUser defines the data structure to store values from a conf file. Viper
// will read these values from the config file or env variables. `mapstructure`
// defines the key name in the conf file.
type TConfUser struct {
	Template string `mapstructure:"template"`
}

var (
	// ConfApp is the basic app settings.
	ConfApp = conf.TConfigFile{
		PathDirConf:        ".",
		NameFileConf:       "config",
		NameTypeConf:       "yaml",
		PathFileConf:       ".gopress", // User defined file path
		IsUsingDefaultConf: false,      // Set to true if conf file not fond
	}

	// ConfUser holds the values read from the config file. The values here are
	// the default.
	ConfUser = TConfUser{
		Template: "default",
	}
)

// Execute is the main function of `cmd` package.
// It adds all the child commands to the root's command tree and sets their flag
// settings. Then runs/executes the `rootCmd` to find appropriate matches for child
// commands with corresponding flags and args.
//
// Usually `cmd.Execute` will be called by the `main.main()` and it only needs to
// happen once to the rootCmd.
// Returns `error` when it fails to execute.
func Execute() error {
	// Read conf file values to ConfUser with ConfApp settings
	return rootCmd.Execute()
}
