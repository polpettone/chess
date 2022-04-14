package model

import (
	"strings"
)

type MoveTestCase struct {
	Number       int
	Movements    []Movement
	InitialBoard Board
	WantedBoard  Board
}

type Case struct {
	Name     string
	Movement Movement
	Board    Board
}

func GenerateTestCases(raw string, board Board) []Case {

	lines := strings.Split(raw, "\n")

	var testCases []Case

	for _, line := range lines {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			movement, _ := MoveFromString(strings.Join(item[1:4], " "))
			c := Case{
				Movement: *movement,
				Name:     line,
				Board:    board,
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

func generateMoves(raw []string) []Movement {

	var movements []Movement
	for _, line := range raw {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			movement, _ := MoveFromString(strings.Join(item[1:4], " "))
			movements = append(movements, *movement)
		}
	}
	return movements
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
		Movements:    generateMoves(lines),
		InitialBoard: initialBoard,
		WantedBoard:  *wantedBoard,
	}

	return moveTestCase, nil
}
