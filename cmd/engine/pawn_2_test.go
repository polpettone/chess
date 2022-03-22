package engine

import (
	"testing"
)

func TestIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WP A3 A4
# WP A2 A5

# WP A2 B2

# WP A3 A7
# WP A1 A3

# BP A7 A4
`
	tests := generateTestCases(testCasesRaw, NewBoard())
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := tt.Piece.Move(tt.Current, tt.Target, tt.Board)
			if err == nil {
				t.Errorf("wanted error, got none")
			}
		})
	}
}

func TestPawnLegalMoves(t *testing.T) {
	testCasesRaw := `
# WP A2 A3
# WP B2 B4
# BP A7 A5
`
	tests := generateTestCases(testCasesRaw, NewBoard())
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := tt.Piece.Move(tt.Current, tt.Target, tt.Board)
			if err != nil {
				t.Errorf("wanted no error, got %s", err)
			}
		})
	}
}
