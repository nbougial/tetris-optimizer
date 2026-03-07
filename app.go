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

	if _, err := parser.ReadInputFile(args[0]); err != nil {
		return err
	}

	return nil
}
