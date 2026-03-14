package main

import "os"

// main stays intentionally small and only forwards CLI arguments into the app layer.
func main() {
	execute(os.Args[1:])
}
