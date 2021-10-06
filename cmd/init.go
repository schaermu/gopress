package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var force bool
var template string

var configFmt = `---
template: %s
`

var slidesFmt = `# First slide for %[1]v
You can now start to create your content.

# Second slide for %[1]v
`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [name]",
	Short: "Create a new presentation based on a template (or the default one).",
	Long:  `Using this command, you can bootstrap a new presentation either using the default template or a configured one.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   execCommand,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&template, "template", "t", "default", "Template to bootstrap the presentation.")
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force re-creation of config file")
}

func execCommand(cmd *cobra.Command, args []string) {
	target := args[0]
	cfgFile := fmt.Sprintf("%s/.gopress.yaml", target)
	slides := fmt.Sprintf("%s/slides.md", target)

	cDefault("[ - ] Creating skeleton inside folder %s...\n", target)
	if _, err := os.Stat(target); err != nil {
		if err := os.MkdirAll(fmt.Sprintf("%s/images", target), 0755); err != nil {
			cError("[ X ] Could not create folder (%v)\n", err)
			os.Exit(1)
		}
	}

	if _, err := os.Stat(cfgFile); err != nil || force {
		f, err := os.Create(cfgFile)
		if err != nil {
			cError("[ X ] Cound not create config (%v)\n", err)
			os.Exit(1)
		}
		defer f.Close()

		if _, err := fmt.Fprintf(f, configFmt, template); err != nil {
			cError("[ X ] Cound not write config (%v)\n", err)
			os.Exit(1)
		}
		cSuccess("[ + ] Generated config file %s.\n", cfgFile)
	} else {
		cDefault("[ - ] Config already existing, skipping init (use -f to overwrite)...\n")
	}

	if _, err := os.Stat(slides); err != nil {
		f, err := os.Create(slides)
		if err != nil {
			cError("[ X ] Cound not write markdown file for slides (%v)\n", err)
			os.Exit(1)
		}
		defer f.Close()

		if _, err := fmt.Fprintf(f, slidesFmt, target); err != nil {
			cError("[ X ] Cound not write markdown file for slides (%v)\n", err)
			os.Exit(1)
		}
		cSuccess("[ + ] Generated slides file %s.\n", slides)
	} else {
		cDefault("[ - ] Slides already existing, skipping init...\n")
	}
}
