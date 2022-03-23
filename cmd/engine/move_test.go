package engine

import (
	"reflect"
	"strconv"
	"testing"
)

func generate() []string {
	return []string{
		`
    A   B   C   D   E   F   G   H  
8 [BR][BN][BB][BQ][BK][BB][BN][BR] 8
7 [BP][BP][BP][BP][BP][BP][BP][BP] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [WP][WP][WP][WP][WP][WP][WP][WP] 2
1 [WR][WN][WB][WQ][WK][WB][WN][WR] 1
    A   B   C   D   E   F   G   H 

# WP A2 A3
# BP A7 A6

    A   B   C   D   E   F   G   H  
8 [BR][BN][BB][BQ][BK][BB][BN][BR] 8
7 [  ][BP][BP][BP][BP][BP][BP][BP] 7
6 [BP][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [WP][  ][  ][  ][  ][  ][  ][  ] 3
2 [  ][WP][WP][WP][WP][WP][WP][WP] 2
1 [WR][WN][WB][WQ][WK][WB][WN][WR] 1
    A   B   C   D   E   F   G   H 
`,

		`
    A   B   C   D   E   F   G   H  
8 [BR][BN][BB][BQ][BK][BB][BN][BR] 8
7 [BP][BP][BP][BP][BP][BP][BP][BP] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [WP][WP][WP][WP][WP][WP][WP][WP] 2
1 [WR][WN][WB][WQ][WK][WB][WN][WR] 1
    A   B   C   D   E   F   G   H 

# BP A7 A6

    A   B   C   D   E   F   G   H  
8 [BR][BN][BB][BQ][BK][BB][BN][BR] 8
7 [  ][BP][BP][BP][BP][BP][BP][BP] 7
6 [BP][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [WP][WP][WP][WP][WP][WP][WP][WP] 2
1 [WR][WN][WB][WQ][WK][WB][WN][WR] 1
    A   B   C   D   E   F   G   H 
`,
	}
}

func TestLegalMoves(t *testing.T) {
	counter := 0
	testCaseRaws := generate()

	for _, testCaseRaw := range testCaseRaws {

		testCase, err := generateMoveTestCase(testCaseRaw, counter)
		counter++

		t.Run(strconv.Itoa(testCase.Number), func(t *testing.T) {
			if err != nil {
				t.Errorf("test setup error %s", err)
			}

			for _, move := range testCase.Moves {
				_, err := move.Piece.Move(move.Current, move.Target, testCase.InitialBoard)
				if err != nil {
					t.Errorf("wanted no error got %s", err)
				}
			}

			if !reflect.DeepEqual(testCase.InitialBoard, testCase.WantedBoard) {
				t.Errorf(" \n wanted: \n%s \n got: \n%s \n",
					testCase.WantedBoard.Print(nil),
					testCase.InitialBoard.Print(nil))
			}
		})
	}
}
