package solver

import "tetris-optimizer/internal/model"

func PlacePiece(board *model.Board, piece model.Piece, anchorRow, anchorCol int) {
	for _, cell := range piece.Cells {
		board.Set(anchorRow+cell.Row, anchorCol+cell.Col, piece.Label)
	}
}

func RemovePiece(board *model.Board, piece model.Piece, anchorRow, anchorCol int) {
	for _, cell := range piece.Cells {
		board.Set(anchorRow+cell.Row, anchorCol+cell.Col, model.EmptyCell)
	}
}
