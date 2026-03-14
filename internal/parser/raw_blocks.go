package parser

import (
	"errors"
	"strings"
)

var errMalformedSeparators = errors.New("malformed block separators")
var errNoBlocks = errors.New("no blocks found")

// SplitRawBlocks breaks the file into ordered 4x4 text blocks separated by one blank line.
func SplitRawBlocks(content []byte) ([][]string, error) {
	normalized := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(normalized, "\n")

	// Ignore one final trailing newline so normal text files do not create an empty block.
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var blocks [][]string
	var current []string

	for _, line := range lines {
		if line == "" {
			// A separator is only valid after collecting a non-empty block.
			if len(current) == 0 {
				return nil, errMalformedSeparators
			}
			blocks = append(blocks, current)
			current = nil
			continue
		}
		current = append(current, line)
	}

	// Ending on a separator would create an empty trailing block, which is invalid.
	if len(current) == 0 {
		return nil, errMalformedSeparators
	}
	blocks = append(blocks, current)

	if len(blocks) == 0 {
		return nil, errNoBlocks
	}

	return blocks, nil
}
