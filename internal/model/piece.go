package model

import (
	"errors"
	"sort"
)

// Coordinate represents one occupied cell of a tetromino relative to an origin.
type Coordinate struct {
	Row int
	Col int
}

// Piece stores a normalized tetromino shape together with its stable label metadata.
type Piece struct {
	Cells      []Coordinate
	LabelIndex int
	Label      rune
}

var errLabelOverflow = errors.New("label overflow")

// LabelForIndex converts input order into the required A-Z piece labels.
func LabelForIndex(index int) (rune, error) {
	if index < 0 || index >= 26 {
		return 0, errLabelOverflow
	}

	return rune('A' + index), nil
}

// NormalizePiece shifts a piece to the top-left origin and sorts cells for stable comparisons.
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
		Label:      piece.Label,
	}

	// Shift every occupied cell so the piece origin becomes the smallest row/col seen.
	for index, cell := range piece.Cells {
		normalized.Cells[index] = Coordinate{
			Row: cell.Row - minRow,
			Col: cell.Col - minCol,
		}
	}

	// Stable ordering keeps solver behavior and test expectations deterministic.
	sort.Slice(normalized.Cells, func(i, j int) bool {
		if normalized.Cells[i].Row != normalized.Cells[j].Row {
			return normalized.Cells[i].Row < normalized.Cells[j].Row
		}
		return normalized.Cells[i].Col < normalized.Cells[j].Col
	})

	return normalized
}
