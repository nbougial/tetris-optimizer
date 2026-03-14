package solver

import "tetris-optimizer/internal/model"

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
