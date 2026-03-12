package parser

import "tetris-optimizer/internal/model"

func BuildPieces(blocks [][]string) ([]model.Piece, error) {
	pieces := make([]model.Piece, 0, len(blocks))

	for index, block := range blocks {
		label, err := model.LabelForIndex(index)
		if err != nil {
			return nil, err
		}

		piece := model.Piece{
			Cells:      make([]model.Coordinate, 0, 4),
			LabelIndex: index,
			Label:      label,
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

		pieces = append(pieces, model.NormalizePiece(piece))
	}

	return pieces, nil
}
