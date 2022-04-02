package model

import (
	"testing"
)

type RookCase struct {
	name    string
	piece   Piece
	current Pos
	target  Pos
	board   Board
}

func TestRookIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WR A1 B2
# WR A1 B8

# BR A1 B2
# BR A1 B8


`
	tests := GenerateTestCases(testCasesRaw, NewBoard())
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := tt.Piece.CheckMoveAllowed(tt.Current, tt.Target)
			if result != false {
				t.Errorf("wanted false, got true")
			}
			if err == nil {
				t.Errorf("wanted error, got none")
			}
		})
	}
}

func TestRookLegalMoves(t *testing.T) {
	testCasesRaw := `
# BR B8 B6
# WR A1 A2
# WR B1 B8

# BR C8 C7

`
	board, _ := NewBoardFromString(boardWithRooks)
	tests := GenerateTestCases(testCasesRaw, *board)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := tt.Piece.CheckMoveAllowed(tt.Current, tt.Target)
			if result != true {
				t.Errorf("wanted true, got false")
			}
			if err != nil {
				t.Errorf("wanted no error, got none %s", err)
			}
		})
	}
}

const boardWithRooks = ` 
    A   B   C   D   E   F   G   H  
8 [BR][BR][BR][  ][  ][  ][  ][  ] 8
7 [  ][  ][  ][  ][  ][  ][  ][  ] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [  ][  ][  ][  ][  ][  ][  ][  ] 2
1 [WR][WR][  ][  ][  ][  ][  ][  ] 1
    A   B   C   D   E   F   G   H 
`
