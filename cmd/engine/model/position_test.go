package model

import (
	"reflect"
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

func Test_Equality(t *testing.T) {
	p1 := PositionFromString("A1")
	p2 := PositionFromString("A1")

	if !reflect.DeepEqual(p1, p2) {
		t.Errorf("should be equal")
	}
}
