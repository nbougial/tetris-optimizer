package solver

import "math"

// InitialBoardSize returns the mathematically smallest square side that could hold all pieces by area.
func InitialBoardSize(pieceCount int) int {
	if pieceCount <= 0 {
		return 0
	}

	return int(math.Ceil(math.Sqrt(float64(pieceCount * 4))))
}
