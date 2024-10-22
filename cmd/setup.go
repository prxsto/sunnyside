package cmd

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/prxsto/sunnyside/internal/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	ENV = "ss"
)

func setup() *cobra.Command {
	init := &cobra.Command{
		Use:     "setup",
		Short:   "init sunnyside cfg",
		Long:    "initial sunnyside configuration file",
		Example: "sunnyside setup",
		Aliases: []string{"s"},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			viper.AutomaticEnv()
			viper.SetEnvPrefix(ENV)
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			if err := tea.Program(tui.NewInitPrompt(viper.GetString(cfgPath), homeDir)).Run(); err != nil {
				return err
			}
			return nil
		},
	}
	return init
}
