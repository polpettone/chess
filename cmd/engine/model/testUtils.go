package model

import (
	"github.com/polpettone/chess/cmd/engine/model/foo"
	"github.com/polpettone/chess/cmd/engine/model/piece"
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
	Piece   piece.Piece
	Current foo.Pos
	Target  foo.Pos
	Board   Board
}

func GenerateTestCases(raw string, board Board) []Case {

	lines := strings.Split(raw, "\n")

	var testCases []Case

	for _, line := range lines {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			c := Case{
				Piece:   PieceFrom(item[1]),
				Current: *foo.PositionFromString(item[2]),
				Target:  *foo.PositionFromString(item[3]),
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
	Piece   piece.Piece
	Current foo.Pos
	Target  foo.Pos
}

func generateMoves(raw []string) []Move {

	var moves []Move
	for _, line := range raw {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			m := Move{
				Piece:   PieceFrom(item[1]),
				Current: *foo.PositionFromString(item[2]),
				Target:  *foo.PositionFromString(item[3]),
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
