package output

import (
	"testing"

	"tetris-optimizer/internal/model"
)

func TestRenderBoardAddsTrailingNewline(t *testing.T) {
	board := model.NewBoard(2)
	board.Set(0, 0, 'A')
	board.Set(1, 1, 'B')

	got := RenderBoard(board)
	want := "A.\n.B\n"

	if got != want {
		t.Fatalf("unexpected rendered board: got %q want %q", got, want)
	}
}

func TestRenderBoardPreservesEmptyCells(t *testing.T) {
	board := model.NewBoard(2)

	got := RenderBoard(board)
	want := "..\n..\n"

	if got != want {
		t.Fatalf("unexpected empty board rendering: got %q want %q", got, want)
	}
}
