package model

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

func TestKingLegalMoves(t *testing.T) {
	testCasesRaw := `
# WK D8 D7
# WK D5 C5

# BK B2 C3
# BK F1 E2
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
