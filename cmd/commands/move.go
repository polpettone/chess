package commands

import (
	"fmt"
	"github.com/polpettone/chess/cmd/engine/repo"

	"github.com/polpettone/chess/cmd/engine"
	"github.com/spf13/cobra"
)

func PlayCmd() *cobra.Command {
	return &cobra.Command{
		Use: "move",

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
	boardFile := "current.chess"
	board, err := repo.LoadBoardFromFileOrCreateNewBoard(boardFile)

	if err != nil {
		fmt.Printf("could not load board from file %s, create new board", boardFile)
		b := engine.NewBoard()
		board = &b
		err = repo.SaveBoardToFile("current.chess", *board)
		if err != nil {
			return err
		}
	}
	fmt.Println(board.Print(nil))

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

	newBoard, err := piece.Move(*currentPos, *targetPos, *board)

	if err != nil {
		return err
	}

	fmt.Println(newBoard.Print([]string{from, to}))

	err = repo.SaveBoardToFile("current.chess", *newBoard)

	if err != nil {
		return err
	}

	return err
}

func init() {
	playCmd := PlayCmd()
	rootCmd.AddCommand(playCmd)
}
