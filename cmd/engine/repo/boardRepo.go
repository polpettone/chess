package repo

import (
	"errors"
	"github.com/polpettone/chess/cmd/engine/model"
	"io/ioutil"
	"os"
)

func SaveBoardToFile(path string, board model.Board) error {
	err := ioutil.WriteFile(path, []byte(board.Print(nil)), 0755)
	return err
}

func LoadBoardFromFileOrCreateNewBoard(path string) (*model.Board, error) {
	_, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return model.NewBoardFromString(newBoard)
	} else {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		return model.NewBoardFromString(string(content))
	}
}

const newBoard string = ` 
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
