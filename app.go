package main

import "errors"

var errInvalidArgCount = errors.New("invalid argument count")

func run(args []string) error {
	if len(args) != 1 {
		return errInvalidArgCount
	}

	_ = args[0]
	return nil
}
