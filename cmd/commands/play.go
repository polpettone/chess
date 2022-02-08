package commands

import (
	"fmt"

	"github.com/polpettone/chess/cmd/engine"
	"github.com/spf13/cobra"
)

func PlayCmd() *cobra.Command {
	return &cobra.Command{
		Use: "play",

		Run: func(command *cobra.Command, args []string) {
			if len(args) != 3 {
				fmt.Println("Need 3 arguments")
				return
			}
			err := handlePlayCommand(args)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
}

func handlePlayCommand(args []string) error {

	fmt.Println("Polpettone Chess")
	board := engine.NewBoard()
	fmt.Println(board.Print())

	p := args[0]
	from := args[1]
	to := args[2]

	piece := engine.PieceFrom(p)
	if piece == nil {
		return fmt.Errorf("%s: unknown piece", p)
	}

	currentPos := engine.P(from)
	if currentPos == nil {
		return fmt.Errorf("%s: invalid position", from)
	}

	targetPos := engine.P(to)
	if targetPos == nil {
		return fmt.Errorf("%s: invalid position", to)
	}

	newBoard, err := piece.Move(*currentPos, *targetPos, board)

	if err != nil {
		return err
	}

	fmt.Println(newBoard.Print())

	return nil
}

func init() {
	playCmd := PlayCmd()
	rootCmd.AddCommand(playCmd)
}
