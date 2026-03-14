package output

import "tetris-optimizer/internal/model"

// RenderBoard converts a solved board into the exact CLI output format with a trailing newline.
func RenderBoard(board model.Board) string {
	return board.String() + "\n"
}
