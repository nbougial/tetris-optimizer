# tetris-optimizer

Command-line Go program that reads tetromino definitions from a text file, validates them, and prints a smallest-square solution.

## Usage

```sh
go run . <path-to-input-file>
```

Example:

```sh
go run . testdata/valid_two_pieces.txt
```

Output:

```text
AAB.
AAB.
..B.
..B.
```

## Input Format

- The program accepts exactly one file path argument.
- Each tetromino is defined as a 4x4 block.
- Only `#` and `.` are allowed.
- Tetrominoes are separated by exactly one blank line.
- Each tetromino must contain exactly four `#` cells.
- The `#` cells must be edge-connected.

Example input:

```text
##..
##..
....
....

#...
#...
#...
#...
```

## Error Behavior

The program prints exactly `ERROR` when:

- the argument count is invalid
- the file cannot be opened or read
- the input format is invalid
- a tetromino has the wrong number of `#` cells
- a tetromino is disconnected

Successful runs print only the solved board.
