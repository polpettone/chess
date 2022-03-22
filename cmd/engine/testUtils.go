package engine

import (
	"strings"
)

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

type Move struct {
	Name    string
	Piece   Piece
	Current Pos
	Target  Pos
}

func generateMoves(raw []string) []Move {

	var moves []Move
	for _, line := range raw {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			m := Move{
				Piece:   PieceFrom(item[1]),
				Current: *P(item[2]),
				Target:  *P(item[3]),
				Name:    line,
			}
			moves = append(moves, m)
		}
	}
	return moves
}

func generateMoveTestCase(raw string) (*MoveTestCase, error) {
	lines := strings.Split(raw, "\n")

	initBoardRaw := strings.Join(lines[:11], "\n")
	wantedBoardRaw := strings.Join(lines[len(lines)-11:], "\n")

	movesRaw := lines[11 : len(lines)-11]

	initialBoard, err := NewBoardFromString(initBoardRaw)
	if err != nil {
		return nil, err
	}
	wantedBoard, err := NewBoardFromString(wantedBoardRaw)
	if err != nil {
		return nil, err
	}

	moveTestCase := &MoveTestCase{
		Name:         "unknown",
		Moves:        generateMoves(movesRaw),
		InitialBoard: *initialBoard,
		WantedBoard:  *wantedBoard,
	}

	return moveTestCase, nil
}

type MoveTestCase struct {
	Name         string
	Moves        []Move
	InitialBoard Board
	WantedBoard  Board
}
