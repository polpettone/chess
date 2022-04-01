package model

import (
	"testing"
)

func TestQueenIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WQ D8 E6
`
	board, _ := NewBoardFromString(boardWithQueens)
	tests := GenerateTestCases(testCasesRaw, *board)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := tt.Piece.Move(tt.Current, tt.Target, tt.Board)
			if err == nil {
				t.Errorf("wanted error, got none")
			}
		})
	}
}

func TestQueenLegalMoves(t *testing.T) {
	testCasesRaw := `
# WQ D8 D6
# WQ D5 H5

# BQ B2 H8
# BQ F1 A6
`
	board, _ := NewBoardFromString(boardWithQueens)
	tests := GenerateTestCases(testCasesRaw, *board)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := tt.Piece.Move(tt.Current, tt.Target, tt.Board)
			if err != nil {
				me, ok := err.(*MoveError)
				if ok {
					t.Errorf("%s \n", me.Err.Error())
				} else {
					t.Errorf("error has wrong type")
				}
			}
		})
	}
}

const boardWithQueens = ` 
    A   B   C   D   E   F   G   H  
8 [  ][  ][  ][WQ][  ][  ][  ][  ] 8
7 [  ][  ][  ][  ][  ][  ][  ][  ] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][WQ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [  ][BQ][  ][  ][  ][  ][  ][  ] 2
1 [  ][  ][  ][  ][  ][BQ][  ][  ] 1
    A   B   C   D   E   F   G   H 
`
