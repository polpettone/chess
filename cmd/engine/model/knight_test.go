package model

import (
	"testing"
)

func TestKnightIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WN B8 C7

# BN D2 E8
# BN F3 H7
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

func TestKnightLegalMoves(t *testing.T) {
	testCasesRaw := `
# WN B8 C6

# BN D2 E4
# BN F3 H4
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
