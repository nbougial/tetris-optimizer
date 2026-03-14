package solver

import (
	"testing"

	"tetris-optimizer/internal/model"
)

func TestSolveFindsValidPlacementWhenOneExists(t *testing.T) {
	board := model.NewBoard(4)
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
		{
			Cells: []model.Coordinate{
				{Row: 0, Col: 0},
				{Row: 0, Col: 1},
				{Row: 1, Col: 0},
				{Row: 1, Col: 1},
			},
			LabelIndex: 1,
			Label:      'B',
		},
	}

	if !Solve(&board, pieces) {
		t.Fatal("expected solver to find a valid placement")
	}

	if board.Get(0, 0) == model.EmptyCell {
		t.Fatal("expected board to contain placed pieces")
	}
}

func TestSolveReturnsFalseAndRestoresBoardWhenNoPlacementExists(t *testing.T) {
	board := model.NewBoard(3)
	original := board.Clone()

	pieces := []model.Piece{
		{
			Cells: []model.Coordinate{
				{Row: 0, Col: 0},
				{Row: 1, Col: 0},
				{Row: 2, Col: 0},
				{Row: 3, Col: 0},
			},
			LabelIndex: 0,
			Label:      'A',
		},
	}

	if Solve(&board, pieces) {
		t.Fatal("expected solver to fail on an undersized board")
	}

	if board.String() != original.String() {
		t.Fatalf("expected failed solve to restore board, got %q want %q", board.String(), original.String())
	}
}
