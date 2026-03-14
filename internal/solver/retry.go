package solver

import "tetris-optimizer/internal/model"

// SolveMinimal keeps increasing the board size until the first valid solution is found.
func SolveMinimal(pieces []model.Piece) (model.Board, bool) {
	size := InitialBoardSize(len(pieces))

	for {
		board := model.NewBoard(size)
		if Solve(&board, pieces) {
			return board, true
		}
		size++
	}
}
