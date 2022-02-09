package engine

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParseBoardFromString(t *testing.T) {

	tests := []struct {
		name      string
		input     string
		want      Board
		wantedErr error
	}{

		{
			name: "initial board",
			input: ` 
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
`,
			want:      NewBoard(),
			wantedErr: nil,
		},

		{
			name: "emptyBoard",
			input: ` 
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
			want:      *NewEmptyBoard(),
			wantedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := NewBoardFromString(tt.input)

			if tt.wantedErr == nil && err != nil {
				t.Errorf("wanted no error got %s", err)
			}

			if !reflect.DeepEqual(actual, &tt.want) {
				t.Errorf(" \n wanted: \n%s \n got: \n%s \n",
					tt.want.Print(nil),
					actual.Print(nil))
			}
		})

	}
}

func TestBoard(t *testing.T) {

	tests := []struct {
		name string
		want int
		err  error
	}{
		{
			name: "test number of squares",
			want: 64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := NewBoard()
			squareCount := len(board.Fields)
			if squareCount != tt.want {
				t.Errorf("wanted %d got %d", tt.want, squareCount)
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

const board = ` 
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
`

func DebugParsing(t *testing.T) {
	slice := strings.Split(board, "\n")
	if len(slice) != 12 {
		t.Errorf("wanted 12 got %d", len(slice))
	}

	fmt.Println("Board Parsing")

	y := 7

	emptyBoard := NewEmptyBoard()

	for _, line := range slice {
		if strings.Contains(line, "[") {
			lineSlice := strings.Split(line, "[")
			x := 0
			for _, l := range lineSlice {
				fmt.Println(l)
				if strings.Contains(l, "]") {
					pieceSymbol := l[0:2]
					fmt.Printf("p: %s %d (%d,%d)\n", pieceSymbol, len(pieceSymbol), x, y)

					emptyBoard.SetPieceAtPos(*NewPos(x, y), PieceFrom(pieceSymbol))

					x = x + 1
				}
			}
			y = y - 1
		}
	}

	fmt.Println(emptyBoard.Print(nil))
}
