package model

import (
	"testing"
)

func TestBishopIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WB D8 E6
# WB A8 A1
`
	board, _ := NewBoardFromString(boardWithBishops)
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

func TestBishopLegalMoves(t *testing.T) {
	testCasesRaw := `
# BB B1 H7
# BB F1 A6
`
	board, _ := NewBoardFromString(boardWithBishops)
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

const boardWithBishops = ` 
    A   B   C   D   E   F   G   H  
8 [WB][  ][  ][WB][  ][  ][  ][  ] 8
7 [  ][  ][  ][  ][  ][  ][  ][  ] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [  ][  ][  ][  ][  ][  ][  ][  ] 2
1 [  ][BB][  ][  ][  ][BB][  ][  ] 1
    A   B   C   D   E   F   G   H 
`
