package engine

import (
	"strings"
	"testing"
)

type RookCase struct {
	name    string
	piece   Piece
	current Pos
	target  Pos
	board   Board
}

func TestRookIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WR A1 B2
# WR A1 B8

# BR A1 B2
# BR A1 B8
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

func TestRookLegalMoves(t *testing.T) {
	testCasesRaw := `
# WR A1 A2
# WR B1 B8

# BR A1 A2
# BR B1 B8
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

func generateRookTestCases(raw string) []Case {

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
