package model

import "testing"

func TestNewBoardInitializesEmptyCells(t *testing.T) {
	board := NewBoard(3)

	if board.Size != 3 {
		t.Fatalf("expected size 3, got %d", board.Size)
	}

	for row := 0; row < board.Size; row++ {
		for col := 0; col < board.Size; col++ {
			if board.Get(row, col) != EmptyCell {
				t.Fatalf("expected empty cell at %d,%d", row, col)
			}
		}
	}
}

func TestBoardInBounds(t *testing.T) {
	board := NewBoard(2)

	if !board.InBounds(0, 0) {
		t.Fatal("expected 0,0 to be in bounds")
	}
	if !board.InBounds(1, 1) {
		t.Fatal("expected 1,1 to be in bounds")
	}
	if board.InBounds(-1, 0) {
		t.Fatal("expected negative row to be out of bounds")
	}
	if board.InBounds(0, -1) {
		t.Fatal("expected negative col to be out of bounds")
	}
	if board.InBounds(2, 0) {
		t.Fatal("expected row 2 to be out of bounds")
	}
	if board.InBounds(0, 2) {
		t.Fatal("expected col 2 to be out of bounds")
	}
}

func TestBoardSetAndGet(t *testing.T) {
	board := NewBoard(2)

	board.Set(1, 0, 'A')

	if got := board.Get(1, 0); got != 'A' {
		t.Fatalf("expected A, got %q", got)
	}
}

func TestBoardCloneCopiesCells(t *testing.T) {
	board := NewBoard(2)
	board.Set(0, 0, 'A')

	cloned := board.Clone()
	cloned.Set(0, 0, 'B')

	if got := board.Get(0, 0); got != 'A' {
		t.Fatalf("expected original board to remain A, got %q", got)
	}
	if got := cloned.Get(0, 0); got != 'B' {
		t.Fatalf("expected cloned board to be B, got %q", got)
	}
}

func TestBoardString(t *testing.T) {
	board := NewBoard(2)
	board.Set(0, 1, 'A')
	board.Set(1, 0, 'B')

	if got := board.String(); got != ".A\nB." {
		t.Fatalf("unexpected board string: %q", got)
	}
}
