package solver

import "tetris-optimizer/internal/model"

// CanPlace checks whether a normalized piece fits at the given board anchor without collisions.
func CanPlace(board model.Board, piece model.Piece, anchorRow, anchorCol int) bool {
	for _, cell := range piece.Cells {
		row := anchorRow + cell.Row
		col := anchorCol + cell.Col

		if !board.InBounds(row, col) {
			return false
		}

		if board.Get(row, col) != model.EmptyCell {
			return false
		}
	}

	return true
}
