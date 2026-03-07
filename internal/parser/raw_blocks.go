package parser

import (
	"errors"
	"strings"
)

var errMalformedSeparators = errors.New("malformed block separators")
var errNoBlocks = errors.New("no blocks found")

func SplitRawBlocks(content []byte) ([][]string, error) {
	normalized := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(normalized, "\n")

	// Ignore final trailing newline only.
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var blocks [][]string
	var current []string

	for _, line := range lines {
		if line == "" {
			if len(current) == 0 {
				return nil, errMalformedSeparators
			}
			blocks = append(blocks, current)
			current = nil
			continue
		}
		current = append(current, line)
	}

	if len(current) == 0 {
		return nil, errMalformedSeparators
	}
	blocks = append(blocks, current)

	if len(blocks) == 0 {
		return nil, errNoBlocks
	}

	return blocks, nil
}
