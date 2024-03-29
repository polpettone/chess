package model

import (
	"strings"
)

type PieceMoveTestCase struct {
	Name    string
	Piece   Piece
	Current Pos
	Target  Pos
}

func GeneratePieceMoveTestCases(raw string) []PieceMoveTestCase {
	lines := strings.Split(raw, "\n")

	var testCases []PieceMoveTestCase

	for _, line := range lines {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			c := PieceMoveTestCase{
				Piece:   PieceFrom(item[1]),
				Current: *PositionFromString(item[2]),
				Target:  *PositionFromString(item[3]),
				Name:    line,
			}
			testCases = append(testCases, c)
		}
	}
	return testCases
}
