package parser

import "tetris-optimizer/internal/model"

func BuildPieces(blocks [][]string) []model.Piece {
	pieces := make([]model.Piece, 0, len(blocks))

	for index, block := range blocks {
		piece := model.Piece{
			Cells:      make([]model.Coordinate, 0, 4),
			LabelIndex: index,
		}

		for row, line := range block {
			for col, cell := range line {
				if cell == '#' {
					piece.Cells = append(piece.Cells, model.Coordinate{
						Row: row,
						Col: col,
					})
				}
			}
		}

		pieces = append(pieces, piece)
	}

	return pieces
}
