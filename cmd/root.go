package cmd

import (
	"fmt"

	"github.com/schaermu/gopress/conf"
	"github.com/spf13/cobra"
)

var rootCmd = createRootCommand()

func createRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gopress",
		Short: "Gopress helps you building impressive, offline-capable presentations using Markdown and impress.js.",
		Long:  `Gopress will enable you to build exciting and modern presentations using impress.js by doing the thing that feels the most natural to us developers: coding.`,
	}

	// OnInitialize appends the passed function to initializers to be run when
	// each command's `Execute` method was called after `init`.
	cobra.OnInitialize(func() {
		// Load user conf file if exists.
		loadConfig(&ConfApp, &ConfUser)
	})

	cmd.PersistentFlags().StringVarP(
		&ConfApp.PathFileConf,
		"config",
		"c",
		"",
		"config file (default is $HOME/.gopress.yaml)",
	)

	return cmd
}

func init() {}

// loadConfig sets the object in the arg with the results exits with an error if user defined conf file didn't exist.
// Otherwise searches the default file and if not found then use the default value.
func loadConfig(configApp *conf.TConfigFile, configUser interface{}) {
	// Overwrite "configUser" with conf file value if file found.
	if err := conf.LoadFile(*configApp, &configUser); err != nil {
		// Exits if user defined conf file fails to read
		if configApp.PathFileConf != "" {
			msg := fmt.Errorf("failed to read configuration file.\n  Error msg: %v", err)
			osExit(EchoStdErrIfError(msg))
		}
		// Conf file not found. Using default. Set flag to true.
		configApp.IsUsingDefaultConf = true
	}
}
