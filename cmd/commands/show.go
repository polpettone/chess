package commands

import (
	"fmt"

	"github.com/polpettone/chess/cmd/engine"
	"github.com/spf13/cobra"
)

func ShowCmd() *cobra.Command {
	return &cobra.Command{
		Use: "show",

		Run: func(command *cobra.Command, args []string) {
			err := handleShowCommand(args)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
}

func handleShowCommand(args []string) error {
	fmt.Println("Polpettone Chess")
	boardFile := "current.chess"
	board, err := engine.LoadBoardFromFile(boardFile)

	if err != nil {
		return err
	}

	fmt.Println(board.Print(nil))
	return nil
}

func init() {
	showCmd := ShowCmd()
	rootCmd.AddCommand(showCmd)
}
