package solver

import "tetris-optimizer/internal/model"

// PlacePiece writes a piece label onto every translated cell it occupies.
func PlacePiece(board *model.Board, piece model.Piece, anchorRow, anchorCol int) {
	for _, cell := range piece.Cells {
		board.Set(anchorRow+cell.Row, anchorCol+cell.Col, piece.Label)
	}
}

// RemovePiece clears every translated cell that was previously written by PlacePiece.
func RemovePiece(board *model.Board, piece model.Piece, anchorRow, anchorCol int) {
	for _, cell := range piece.Cells {
		board.Set(anchorRow+cell.Row, anchorCol+cell.Col, model.EmptyCell)
	}
}
