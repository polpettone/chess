package tests

import (
	"github.com/polpettone/chess/cmd/engine/model"
	"github.com/polpettone/chess/cmd/engine/model/foo"
	"github.com/polpettone/chess/cmd/engine/model/piece"
	"strings"
)

type PieceMoveTestCase struct {
	Name    string
	Piece   piece.Piece
	Current foo.Pos
	Target  foo.Pos
}

func GeneratePieceMoveTestCases(raw string) []PieceMoveTestCase {
	lines := strings.Split(raw, "\n")

	var testCases []PieceMoveTestCase

	for _, line := range lines {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			c := PieceMoveTestCase{
				Piece:   model.PieceFrom(item[1]),
				Current: *foo.PositionFromString(item[2]),
				Target:  *foo.PositionFromString(item[3]),
				Name:    line,
			}
			testCases = append(testCases, c)
		}
	}
	return testCases
}
