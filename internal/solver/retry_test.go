package solver

import (
	"testing"

	"tetris-optimizer/internal/model"
)

func TestSolveMinimalUsesInitialSizeWhenItFits(t *testing.T) {
	pieces := []model.Piece{
		{
			Cells: []model.Coordinate{
				{Row: 0, Col: 0},
				{Row: 0, Col: 1},
				{Row: 1, Col: 0},
				{Row: 1, Col: 1},
			},
			LabelIndex: 0,
			Label:      'A',
		},
	}

	board, solved := SolveMinimal(pieces)
	if !solved {
		t.Fatal("expected solver to find a solution")
	}

	if board.Size != 2 {
		t.Fatalf("expected minimal board size 2, got %d", board.Size)
	}
}

func TestSolveMinimalRetriesWithLargerBoardWhenNeeded(t *testing.T) {
	pieces := []model.Piece{
		{
			Cells: []model.Coordinate{
				{Row: 0, Col: 0},
				{Row: 0, Col: 1},
				{Row: 0, Col: 2},
				{Row: 0, Col: 3},
			},
			LabelIndex: 0,
			Label:      'A',
		},
		{
			Cells: []model.Coordinate{
				{Row: 0, Col: 0},
				{Row: 0, Col: 1},
				{Row: 0, Col: 2},
				{Row: 0, Col: 3},
			},
			LabelIndex: 1,
			Label:      'B',
		},
	}

	board, solved := SolveMinimal(pieces)
	if !solved {
		t.Fatal("expected solver to find a solution after retrying")
	}

	if board.Size != 4 {
		t.Fatalf("expected minimal board size 4 after retry, got %d", board.Size)
	}
}
