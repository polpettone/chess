package model

import (
	"reflect"
	"testing"
)

func Test_PrintMovement(t *testing.T) {
	raw := "WP A2 A4"
	movement, _ := MoveFromString(raw)
	actual := movement.Print()
	if actual != raw {
		t.Errorf("actual %s not equal wanted %s", actual, raw)
	}
}

func Test_ValidMoveFromString(t *testing.T) {
	tests := []struct {
		move string
		want Movement
	}{

		{
			move: "WP A2 A4",
			want: Movement{
				Piece: PieceFrom("WP"),
				From:  *PositionFromString("A2"),
				To:    *PositionFromString("A4"),
			},
		},

		{
			move: "BP A2 A4",
			want: Movement{
				Piece: PieceFrom("BP"),
				From:  *PositionFromString("A2"),
				To:    *PositionFromString("A4"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.move, func(t *testing.T) {
			actual, err := MoveFromString(tt.move)
			if err != nil {
				t.Errorf("wanted no error, got %s", err)
				return
			}
			if !reflect.DeepEqual(*actual, tt.want) {
				t.Errorf("%v not equal %v", actual, tt.want)
			}
		})
	}
}

func Test_InvalidMoveFromString(t *testing.T) {
	tests := []struct {
		move string
	}{

		{move: "A4"},
		{move: "A2 A4"},
		{move: "X2 A4"},
		{move: ""},
	}
	for _, tt := range tests {
		t.Run(tt.move, func(t *testing.T) {
			_, err := MoveFromString(tt.move)
			if err == nil {
				t.Errorf("wanted error, got none")
			}
		})
	}
}
