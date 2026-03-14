package solver

import (
	"testing"

	"tetris-optimizer/internal/model"
)

func TestSolveMinimalProducesMinimalBoardForThreeSquares(t *testing.T) {
	pieces := []model.Piece{
		squarePiece('A', 0),
		squarePiece('B', 1),
		squarePiece('C', 2),
	}

	board, solved := SolveMinimal(pieces)
	if !solved {
		t.Fatal("expected solver to find a solution")
	}

	if board.Size != 4 {
		t.Fatalf("expected minimal board size 4, got %d", board.Size)
	}
}

func TestSolveMinimalPlacesEachLabelExactlyFourTimes(t *testing.T) {
	pieces := []model.Piece{
		squarePiece('A', 0),
		squarePiece('B', 1),
		horizontalLinePiece('C', 2),
	}

	board, solved := SolveMinimal(pieces)
	if !solved {
		t.Fatal("expected solver to find a solution")
	}

	for _, piece := range pieces {
		if got := countLabel(board, piece.Label); got != 4 {
			t.Fatalf("expected label %q to appear 4 times, got %d", piece.Label, got)
		}
	}
}

func squarePiece(label rune, index int) model.Piece {
	return model.Piece{
		Cells: []model.Coordinate{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 1, Col: 0},
			{Row: 1, Col: 1},
		},
		LabelIndex: index,
		Label:      label,
	}
}

func horizontalLinePiece(label rune, index int) model.Piece {
	return model.Piece{
		Cells: []model.Coordinate{
			{Row: 0, Col: 0},
			{Row: 0, Col: 1},
			{Row: 0, Col: 2},
			{Row: 0, Col: 3},
		},
		LabelIndex: index,
		Label:      label,
	}
}

func countLabel(board model.Board, label rune) int {
	count := 0

	for row := 0; row < board.Size; row++ {
		for col := 0; col < board.Size; col++ {
			if board.Get(row, col) == label {
				count++
			}
		}
	}

	return count
}
