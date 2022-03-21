package engine

import (
	"testing"
)

func TestKingIllegalMoves(t *testing.T) {
	testCasesRaw := `

# WK D8 D6
# WK D5 B5

# BK B2 D4
# BK F1 H3
`
	board, _ := NewBoardFromString(boardWithKings)
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

func TestKingLegalMoves(t *testing.T) {
	testCasesRaw := `
# WK D8 D7
# WK D5 C5

# BK B2 C3
# BK F1 E2
`
	board, _ := NewBoardFromString(boardWithKings)
	tests := generateTestCases(testCasesRaw, *board)

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

const boardWithKings = ` 
    A   B   C   D   E   F   G   H  
8 [  ][  ][  ][WK][  ][  ][  ][  ] 8
7 [  ][  ][  ][  ][  ][  ][  ][  ] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][WK][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [  ][BK][  ][  ][  ][  ][  ][  ] 2
1 [  ][  ][  ][  ][  ][BK][  ][  ] 1
    A   B   C   D   E   F   G   H 
`
