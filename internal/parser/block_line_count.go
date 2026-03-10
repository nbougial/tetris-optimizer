package parser

import "errors"

var errInvalidBlockLineCount = errors.New("invalid block line count")

func ValidateBlockLineCounts(blocks [][]string) error {
	for _, block := range blocks {
		if len(block) != 4 {
			return errInvalidBlockLineCount
		}
	}

	return nil
}
