package model

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestIsCheck(t *testing.T) {

	testCases := generateIsCheckTestCases()

	for i, tC := range testCases {

		t.Run(strconv.Itoa(i), func(t *testing.T) {

			board, err := NewBoardFromString(tC.Board)
			if err != nil {
				t.Errorf("Setup Failure")
				return
			}
			actual, err := board.IsCheck(tC.Color)

			if err != nil && !tC.WantErr {
				t.Errorf("wanted no error got %s", err)
			}

			if tC.IsCheck != actual {
				t.Errorf(" \n wanted check %t, \n but got %t, for \n %s",
					tC.IsCheck, actual, board.Print(nil))
			}
		})
	}
}

func TestIllegalMoves(t *testing.T) {
	counter := 0
	testCaseRaws := generatIllegalMoves()

	for _, testCaseRaw := range testCaseRaws {

		testCase, err := generateMoveTestCase(testCaseRaw, counter)
		counter++

		t.Run(strconv.Itoa(testCase.Number), func(t *testing.T) {
			if err != nil {
				t.Errorf("test setup error %s", err)
			}

			for _, movement := range testCase.Moves {

				_, err := testCase.InitialBoard.MovePiece(movement)
				if err == nil {
					t.Errorf("wanted error got none")
				}
			}
		})
	}

}

func TestLegalMoves(t *testing.T) {
	counter := 0
	testCaseRaws := generatLegalMoves()

	for _, testCaseRaw := range testCaseRaws {

		testCase, err := generateMoveTestCase(testCaseRaw, counter)
		counter++

		t.Run(strconv.Itoa(testCase.Number), func(t *testing.T) {
			if err != nil {
				t.Errorf("test setup error %s", err)
			}

			for _, movement := range testCase.Moves {
				_, err := testCase.InitialBoard.MovePiece(movement)
				if err != nil {
					t.Errorf("wanted no error got %s", err)
				}
			}

			if !reflect.DeepEqual(testCase.InitialBoard.Fields,
				testCase.WantedBoard.Fields) {
				t.Errorf(" \n wanted: \n%s \n got: \n%s \n",
					testCase.WantedBoard.Print(nil),
					testCase.InitialBoard.Print(nil))
			}

		})
	}
}
func TestLegalMovePieceTo(t *testing.T) {
	testCasesRaw := `
# WP A2 A4
# WP B2 B4
# BP A7 A6
`
	tests := GenerateTestCases(testCasesRaw, NewBoard())

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := tt.Board.MovePiece(tt.Move)
			if err != nil {
				t.Errorf("wanted no error, got %s", err)
			}
		})
	}
}

func TestIllegalMovePieceTo(t *testing.T) {

	testCasesRaw := `
# WR A1 A2
# WR A2 A3
# WP C2 C2
`
	tests := GenerateTestCases(testCasesRaw, NewBoard())

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := tt.Board.MovePiece(tt.Move)
			if err != nil {
				me, ok := err.(*MoveError)
				if ok {
					fmt.Printf("%s \n", me.Err.Error())
				} else {
					t.Errorf("error has wrong type")
				}
			} else {
				t.Errorf("wanted error, got none")
			}
		})
	}
}

func TestFindPiecePos(t *testing.T) {
	tests := []struct {
		board     Board
		piece     Piece
		positions []Pos
	}{
		{
			board:     NewBoard(),
			piece:     PieceFrom("BQ"),
			positions: []Pos{*PositionFromString("D8")},
		},

		{
			board:     NewBoard(),
			piece:     PieceFrom("WQ"),
			positions: []Pos{*PositionFromString("D1")},
		},

		{
			board:     NewBoard(),
			piece:     PieceFrom("WR"),
			positions: []Pos{*PositionFromString("A1"), *PositionFromString("H1")},
		},

		{
			board: NewBoard(),
			piece: PieceFrom("WP"),
			positions: []Pos{
				*PositionFromString("A2"),
				*PositionFromString("B2"),
				*PositionFromString("C2"),
				*PositionFromString("D2"),
				*PositionFromString("E2"),
				*PositionFromString("F2"),
				*PositionFromString("G2"),
				*PositionFromString("H2"),
			},
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("Find Positions of %s", tt.piece.GetSymbol())
		t.Run(name, func(t *testing.T) {
			board := NewBoard()
			actualPositions := board.FindPiecePositions(tt.piece)

			if !reflect.DeepEqual(tt.positions, actualPositions) {
				t.Errorf("not equal wanted: %v, actual %v", tt.positions, actualPositions)
			}
		})
	}
}

func TestGetPieceAtPos(t *testing.T) {

	tests := []struct {
		name       string
		board      Board
		pos        Pos
		wantColor  Color
		wantSymbol string
	}{

		{
			pos:        *NewPos(0, 1),
			name:       fmt.Sprintf("Test Piece at Pos:  %s", NewPos(0, 1)),
			board:      NewBoard(),
			wantColor:  WHITE,
			wantSymbol: "WP",
		},

		{
			pos:        *NewPos(0, 0),
			name:       fmt.Sprintf("Test Piece at Pos:  %s", NewPos(0, 0)),
			board:      NewBoard(),
			wantColor:  WHITE,
			wantSymbol: "WR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := NewBoard()

			piece := board.GetPieceAtPos(tt.pos)

			if piece == nil {
				t.Errorf("expected piece at pos %s but got nil", tt.pos)
			}

			if piece.GetSymbol() != tt.wantSymbol {
				t.Errorf("expected piece with symbol %s and color %s at pos %s but got %s",
					tt.wantSymbol,
					fmt.Sprint(tt.wantColor),
					tt.pos,
					piece.GetSymbol())
			}
		})
	}
}
