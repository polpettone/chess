package engine

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMovePawn(t *testing.T) {
	tests := []struct {
		name         string
		currentBoard map[string]string
		move         []string
		wantedBoard  map[string]string
		wantedErr    error
	}{

		{
			name: "white pawn 1",
			move: []string{"WP", "A2", "A3"},
			wantedBoard: map[string]string{
				"A2": "",
				"A3": "WP",
			},
			wantedErr: nil,
		},

		{
			name: "white pawn init 2",
			move: []string{"WP", "A2", "A4"},
			wantedBoard: map[string]string{
				"A2": "",
				"A4": "WP",
			},
			wantedErr: nil,
		},

		{
			name: "white pawn 2",
			currentBoard: map[string]string{
				"A2": "",
				"A3": "WP",
			},
			move:        []string{"WP", "A3", "A5"},
			wantedBoard: nil,
			wantedErr:   fmt.Errorf("not allowed"),
		},

		{
			name: "white pawn back",
			currentBoard: map[string]string{
				"A2": "",
				"A3": "WP",
			},
			move:        []string{"WP", "A3", "A2"},
			wantedBoard: nil,
			wantedErr:   fmt.Errorf("not allowed"),
		},

		{
			name: "white pawn strike 1",
			currentBoard: map[string]string{
				"B3": "BP",
				"B7": "",
			},
			move: []string{"WP", "A2", "B3"},
			wantedBoard: map[string]string{
				"A2": "",
				"B3": "WP",
				"B7": "",
			},
			wantedErr: nil,
		},

		{
			name: "white pawn strike own color",
			currentBoard: map[string]string{
				"B2": "",
				"B3": "WP",
			},
			move:      []string{"WP", "A2", "B3"},
			wantedErr: fmt.Errorf("not allowed"),
		},

		{
			name: "black pawn strike own color",
			currentBoard: map[string]string{
				"B7": "",
				"B6": "BP",
			},
			move:      []string{"BP", "A7", "B6"},
			wantedErr: fmt.Errorf("not allowed"),
		},

		{
			name: "black pawn strike 1",
			currentBoard: map[string]string{
				"B6": "WP",
				"B2": "",
			},
			move: []string{"BP", "A7", "B6"},
			wantedBoard: map[string]string{
				"B2": "",
				"B6": "BP",
				"A7": "",
			},
			wantedErr: nil,
		},

		{
			name: "black pawn 1",
			move: []string{"BP", "A7", "A6"},
			wantedBoard: map[string]string{
				"A7": "",
				"A6": "BP",
			},
			wantedErr: nil,
		},

		{
			name: "black pawn init 2",
			move: []string{"BP", "A7", "A5"},
			wantedBoard: map[string]string{
				"A7": "",
				"A5": "BP",
			},
			wantedErr: nil,
		},

		{
			name: "black pawn 2",
			currentBoard: map[string]string{
				"A6": "BP",
				"A7": "",
			},
			move:        []string{"BP", "A6", "A4"},
			wantedBoard: nil,
			wantedErr:   fmt.Errorf("not allowed"),
		},

		{
			name: "black pawn back",
			move: []string{"BP", "A6", "A7"},
			currentBoard: map[string]string{
				"A6": "BP",
				"A7": "",
			},
			wantedBoard: nil,
			wantedErr:   fmt.Errorf("not allowed"),
		},

		{
			name: "pawn frontstep not allowed if piece is in the way",
			move: []string{"WP", "A2", "A3"},
			currentBoard: map[string]string{
				"A7": "",
				"A3": "BP",
			},
			wantedBoard: nil,
			wantedErr:   fmt.Errorf("not allowed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			piece := PieceFrom(tt.move[0])
			board := NewBoard()
			wantedBoard := NewBoard()

			if tt.currentBoard != nil {
				board = changePiecesOnBoard(board, tt.currentBoard)
			}

			wantedBoard = changePiecesOnBoard(wantedBoard, tt.wantedBoard)
			newBoard, err := piece.Move(*P(tt.move[1]), *P(tt.move[2]), board)

			if tt.wantedErr == nil && err != nil {
				t.Errorf("wanted no error got %v", err)
			}

			if tt.wantedErr != nil && err == nil {
				t.Errorf("wanted error '%v' got none", tt.wantedErr)
			}

			if tt.wantedErr == nil {
				if newBoard != nil {
					if !reflect.DeepEqual(*newBoard, wantedBoard) {
						t.Errorf(" \n wanted: \n%s \n got: \n%s \n",
							wantedBoard.Print(nil),
							newBoard.Print(nil))
					}
				} else {
					t.Errorf("board should not be nil")
				}
			}
		})
	}
}
