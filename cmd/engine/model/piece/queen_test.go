package piece

import (
	"testing"
)

func TestQueenIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WQ D8 E6
`

	tests := GeneratePieceMoveTestCases(testCasesRaw)
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

func TestQueenLegalMoves(t *testing.T) {
	testCasesRaw := `
# WQ D8 D6
# WQ D5 H5

# BQ B2 H8
# BQ F1 A6
`

	tests := GeneratePieceMoveTestCases(testCasesRaw)
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
