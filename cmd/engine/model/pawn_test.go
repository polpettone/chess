package model_test

import (
	"testing"

	"github.com/polpettone/chess/cmd/engine/model"
)

func TestIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WP A2 A5

# WP A2 B2

# WP A3 A7
# WP A1 A3

# BP A7 A4

# BP E7 H6
`

	tests := model.GeneratePieceMoveTestCases(testCasesRaw)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := tt.Piece.CheckMoveAllowed(tt.Current, tt.Target)
			if err == nil {
				t.Errorf("wanted error, got none")
			}
			if result != false {
				t.Errorf("wanted false, got true")
			}
		})
	}
}

func TestPawnLegalMoves(t *testing.T) {
	testCasesRaw := `
# WP A2 A3
# WP B2 B4
# BP A7 A5
# WP C2 D3
`

	tests := model.GeneratePieceMoveTestCases(testCasesRaw)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := tt.Piece.CheckMoveAllowed(tt.Current, tt.Target)
			if result != true {
				t.Errorf("wanted true, got false")
			}
			if err != nil {
				t.Errorf("wanted no error, got %s", err)
			}

		})
	}
}
