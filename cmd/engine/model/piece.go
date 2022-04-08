package model

import (
	"github.com/polpettone/chess/cmd/engine/model/foo"
	piecePackage "github.com/polpettone/chess/cmd/engine/model/piece"
)

type Piece interface {
	CheckMoveAllowed(current, target foo.Pos) (bool, error)
	GetColor() piecePackage.Color
	GetSymbol() string
}
