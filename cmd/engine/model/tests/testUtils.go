package tests

import (
	"github.com/polpettone/chess/cmd/engine/model"
	"github.com/polpettone/chess/cmd/engine/model/foo"
	"github.com/polpettone/chess/cmd/engine/model/piece"
	"strings"
)

type MoveTestCase struct {
	Number       int
	Movements    []model.Movement
	InitialBoard model.Board
	WantedBoard  model.Board
}

type Case struct {
	Name     string
	Movement model.Movement
	Board    model.Board
}

func GenerateTestCases(raw string, board model.Board) []Case {

	lines := strings.Split(raw, "\n")

	var testCases []Case

	for _, line := range lines {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			movement, _ := model.MoveFromString(strings.Join(item[1:4], " "))
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
	Piece   piece.Piece
	Current foo.Pos
	Target  foo.Pos
}

func generateMoves(raw []string) []model.Movement {

	var movements []model.Movement
	for _, line := range raw {
		if strings.Contains(line, "#") {
			item := strings.Split(line, " ")
			movement, _ := model.MoveFromString(strings.Join(item[1:4], " "))
			movements = append(movements, *movement)
		}
	}
	return movements
}

func generateMoveTestCase(raw string, number int) (*MoveTestCase, error) {
	lines := strings.Split(raw, "\n")

	var initialBoard model.Board
	if strings.Contains(lines[1], "NEW") {
		initialBoard = model.NewBoard()
	} else if strings.Contains(lines[1], "EMPTY") {
		initialBoard = *model.NewEmptyBoard()
	} else {
		initBoardRaw := strings.Join(lines[:11], "\n")
		i, err := model.NewBoardFromString(initBoardRaw)
		if err != nil {
			return nil, err
		}
		initialBoard = *i
	}

	wantedBoardRaw := strings.Join(lines[len(lines)-11:], "\n")
	wantedBoard, err := model.NewBoardFromString(wantedBoardRaw)
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
