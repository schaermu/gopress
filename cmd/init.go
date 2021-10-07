package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type TypePropertyInit struct {
	force    bool
	template string
}

var propertyInit = &TypePropertyInit{}

var configFmt = `---
template: %s
`

var slidesFmt = `# First slide for %[1]v
You can now start to create your content.

# Second slide for %[1]v
`

func createInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [name]",
		Short: "Create a new presentation based on a template (or the default one).",
		Long:  `Using this command, you can bootstrap a new presentation either using the default template or a configured one.`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return execInitCommand(cmd, args)
		},
	}

	cmd.Flags().StringVarP(&propertyInit.template, "template", "t", "", "Template to bootstrap the presentation.")
	cmd.Flags().BoolVarP(&propertyInit.force, "force", "f", false, "Force re-creation of config file")

	return cmd
}

func init() {
	rootCmd.AddCommand(createInitCommand())
}

func execInitCommand(cmd *cobra.Command, args []string) error {
	target := args[0]
	cfgFile := fmt.Sprintf("%s/.gopress.yaml", target)
	slides := fmt.Sprintf("%s/slides.md", target)

	if propertyInit.template == "" {
		// use default from config
		propertyInit.template = ConfUser.Template
	}

	cDefault(os.Stdout, "[ - ] Creating skeleton inside folder %q...\n", target)
	if _, err := os.Stat(target); err != nil {
		if err := os.MkdirAll(fmt.Sprintf("%s/images", target), 0755); err != nil {
			return err
		}
	}

	if _, err := os.Stat(cfgFile); err != nil || propertyInit.force {
		f, err := os.Create(cfgFile)
		if err != nil {
			return err
		}
		defer f.Close()

		fmt.Fprintf(f, configFmt, propertyInit.template)

		cSuccess(os.Stdout, "[ + ] Generated config file %q.\n", cfgFile)
	} else {
		cDefault(os.Stdout, "[ - ] Config already exists, skipping init (use -f to overwrite)...\n")
	}

	if _, err := os.Stat(slides); err != nil {
		f, err := os.Create(slides)
		if err != nil {
			return err
		}
		defer f.Close()

		fmt.Fprintf(f, slidesFmt, target)
		cSuccess(os.Stdout, "[ + ] Generated slides file %q.\n", slides)
	} else {
		cDefault(os.Stdout, "[ - ] Slides already existing, skipping init...\n")
	}

	return nil
}
