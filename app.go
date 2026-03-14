package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"tetris-optimizer/internal/output"
	"tetris-optimizer/internal/parser"
	"tetris-optimizer/internal/solver"
)

var errInvalidArgCount = errors.New("invalid argument count")

const errorOutput = "ERROR"

func execute(args []string) {
	executeWithWriter(args, os.Stdout)
}

func executeWithWriter(args []string, writer io.Writer) {
	rendered, err := run(args)
	if err != nil {
		writeError(writer)
		return
	}

	fmt.Fprint(writer, rendered)
}

func writeError(writer io.Writer) {
	fmt.Fprintln(writer, errorOutput)
}

func run(args []string) (string, error) {
	if len(args) != 1 {
		return "", errInvalidArgCount
	}

	content, err := parser.ReadInputFile(args[0])
	if err != nil {
		return "", err
	}

	blocks, err := parser.SplitRawBlocks(content)
	if err != nil {
		return "", err
	}

	if err := parser.ValidateBlockLineCounts(blocks); err != nil {
		return "", err
	}

	if err := parser.ValidateBlockColumnCounts(blocks); err != nil {
		return "", err
	}

	if err := parser.ValidateBlockCharacters(blocks); err != nil {
		return "", err
	}

	if err := parser.ValidateBlockHashCounts(blocks); err != nil {
		return "", err
	}

	if err := parser.ValidateBlockConnectivity(blocks); err != nil {
		return "", err
	}

	pieces, err := parser.BuildPieces(blocks)
	if err != nil {
		return "", err
	}

	board, solved := solver.SolveMinimal(pieces)
	if !solved {
		return "", errors.New("no solution found")
	}

	return output.RenderBoard(board), nil
}
