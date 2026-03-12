package model

import "sort"

type Coordinate struct {
	Row int
	Col int
}

type Piece struct {
	Cells      []Coordinate
	LabelIndex int
}

func NormalizePiece(piece Piece) Piece {
	if len(piece.Cells) == 0 {
		return piece
	}

	minRow := piece.Cells[0].Row
	minCol := piece.Cells[0].Col

	for _, cell := range piece.Cells[1:] {
		if cell.Row < minRow {
			minRow = cell.Row
		}
		if cell.Col < minCol {
			minCol = cell.Col
		}
	}

	normalized := Piece{
		Cells:      make([]Coordinate, len(piece.Cells)),
		LabelIndex: piece.LabelIndex,
	}

	for index, cell := range piece.Cells {
		normalized.Cells[index] = Coordinate{
			Row: cell.Row - minRow,
			Col: cell.Col - minCol,
		}
	}

	sort.Slice(normalized.Cells, func(i, j int) bool {
		if normalized.Cells[i].Row != normalized.Cells[j].Row {
			return normalized.Cells[i].Row < normalized.Cells[j].Row
		}
		return normalized.Cells[i].Col < normalized.Cells[j].Col
	})

	return normalized
}
