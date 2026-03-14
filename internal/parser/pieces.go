package parser

import "tetris-optimizer/internal/model"

// BuildPieces converts validated text blocks into normalized model pieces with stable labels.
func BuildPieces(blocks [][]string) ([]model.Piece, error) {
	pieces := make([]model.Piece, 0, len(blocks))

	for index, block := range blocks {
		// Labels are assigned in input order so the solver and renderer stay deterministic.
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

		// The solver only works with normalized coordinates, not the original 4x4 offsets.
		pieces = append(pieces, model.NormalizePiece(piece))
	}

	return pieces, nil
}
