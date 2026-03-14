package parser

import "errors"

var errInvalidBlockColumnCount = errors.New("invalid block column count")

// ValidateBlockColumnCounts ensures every row in every tetromino block is exactly four cells wide.
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
