package commands

import (
	"fmt"
	"github.com/polpettone/chess/cmd/engine/model"
	"github.com/polpettone/chess/cmd/engine/repo"

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
		fmt.Printf("could not load model from file %s, create new model", boardFile)
		b := model.NewBoard()
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

	piece := model.PieceFrom(p)
	if piece == nil {
		return fmt.Errorf("%s: unknown piece", p)
	}

	currentPos := model.PositionFromString(from)
	if currentPos == nil {
		return fmt.Errorf("%s: invalid position", from)
	}

	targetPos := model.PositionFromString(to)
	if targetPos == nil {
		return fmt.Errorf("%s: invalid position", to)
	}

	_, err = board.MovePiece(model.Move{From: *currentPos, To: *targetPos, Piece: piece})

	if err != nil {
		return err
	}

	fmt.Println(board.Print([]string{from, to}))

	err = repo.SaveBoardToFile("current.chess", *board)

	if err != nil {
		return err
	}

	return err
}

func init() {
	playCmd := PlayCmd()
	rootCmd.AddCommand(playCmd)
}
