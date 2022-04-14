package model

import (
	"testing"
)

func Test_PositionFromString(t *testing.T) {
	raw := "H8"
	pos := PositionFromString(raw)
	if pos.X != 7 {
		t.Errorf("error")
	}
	if pos.Y != 7 {
		t.Errorf("error")
	}
}

func Test_Print(t *testing.T) {
	raw := "H8"
	pos := PositionFromString(raw)
	if pos.Print() != raw {
		t.Errorf("wanted %s got %s", raw, pos.Print())
	}
}
