package model

import (
	"fmt"
	"github.com/polpettone/chess/cmd/engine/model/foo"
	"github.com/polpettone/chess/cmd/engine/model/piece"
	"reflect"
	"strings"
	"testing"
)

func TestLegalMovePieceTo(t *testing.T) {
	testCasesRaw := `
# WP A2 A4
# WP B2 B4
# BP A7 A6
`
	tests := GenerateTestCases(testCasesRaw, NewBoard())

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := tt.Board.MovePiece(tt.Current, tt.Target, tt.Piece)
			if err != nil {
				t.Errorf("wanted no error, got %s", err)
			}

		})
	}
}

func TestIllegalMovePieceTo(t *testing.T) {

	testCasesRaw := `
# WR A1 A2
# WR A2 A3
`
	tests := GenerateTestCases(testCasesRaw, NewBoard())

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := tt.Board.MovePiece(tt.Current, tt.Target, tt.Piece)
			if err != nil {
				me, ok := err.(*MoveError)
				if ok {
					fmt.Printf("%s \n", me.Err.Error())
				} else {
					t.Errorf("error has wrong type")
				}
			} else {
				t.Errorf("wanted error, got none")
			}
		})
	}
}

func TestMovePieceTo(t *testing.T) {
	testCasesRaw := `
# WP A2 A3
`
	tests := GenerateTestCases(testCasesRaw, NewBoard())

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			_, err := tt.Board.MovePiece(tt.Current, tt.Target, tt.Piece)
			wanted, _ := NewBoardFromString(wantedBoard)

			if err != nil {
				me, ok := err.(*MoveError)
				if ok {
					t.Errorf("%s \n", me.Err.Error())
				}
			}

			if !reflect.DeepEqual(tt.Board, *wanted) {
				t.Errorf(" \n wanted: \n%s \n got: \n%s \n",
					wanted.Print(nil),
					tt.Board.Print(nil))
			}
		})
	}

}

func TestGetPieceAtPos(t *testing.T) {

	tests := []struct {
		name       string
		board      Board
		pos        foo.Pos
		wantColor  piece.Color
		wantSymbol string
	}{

		{
			pos:        *foo.NewPos(0, 1),
			name:       fmt.Sprintf("Test Piece at Pos:  %s", foo.NewPos(0, 1)),
			board:      NewBoard(),
			wantColor:  piece.WHITE,
			wantSymbol: "WP",
		},

		{
			pos:        *foo.NewPos(0, 0),
			name:       fmt.Sprintf("Test Piece at Pos:  %s", foo.NewPos(0, 0)),
			board:      NewBoard(),
			wantColor:  piece.WHITE,
			wantSymbol: "WR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := NewBoard()

			piece := board.GetPieceAtPos(tt.pos)

			if piece == nil {
				t.Errorf("expected piece at pos %s but got nil", tt.pos)
			}

			if piece.GetSymbol() != tt.wantSymbol {
				t.Errorf("expected piece with symbol %s and color %s at pos %s but got %s",
					tt.wantSymbol,
					fmt.Sprint(tt.wantColor),
					tt.pos,
					piece.GetSymbol())
			}
		})
	}
}

const board = ` 
    A   B   C   D   E   F   G   H  
8 [BR][BN][BB][BQ][BK][BB][BN][BR] 8
7 [BP][BP][BP][BP][BP][BP][BP][BP] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [WP][WP][WP][WP][WP][WP][WP][WP] 2
1 [WR][WN][WB][WQ][WK][WB][WN][WR] 1
    A   B   C   D   E   F   G   H 
`

const wantedBoard = ` 
    A   B   C   D   E   F   G   H  
8 [BR][BN][BB][BQ][BK][BB][BN][BR] 8
7 [BP][BP][BP][BP][BP][BP][BP][BP] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [WP][  ][  ][  ][  ][  ][  ][  ] 3
2 [  ][WP][WP][WP][WP][WP][WP][WP] 2
1 [WR][WN][WB][WQ][WK][WB][WN][WR] 1
    A   B   C   D   E   F   G   H 
`

func DebugParsing(t *testing.T) {
	slice := strings.Split(board, "\n")
	if len(slice) != 12 {
		t.Errorf("wanted 12 got %d", len(slice))
	}

	fmt.Println("Board Parsing")

	y := 7

	emptyBoard := NewEmptyBoard()

	for _, line := range slice {
		if strings.Contains(line, "[") {
			lineSlice := strings.Split(line, "[")
			x := 0
			for _, l := range lineSlice {
				fmt.Println(l)
				if strings.Contains(l, "]") {
					pieceSymbol := l[0:2]
					fmt.Printf("p: %s %d (%d,%d)\n", pieceSymbol, len(pieceSymbol), x, y)

					emptyBoard.SetPieceAtPos(*foo.NewPos(x, y), PieceFrom(pieceSymbol))

					x = x + 1
				}
			}
			y = y - 1
		}
	}

	fmt.Println(emptyBoard.Print(nil))
}
