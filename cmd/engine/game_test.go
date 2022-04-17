package engine

import (
	"testing"

	"github.com/polpettone/chess/cmd/engine/model"
)

func TestRandomWhitePiece(t *testing.T) {
	piece := randomPiece(model.WHITE)
	if piece.GetColor() != model.WHITE {
		t.Errorf("wanted WHITE, got BLACK")
	}
}

func TestRandomBlackPiece(t *testing.T) {
	piece := randomPiece(model.WHITE)
	if piece.GetColor() != model.WHITE {
		t.Errorf("wanted WHITE, got BLACK")
	}
}
