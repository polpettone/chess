package model

import (
	"github.com/polpettone/chess/cmd/engine/model/foo"
)

type Piece interface {
	CheckMoveAllowed(current, target foo.Pos) (bool, error)
	GetColor() foo.Color
	GetSymbol() string
}
