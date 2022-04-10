package tests

import (
	"github.com/polpettone/chess/cmd/engine/model"
	"reflect"
	"testing"
)

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
			board := model.NewBoard()
			squareCount := len(board.Fields)
			if squareCount != tt.want {
				t.Errorf("wanted %d got %d", tt.want, squareCount)
			}
		})
	}

}

func TestParseBoardFromString(t *testing.T) {

	tests := []struct {
		name      string
		input     string
		want      model.Board
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
			want:      model.NewBoard(),
			wantedErr: nil,
		},

		{
			name: "empty board",
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
			want:      *model.NewEmptyBoard(),
			wantedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := model.NewBoardFromString(tt.input)

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
