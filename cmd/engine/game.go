package engine

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/polpettone/chess/cmd/engine/model"
)

type Game struct {
}

func (g Game) Play() error {
	board := model.NewBoard()

	pieces := []string{"P", "R", "N", "B", "K", "Q"}

	white := false

	rand.Seed(time.Now().UnixNano())
	for n := 0; n < 1000*1000*1000; n++ {
		v := rand.Intn(len(pieces)-1) + 1
		piece := pieces[v]
		var p string
		if white {
			p = "W" + piece
			white = false
		} else {
			p = "B" + piece
			white = true
		}
		x := rand.Intn(7)
		y := rand.Intn(7)

		from := model.Pos{X: x, Y: y}
		to := model.Pos{X: x, Y: y}

		choosenPiece := model.PieceFrom(p)

		move := model.Move{
			From:  from,
			To:    to,
			Piece: choosenPiece,
		}
		_, err := board.MovePiece(move)

		if err != nil {

		} else {
			fmt.Printf("n: %d\n", n)
			fmt.Println(board.Print([]string{move.From.Print(), move.To.Print()}))
		}

	}

	return nil
}
