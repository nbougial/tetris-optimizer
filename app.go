package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"tetris-optimizer/internal/parser"
)

var errInvalidArgCount = errors.New("invalid argument count")
const errorOutput = "ERROR"

func execute(args []string) {
	executeWithWriter(args, os.Stdout)
}

func executeWithWriter(args []string, writer io.Writer) {
	if err := run(args); err != nil {
		writeError(writer)
	}
}

func writeError(writer io.Writer) {
	fmt.Fprintln(writer, errorOutput)
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

	if err := parser.ValidateBlockConnectivity(blocks); err != nil {
		return err
	}

	if _, err := parser.BuildPieces(blocks); err != nil {
		return err
	}

	return nil
}
