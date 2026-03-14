package output

import "tetris-optimizer/internal/model"

func RenderBoard(board model.Board) string {
	return board.String() + "\n"
}
