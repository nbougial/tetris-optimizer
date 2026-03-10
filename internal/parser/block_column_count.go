package parser

import "errors"

var errInvalidBlockColumnCount = errors.New("invalid block column count")

func ValidateBlockColumnCounts(blocks [][]string) error {
	for _, block := range blocks {
		for _, line := range block {
			if len(line) != 4 {
				return errInvalidBlockColumnCount
			}
		}
	}

	return nil
}
