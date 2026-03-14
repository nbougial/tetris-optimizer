package solver

import (
	"testing"

	"tetris-optimizer/internal/model"
)

func TestPlacePieceWritesLabelToBoard(t *testing.T) {
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

	PlacePiece(&board, piece, 1, 1)

	if board.Get(1, 1) != 'A' || board.Get(1, 2) != 'A' || board.Get(2, 1) != 'A' || board.Get(2, 2) != 'A' {
		t.Fatal("expected placed cells to contain piece label")
	}
}

func TestRemovePieceRestoresEmptyCells(t *testing.T) {
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

	PlacePiece(&board, piece, 1, 1)
	RemovePiece(&board, piece, 1, 1)

	if board.Get(1, 1) != model.EmptyCell || board.Get(1, 2) != model.EmptyCell || board.Get(2, 1) != model.EmptyCell || board.Get(2, 2) != model.EmptyCell {
		t.Fatal("expected removed cells to be empty")
	}
}

func TestPlaceThenRemoveRestoresBoardState(t *testing.T) {
	board := model.NewBoard(4)
	board.Set(0, 0, 'X')
	original := board.Clone()

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

	PlacePiece(&board, piece, 1, 1)
	RemovePiece(&board, piece, 1, 1)

	if board.String() != original.String() {
		t.Fatalf("expected board state to be restored, got %q want %q", board.String(), original.String())
	}
}
