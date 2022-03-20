package engine

import "strings"

type Case struct {
	Name    string
	Piece   Piece
	Current Pos
	Target  Pos
	Board   Board
}

func generateTestCases(raw string, board Board) []Case {

	lines := strings.Split(raw, "\n")

	var testCases []Case

	for _, line := range lines {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			c := Case{
				Piece:   PieceFrom(item[1]),
				Current: *P(item[2]),
				Target:  *P(item[3]),
				Name:    line,
				Board:   board,
			}
			testCases = append(testCases, c)
		}
	}
	return testCases
}
