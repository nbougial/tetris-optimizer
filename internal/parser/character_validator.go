package parser

import "errors"

var errInvalidBlockCharacter = errors.New("invalid block character")

// ValidateBlockCharacters rejects any cell that is not one of the two legal tetromino symbols.
func ValidateBlockCharacters(blocks [][]string) error {
	for _, block := range blocks {
		for _, line := range block {
			for _, cell := range line {
				if cell != '#' && cell != '.' {
					return errInvalidBlockCharacter
				}
			}
		}
	}

	return nil
}
