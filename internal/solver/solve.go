package solver

import "tetris-optimizer/internal/model"

func Solve(board *model.Board, pieces []model.Piece) bool {
	return solveFromIndex(board, pieces, 0)
}

func solveFromIndex(board *model.Board, pieces []model.Piece, index int) bool {
	if index == len(pieces) {
		return true
	}

	piece := pieces[index]

	for row := 0; row < board.Size; row++ {
		for col := 0; col < board.Size; col++ {
			if !CanPlace(*board, piece, row, col) {
				continue
			}

			PlacePiece(board, piece, row, col)
			if solveFromIndex(board, pieces, index+1) {
				return true
			}
			RemovePiece(board, piece, row, col)
		}
	}

	return false
}
