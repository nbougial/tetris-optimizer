# Product Requirements Document (PRD)

## Project
Tetris Optimizer (Tetromino Square Assembler)

## 1. Purpose
Build a command-line Go program that reads tetromino definitions from a text file and arranges all tetrominoes into the smallest possible square.

## 2. Problem Statement
Given a list of tetrominoes in a strict text format, the program must compute a valid placement of all pieces, preserving piece order, and output a square map labeled with uppercase letters (`A`, `B`, `C`, ...).

## 3. Goals
- Parse and validate a file containing one or more tetrominoes.
- Reject malformed input with a single `ERROR` output.
- Place all tetrominoes in the smallest possible square.
- Output a human-readable solved square using piece labels and `.` for empty cells.

## 4. Non-Goals
- GUI or web interface.
- Support for shapes other than tetrominoes.
- External dependencies beyond Go standard library.

## 5. Users
- Students and evaluators running the project from terminal.
- Developers validating algorithmic and parsing correctness.

## 6. Functional Requirements
1. Program accepts exactly one CLI argument: path to input text file.
2. Program reads all tetrominoes from the file.
3. Program validates file structure:
- Each tetromino is a 4x4 grid of `#` and `.`.
- Tetrominoes are separated by exactly one blank line.
- At least one tetromino must be present.
4. Program validates tetromino correctness:
- Exactly 4 `#` blocks per tetromino.
- Blocks must form a connected shape by edges (not diagonals).
5. On invalid file path, read failure, or invalid format/content: print `ERROR`.
6. On valid input, solver must place all tetrominoes.
7. Output map uses:
- `A` for first tetromino, `B` for second, etc.
- `.` for empty cells.
8. Output must be the smallest square dimension that can contain all pieces.

## 7. Input/Output Specification
### Input
- CLI: `go run . <path-to-file>`
- File format example:

```txt
#...
#...
#...
#...

....
....
..##
..##
```

### Output
- On success: solved square printed line-by-line.
- On failure: exactly `ERROR`.

## 8. Constraints
- Language: Go.
- Allowed packages: Go standard library only.
- Follow clean code and good engineering practices.

## 9. Quality Requirements
- Deterministic output for the same input.
- Clear error handling path.
- Modular package structure (parsing, validation, solving, printing).
- Unit tests recommended for parser, validator, and solver behavior.

## 10. Suggested Technical Approach
1. Parse file into raw 4x4 blocks.
2. Normalize tetromino coordinates (top-left shift).
3. Validate connectivity and block count.
4. Start board size at `ceil(sqrt(4 * nPieces))`.
5. Use backtracking to place tetrominoes in order.
6. If unsolved, increment board size and retry.
7. Print first valid solution for smallest size.

## 11. Edge Cases
- Empty file.
- Missing/extra lines in tetromino blocks.
- Invalid characters.
- Wrong separator usage.
- Disconnected `#` cells.
- More than 26 tetrominoes (label overflow decision should be defined; recommended: return `ERROR`).

## 12. Acceptance Criteria
- `go build` completes successfully.
- Invalid inputs always produce `ERROR`.
- Valid inputs produce a square containing all tetrominoes.
- Output uses uppercase labels by input order.
- Solver returns minimal possible square size.
- Project includes or is ready for unit tests.

## 13. Test Plan (Minimum)
- Parser valid/invalid format cases.
- Tetromino validation cases (count/connectivity).
- Small known solvable files with expected output size.
- Unsolvable-at-size-X but solvable-at-size-(X+1) scenarios.
- CLI argument and file open error paths.
