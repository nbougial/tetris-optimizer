package solver

import "math"

func InitialBoardSize(pieceCount int) int {
	if pieceCount <= 0 {
		return 0
	}

	return int(math.Ceil(math.Sqrt(float64(pieceCount * 4))))
}
