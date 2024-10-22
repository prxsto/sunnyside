package cmd

import (
	"github.com/spf13/cobra"
)

func setup() *cobra.Command {
	init := &cobra.Command{
		Use:     "setup",
		Short:   "init sunnyside cfg",
		Long:    "initial sunnyside configuration file",
		Example: "sunnyside setup",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return init
}
