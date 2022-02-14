package engine

import (
	"strings"
	"testing"
)

type Case struct {
	name    string
	piece   Piece
	current Pos
	target  Pos
	board   Board
}

func TestIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WP A3 A4
# WP A2 A5

# WP A2 B2

# WP A3 A7
# WP A1 A3

# BP A7 A4
`
	tests := generateTestCases(testCasesRaw)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.piece.Move(tt.current, tt.target, tt.board)
			if err == nil {
				t.Errorf("wanted error, got none")
			}
		})
	}
}

func TestLegalMoves(t *testing.T) {
	testCasesRaw := `
# WP A2 A3
# WP A2 A4
# BP A7 A5
`
	tests := generateTestCases(testCasesRaw)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.piece.Move(tt.current, tt.target, tt.board)
			if err != nil {
				t.Errorf("wanted no error, got %s", err)
			}
		})
	}
}

func generateTestCases(raw string) []Case {

	lines := strings.Split(raw, "\n")

	var testCases []Case

	for _, line := range lines {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			c := Case{
				piece:   PieceFrom(item[1]),
				current: *P(item[2]),
				target:  *P(item[3]),
				name:    line,
				board:   NewBoard(),
			}
			testCases = append(testCases, c)
		}
	}
	return testCases
}
