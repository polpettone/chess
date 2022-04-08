package piece_test

import (
	"github.com/polpettone/chess/cmd/engine/model/piece_test"
	"testing"
)

func TestKnightIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WN B8 C7

# BN D2 E8
# BN F3 H7
`

	tests := piece.GeneratePieceMoveTestCases(testCasesRaw)
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

	tests := piece.GeneratePieceMoveTestCases(testCasesRaw)
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
