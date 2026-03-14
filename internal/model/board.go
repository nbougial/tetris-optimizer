package model

import "strings"

// EmptyCell is the rune used for every unoccupied board position.
const EmptyCell = '.'

// Board stores a square grid and its size so solver code can work with one value.
type Board struct {
	Size  int
	Cells [][]rune
}

// NewBoard allocates a square board and fills every position with EmptyCell.
func NewBoard(size int) Board {
	board := Board{
		Size:  size,
		Cells: make([][]rune, size),
	}

	for row := 0; row < size; row++ {
		board.Cells[row] = make([]rune, size)
		for col := 0; col < size; col++ {
			board.Cells[row][col] = EmptyCell
		}
	}

	return board
}

// InBounds reports whether a translated cell lands inside the board.
func (b Board) InBounds(row, col int) bool {
	return row >= 0 && row < b.Size && col >= 0 && col < b.Size
}

// Get reads a single board cell.
func (b Board) Get(row, col int) rune {
	return b.Cells[row][col]
}

// Set writes a single board cell.
func (b *Board) Set(row, col int, value rune) {
	b.Cells[row][col] = value
}

// Clone copies the board grid so tests and backtracking logic can compare states safely.
func (b Board) Clone() Board {
	cloned := Board{
		Size:  b.Size,
		Cells: make([][]rune, b.Size),
	}

	for row := 0; row < b.Size; row++ {
		cloned.Cells[row] = make([]rune, b.Size)
		copy(cloned.Cells[row], b.Cells[row])
	}

	return cloned
}

// String renders the board exactly as the CLI expects it, without a trailing newline.
func (b Board) String() string {
	rows := make([]string, b.Size)

	for row := 0; row < b.Size; row++ {
		rows[row] = string(b.Cells[row])
	}

	return strings.Join(rows, "\n")
}
