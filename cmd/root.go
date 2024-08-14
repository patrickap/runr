package cmd

import (
	"github.com/patrickap/runr/m/v2/internal/command"
	"github.com/patrickap/runr/m/v2/internal/config"
	"github.com/patrickap/runr/m/v2/internal/log"
	"github.com/spf13/cobra"
)

var version string

var rootCmd = &cobra.Command{
	Use:           "runr",
	Version:       version,
	Args:          cobra.ExactArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	commands := config.Instance().GetCommands()

	for commandName, commandConfig := range commands {
		commandName := commandName
		commandConfig := commandConfig

		childCmd := &cobra.Command{
			Use:          commandName,
			SilenceUsage: true,
			RunE: func(c *cobra.Command, args []string) error {
				log.Instance().Info().Msgf("Running command: %s", commandName)
				cmd := command.BuildCommand(&commandConfig)
				err := cmd.Run()
				if err != nil {
					log.Instance().Error().Msgf("Failed to run command: %s: %v", commandName, err)
					return err
				}

				return nil
			},
		}

		rootCmd.AddCommand(childCmd)
	}
}

func Execute() error {
	return rootCmd.Execute()
}
