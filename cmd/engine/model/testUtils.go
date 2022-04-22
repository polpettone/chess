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
	Name  string
	Move  Move
	Board Board
}

func GenerateTestCases(raw string, board Board) []Case {

	lines := strings.Split(raw, "\n")

	var testCases []Case

	for _, line := range lines {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			move, _ := MoveFromString(strings.Join(item[1:4], " "))
			c := Case{
				Move:  *move,
				Name:  line,
				Board: board,
			}
			testCases = append(testCases, c)
		}
	}
	return testCases
}

func generateMoves(raw []string) []Move {

	var moves []Move
	for _, line := range raw {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			move, _ := MoveFromString(strings.Join(item[1:4], " "))
			moves = append(moves, *move)
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

	startIndexWantedBoard := len(lines) - 11
	var wantedBoard Board
	if startIndexWantedBoard > 0 && len(lines[startIndexWantedBoard:]) > 9 {
		wantedBoardRaw := strings.Join(lines[len(lines)-11:], "\n")
		mayWantedBoard, err := NewBoardFromString(wantedBoardRaw)
		if mayWantedBoard != nil {
			wantedBoard = *mayWantedBoard
		}
		if err != nil {
			return nil, err
		}
	}

	moveTestCase := &MoveTestCase{
		Number:       number,
		Moves:        generateMoves(lines),
		InitialBoard: initialBoard,
		WantedBoard:  wantedBoard,
	}

	return moveTestCase, nil
}
