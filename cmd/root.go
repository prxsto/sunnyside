package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// entry function
func Execute() error {
	rootCmd := &cobra.Command{
		Version: "v0.0.1",
		Use:     "sunnyside",
		Long:    "sunnyside is a TUI-based recipe repository & meal planner",
		Example: "sunnyside",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	// add sub-commands
	rootCmd.AddCommand(setup())

	// dir, err := os.UserHomeDir()
	// if err != nil {
	// 	return err
	// }

	return rootCmd.ExecuteContext(context.Background())
}
