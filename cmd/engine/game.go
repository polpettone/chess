package engine

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/polpettone/chess/cmd/engine/model"
)

type Game struct {
}

func randomPiece(color model.Color) model.Piece {
	pieceChars := []string{"P"}
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(len(pieceChars))
	pieceChar := pieceChars[v]

	var piece model.Piece
	if color == model.WHITE {
		piece = model.PieceFrom("W" + pieceChar)
	} else {
		piece = model.PieceFrom("B" + pieceChar)
	}
	return piece
}

func (g Game) Play(errorPrinting bool) error {
	board := model.NewBoard()

	white := true

	var choosenPiece model.Piece
	lastMoveSuccess := true
	for n := 0; n < 1000*1000*1000; n++ {
		if lastMoveSuccess {
			if white {
				choosenPiece = randomPiece(model.WHITE)
			} else {
				choosenPiece = randomPiece(model.BLACK)
			}
		}

		fromPositions := board.FindPiecePositions(choosenPiece)
		from := fromPositions[0]

		rand.Seed(time.Now().UnixNano())
		tx := rand.Intn(7)
		ty := rand.Intn(7)
		to := model.Pos{X: tx, Y: ty}

		move := model.Move{
			From:  from,
			To:    to,
			Piece: choosenPiece,
		}
		_, err := board.MovePiece(move)

		if err != nil {
			lastMoveSuccess = false
			if errorPrinting {
				fmt.Printf("%s %s %s : ", choosenPiece.GetSymbol(), from.Print(), to.Print())
				fmt.Println(err)
			}
		} else {
			lastMoveSuccess = true
			if white {
				white = false
			} else {
				white = true
			}

			fmt.Printf("n: %d\n", n)
			fmt.Println(board.Print([]string{move.From.Print(), move.To.Print()}))
		}
	}

	return nil
}
