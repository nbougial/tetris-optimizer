package main

import (
	"errors"
	"fmt"

	"tetris-optimizer/internal/parser"
)

var errInvalidArgCount = errors.New("invalid argument count")

func execute(args []string) {
	if err := run(args); err != nil {
		fmt.Println("ERROR")
	}
}

func run(args []string) error {
	if len(args) != 1 {
		return errInvalidArgCount
	}

	content, err := parser.ReadInputFile(args[0])
	if err != nil {
		return err
	}

	blocks, err := parser.SplitRawBlocks(content)
	if err != nil {
		return err
	}

	if err := parser.ValidateBlockLineCounts(blocks); err != nil {
		return err
	}

	if err := parser.ValidateBlockColumnCounts(blocks); err != nil {
		return err
	}

	if err := parser.ValidateBlockCharacters(blocks); err != nil {
		return err
	}

	if err := parser.ValidateBlockHashCounts(blocks); err != nil {
		return err
	}

	return nil
}
