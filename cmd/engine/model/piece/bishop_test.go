package piece

import (
	"testing"
)

func TestBishopIllegalMoves(t *testing.T) {
	testCasesRaw := `
# WB D8 E6
# WB A8 A1
`
	tests := GeneratePieceMoveTestCases(testCasesRaw)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := tt.Piece.CheckMoveAllowed(tt.Current, tt.Target)
			if result != false {
				t.Errorf("wanted false, got true")
			}
			if err == nil {
				t.Errorf("wanted error, got none")
			}
		})
	}
}

func TestBishopLegalMoves(t *testing.T) {
	testCasesRaw := `
# BB B1 H7
# BB F1 A6
`
	tests := GeneratePieceMoveTestCases(testCasesRaw)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := tt.Piece.CheckMoveAllowed(tt.Current, tt.Target)

			if result != true {
				t.Errorf("wanted true, got false")
			}
			if err != nil {
				t.Errorf("wanted no error, got none %s", err)
			}

		})
	}
}

const boardWithBishops = ` 
    A   B   C   D   E   F   G   H  
8 [WB][  ][  ][WB][  ][  ][  ][  ] 8
7 [  ][  ][  ][  ][  ][  ][  ][  ] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [  ][  ][  ][  ][  ][  ][  ][  ] 2
1 [  ][BB][  ][  ][  ][BB][  ][  ] 1
    A   B   C   D   E   F   G   H 
`
