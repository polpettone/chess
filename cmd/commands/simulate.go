package commands

import (
	"fmt"

	"github.com/polpettone/chess/cmd/engine"
	"github.com/spf13/cobra"
)

func SimulateCmd() *cobra.Command {
	return &cobra.Command{
		Use: "simulate",

		Run: func(command *cobra.Command, args []string) {
			err := handleSimulateCommand()
			if err != nil {
				fmt.Println(err)
			}
		},
	}
}

func handleSimulateCommand() error {
	game := engine.Game{}
	err := game.Play(false)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	simulateCmd := SimulateCmd()
	rootCmd.AddCommand(simulateCmd)
}
