package model

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_PrintMovement(t *testing.T) {
	raw := "WP A2 A4"
	movement, _ := MoveFromString(raw)
	actual := movement.Print()
	if actual != raw {
		t.Errorf("actual %s not equal wanted %s", actual, raw)
	}
}

func Test_ValidMoveFromString(t *testing.T) {
	tests := []struct {
		move string
		want Movement
	}{

		{
			move: "WP A2 A4",
			want: Movement{
				Piece: PieceFrom("WP"),
				From:  *PositionFromString("A2"),
				To:    *PositionFromString("A4"),
			},
		},

		{
			move: "BP A2 A4",
			want: Movement{
				Piece: PieceFrom("BP"),
				From:  *PositionFromString("A2"),
				To:    *PositionFromString("A4"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.move, func(t *testing.T) {
			actual, err := MoveFromString(tt.move)
			if err != nil {
				t.Errorf("wanted no error, got %s", err)
				return
			}
			if !reflect.DeepEqual(*actual, tt.want) {
				t.Errorf("%v not equal %v", actual, tt.want)
			}
		})
	}
}

func Test_InvalidMoveFromString(t *testing.T) {
	tests := []struct {
		move string
	}{

		{move: "A4"},
		{move: "A2 A4"},
		{move: "X2 A4"},
		{move: ""},
	}
	for _, tt := range tests {
		t.Run(tt.move, func(t *testing.T) {
			_, err := MoveFromString(tt.move)
			if err == nil {
				t.Errorf("wanted error, got none")
			}
		})
	}
}

func TestLegalMovePieceTo(t *testing.T) {
	testCasesRaw := `
# WP A2 A4
# WP B2 B4
# BP A7 A6
`
	tests := GenerateTestCases(testCasesRaw, NewBoard())

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := tt.Board.MovePiece(tt.Movement)
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
			_, err := tt.Board.MovePiece(tt.Movement)
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

			_, err := tt.Board.MovePiece(tt.Movement)
			wanted, _ := NewBoardFromString(wantedBoard)

			if err != nil {
				me, ok := err.(*MoveError)
				if ok {
					t.Errorf("%s \n", me.Err.Error())
				}
			}

			if !reflect.DeepEqual(tt.Board.Fields, wanted.Fields) {
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
		pos        Pos
		wantColor  Color
		wantSymbol string
	}{

		{
			pos:        *NewPos(0, 1),
			name:       fmt.Sprintf("Test Piece at Pos:  %s", NewPos(0, 1)),
			board:      NewBoard(),
			wantColor:  WHITE,
			wantSymbol: "WP",
		},

		{
			pos:        *NewPos(0, 0),
			name:       fmt.Sprintf("Test Piece at Pos:  %s", NewPos(0, 0)),
			board:      NewBoard(),
			wantColor:  WHITE,
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