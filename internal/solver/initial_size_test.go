package solver

import "testing"

func TestInitialBoardSize(t *testing.T) {
	testCases := []struct {
		pieceCount int
		want      int
	}{
		{pieceCount: 0, want: 0},
		{pieceCount: 1, want: 2},
		{pieceCount: 2, want: 3},
		{pieceCount: 3, want: 4},
		{pieceCount: 4, want: 4},
		{pieceCount: 5, want: 5},
	}

	for _, testCase := range testCases {
		if got := InitialBoardSize(testCase.pieceCount); got != testCase.want {
			t.Fatalf("piece count %d: expected %d, got %d", testCase.pieceCount, testCase.want, got)
		}
	}
}
