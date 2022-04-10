package tests

import (
	"github.com/polpettone/chess/cmd/engine/model/foo"
	"testing"
)

func Test_PositionFromString(t *testing.T) {
	raw := "H8"
	pos := foo.PositionFromString(raw)
	if pos.X != 7 {
		t.Errorf("error")
	}
	if pos.Y != 7 {
		t.Errorf("error")
	}
}
