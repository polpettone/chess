package model

import (
	"reflect"
	"strconv"
	"testing"
)

type SquaresAroundTestCase struct {
	Board         string
	Square        Square
	SquaresAround []Square
	WantErr       bool
}

func generatePieceAroundTestCase() []SquaresAroundTestCase {

	return []SquaresAroundTestCase{

		{
			WantErr:       true,
			Square:        SquareFromString("BK", "D6"),
			SquaresAround: nil,
			Board: `
    A   B   C   D   E   F   G   H  
8 [  ][  ][  ][  ][  ][  ][  ][  ] 8
7 [  ][  ][  ][  ][  ][  ][  ][  ] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [  ][  ][  ][  ][  ][  ][  ][  ] 2
1 [  ][  ][  ][  ][  ][  ][  ][  ] 1
    A   B   C   D   E   F   G   H 
`,
		},

		{
			Square: SquareFromString("BK", "D6"),
			SquaresAround: []Square{
				SquareFromString("WQ", "D4"),
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

			squares, err := FindSquaresAround(*board, tC.Square)

			if !tC.WantErr && err != nil {
				t.Errorf("Wanted no error, got: %v", err)
				return
			}

			if tC.WantErr && err == nil {
				t.Errorf("Wanted error, got none")
			}

			if !reflect.DeepEqual(squares, tC.SquaresAround) {
				t.Errorf("wanted %s got %s", tC.SquaresAround, squares)
			}

		})
	}

}
