package piece_test

import (
	"github.com/polpettone/chess/cmd/engine/model/piece_test"
	"testing"
)

func TestRookIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WR A1 B2
# WR A1 B8

# BR A1 B2
# BR A1 B8


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

func TestRookLegalMoves(t *testing.T) {
	testCasesRaw := `
# BR B8 B6
# WR A1 A2
# WR B1 B8

# BR C8 C7

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
