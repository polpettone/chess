package model

import (
	"strings"
)

type MoveTestCase struct {
	Number       int
	Moves        []Move
	InitialBoard Board
	WantedBoard  Board
}

type Case struct {
	Name    string
	Piece   Piece
	Current Pos
	Target  Pos
	Board   Board
}

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
				Current: *P(item[2]),
				Target:  *P(item[3]),
				Name:    line,
			}
			testCases = append(testCases, c)
		}
	}
	return testCases
}

func GenerateTestCases(raw string, board Board) []Case {

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

func generateMoveTestCase(raw string, number int) (*MoveTestCase, error) {
	lines := strings.Split(raw, "\n")

	var initialBoard Board
	if strings.Contains(lines[1], "NEW") {
		initialBoard = NewBoard()
	} else if strings.Contains(lines[1], "EMPTY") {
		initialBoard = *NewEmptyBoard()
	} else {
		initBoardRaw := strings.Join(lines[:11], "\n")
		i, err := NewBoardFromString(initBoardRaw)
		if err != nil {
			return nil, err
		}
		initialBoard = *i
	}

	wantedBoardRaw := strings.Join(lines[len(lines)-11:], "\n")
	wantedBoard, err := NewBoardFromString(wantedBoardRaw)
	if err != nil {
		return nil, err
	}

	moveTestCase := &MoveTestCase{
		Number:       number,
		Moves:        generateMoves(lines),
		InitialBoard: initialBoard,
		WantedBoard:  *wantedBoard,
	}

	return moveTestCase, nil
}
