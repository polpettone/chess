package commands

import (
	"github.com/polpettone/chess/cmd/config"
	"github.com/polpettone/chess/cmd/engine"
	"github.com/spf13/cobra"
)

func SimulateCmd() *cobra.Command {
	return &cobra.Command{
		Use: "simulate",

		Run: func(command *cobra.Command, args []string) {
			err := handleSimulateCommand()
			if err != nil {
				config.Log.ErrorLog.Println(err)
			}
		},
	}
}

func handleSimulateCommand() error {
	game := engine.Simulation{}
	err := game.Play(true)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	simulateCmd := SimulateCmd()
	rootCmd.AddCommand(simulateCmd)
}
