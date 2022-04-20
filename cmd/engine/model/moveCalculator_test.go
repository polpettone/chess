package model

import (
	"testing"
)

func TestIsDiagonalMove(t *testing.T) {

	tests := []struct {
		name string
		from Pos
		to   Pos
		want bool
	}{

		{
			from: *PositionFromString("A1"),
			to:   *PositionFromString("B2"),
			want: true,
		},
		{
			from: *PositionFromString("A1"),
			to:   *PositionFromString("A2"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isDiagonalMove(tt.from, tt.to)
			if actual != tt.want {
				t.Errorf("wanted %t got %t", tt.want, actual)
			}

		})
	}
}

func TestIsDiagonalMoveDistanceOne(t *testing.T) {
	tests := []struct {
		name string
		from Pos
		to   Pos
		want bool
	}{
		{
			from: *PositionFromString("A1"),
			to:   *PositionFromString("B2"),
			want: true,
		},
		{
			from: *PositionFromString("A1"),
			to:   *PositionFromString("C3"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isDiagonalMoveDistanceOne(tt.from, tt.to)
			if actual != tt.want {
				t.Errorf("wanted %t got %t", tt.want, actual)
			}

		})
	}
}
