package parser

import "os"

func ReadInputFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
