package solver

import (
	"testing"

	"tetris-optimizer/internal/model"
)

func TestCanPlaceReturnsTrueForValidPlacement(t *testing.T) {
	board := model.NewBoard(4)
	piece := model.Piece{
		Cells: []model.Coordinate{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		LabelIndex: 0,
		Label:      'A',
	}

	if !CanPlace(board, piece, 1, 1) {
		t.Fatal("expected piece to fit at 1,1")
	}
}

func TestCanPlaceReturnsFalseWhenOutOfBounds(t *testing.T) {
	board := model.NewBoard(4)
	piece := model.Piece{
		Cells: []model.Coordinate{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		LabelIndex: 0,
		Label:      'A',
	}

	if CanPlace(board, piece, 3, 3) {
		t.Fatal("expected piece to be out of bounds at 3,3")
	}
}

func TestCanPlaceReturnsFalseWhenCellsOccupied(t *testing.T) {
	board := model.NewBoard(4)
	board.Set(1, 2, 'X')

	piece := model.Piece{
		Cells: []model.Coordinate{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		LabelIndex: 0,
		Label:      'A',
	}

	if CanPlace(board, piece, 1, 1) {
		t.Fatal("expected piece placement to collide with occupied cell")
	}
}
