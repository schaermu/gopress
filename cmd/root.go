package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gopress",
	Short: "Gopress helps you building impressive, offline-capable presentations using Markdown and impress.js.",
	Long: `Gopress will enable you to build exciting and modern presentations using impress.js by doing the thing that feels the most natural to us developers: coding.
Creating the content is as natural as writing markdown files, building a presentable version of your presentation is done by using a CLI.
Since everything is stored in code, you can even use version control to manage your presentations!

An example workflow looks like this:
1. Start by creating a directory for your presentations: mkdir -p my-awesome-slides && cd my-awsome-slides.
2. Create your first presentation: gopress create gopress-101.
3. Start writing your content using GitHub-Flavored markdown: echo '# GoPress\n## ...simply\n## ...works' > 00_intro.md.
4. Build your presentation and open it in your default browser: gopress present.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gopress.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gopress" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gopress")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
