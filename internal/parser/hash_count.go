package parser

import "errors"

var errInvalidHashCount = errors.New("invalid hash count")

func ValidateBlockHashCounts(blocks [][]string) error {
	for _, block := range blocks {
		hashCount := 0
		for _, line := range block {
			for _, cell := range line {
				if cell == '#' {
					hashCount++
				}
			}
		}

		if hashCount != 4 {
			return errInvalidHashCount
		}
	}

	return nil
}
