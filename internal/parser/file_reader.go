package parser

import "os"

// ReadInputFile loads the entire source file so later parser stages can validate it strictly.
func ReadInputFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
