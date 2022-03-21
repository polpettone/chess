package engine

import (
	"testing"
)

func TestKnightIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WN B8 C7

# BN D2 E8
# BN F3 H7
`
	board, _ := NewBoardFromString(boardWithKnights)
	tests := generateTestCases(testCasesRaw, *board)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := tt.Piece.Move(tt.Current, tt.Target, tt.Board)
			if err == nil {
				t.Errorf("wanted error, got none")
			}
		})
	}
}

func TestKnightLegalMoves(t *testing.T) {
	testCasesRaw := `
# WN B8 C6

# BN D2 E4
# BN F3 H4
`
	board, _ := NewBoardFromString(boardWithKnights)
	tests := generateTestCases(testCasesRaw, *board)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := tt.Piece.Move(tt.Current, tt.Target, tt.Board)
			if err != nil {
				me, ok := err.(*MoveError)
				if ok {
					t.Errorf("%s \n", me.Error())
				} else {
					t.Errorf("error has wrong type")
				}
			}
		})
	}
}

const boardWithKnights = ` 
    A   B   C   D   E   F   G   H  
8 [  ][WN][  ][  ][  ][  ][  ][  ] 8
7 [  ][  ][  ][  ][  ][  ][  ][  ] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][BN][  ][  ] 3
2 [  ][  ][  ][BN][  ][  ][  ][  ] 2
1 [  ][  ][  ][  ][  ][  ][  ][  ] 1
    A   B   C   D   E   F   G   H 
`
