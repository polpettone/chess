package model

import (
	"reflect"
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

func TestGetAllPositionsBetween(t *testing.T) {
	tests := []struct {
		name string
		from Pos
		to   Pos
		want []Pos
	}{

		{
			from: *PositionFromString("D1"),
			to:   *PositionFromString("G4"),
			want: []Pos{
				*PositionFromString("E2"),
				*PositionFromString("F3"),
			},
		},
		{
			from: *PositionFromString("A1"),
			to:   *PositionFromString("B2"),
			want: []Pos{},
		},

		{
			from: *PositionFromString("A1"),
			to:   *PositionFromString("A3"),
			want: []Pos{*PositionFromString("A2")},
		},

		{
			from: *PositionFromString("A1"),
			to:   *PositionFromString("A4"),
			want: []Pos{*PositionFromString("A2"),
				*PositionFromString("A3")},
		},

		{
			from: *PositionFromString("A4"),
			to:   *PositionFromString("A1"),
			want: []Pos{*PositionFromString("A2"),
				*PositionFromString("A3")},
		},

		{
			from: *PositionFromString("A1"),
			to:   *PositionFromString("D1"),
			want: []Pos{*PositionFromString("B1"),
				*PositionFromString("C1")},
		},

		{
			from: *PositionFromString("D1"),
			to:   *PositionFromString("A1"),
			want: []Pos{*PositionFromString("B1"),
				*PositionFromString("C1")},
		},

		{
			from: *PositionFromString("A1"),
			to:   *PositionFromString("D4"),
			want: []Pos{*PositionFromString("B2"),
				*PositionFromString("C3")},
		},

		{
			from: *PositionFromString("D4"),
			to:   *PositionFromString("A1"),
			want: []Pos{*PositionFromString("B2"),
				*PositionFromString("C3")},
		},

		{
			from: *PositionFromString("H1"),
			to:   *PositionFromString("E4"),
			want: []Pos{
				*PositionFromString("F3"),
				*PositionFromString("G2"),
			},
		},

		{
			from: *PositionFromString("E4"),
			to:   *PositionFromString("H1"),
			want: []Pos{
				*PositionFromString("F3"),
				*PositionFromString("G2"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := getAllPositionsBetween(tt.from, tt.to)
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("wanted %v got %v", tt.want, actual)
			}

		})
	}

}
