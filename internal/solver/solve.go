package solver

import "tetris-optimizer/internal/model"

// Solve tries to place all pieces on the current board size without resizing it.
func Solve(board *model.Board, pieces []model.Piece) bool {
	return solveFromIndex(board, pieces, 0)
}

// solveFromIndex performs the recursive backtracking search in input order.
func solveFromIndex(board *model.Board, pieces []model.Piece, index int) bool {
	if index == len(pieces) {
		return true
	}

	piece := pieces[index]

	// Search in row-major order so the first solution stays deterministic across runs.
	for row := 0; row < board.Size; row++ {
		for col := 0; col < board.Size; col++ {
			if !CanPlace(*board, piece, row, col) {
				continue
			}

			PlacePiece(board, piece, row, col)
			if solveFromIndex(board, pieces, index+1) {
				return true
			}
			// Roll back before trying the next candidate position.
			RemovePiece(board, piece, row, col)
		}
	}

	return false
}
