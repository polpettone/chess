package model

import (
	"strconv"
	"testing"
)

type SquaresAroundTestCase struct {
	Board        string
	Square       Square
	SquaresAround []Square
}

func generatePieceAroundTestCase() []SquaresAroundTestCase {

	return []SquaresAroundTestCase{

		{
			Square:        
			Square{Piece: PieceFrom("BK"), Pos: *PositionFromString("D6")},
			SquaresAround: []Squares{
					{
					Piece: PieceFrom("WQ")
				    Pos: *PositionFromString("D4")
			},


			Board: `
    A   B   C   D   E   F   G   H  
8 [  ][  ][  ][  ][  ][  ][  ][  ] 8
7 [  ][  ][  ][  ][  ][  ][  ][  ] 7
6 [  ][  ][  ][BK][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][WQ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [  ][  ][  ][  ][  ][  ][  ][  ] 2
1 [  ][  ][  ][  ][  ][  ][  ][  ] 1
    A   B   C   D   E   F   G   H 
`,
		},
	}
}

func TestPiecesAround(t *testing.T) {

	testCases := generatePieceAroundTestCase()

	for i, tC := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			board, err := NewBoardFromString(tC.Board)
			if err != nil {
				t.Errorf("Setup Failure")
				return
			}

			squares := FindSqauresAround(board, square)

		})
	}

}
