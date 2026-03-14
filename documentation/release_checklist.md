# Release Checklist

Submission-ready verification completed on 2026-03-14.

## Acceptance Criteria

- [x] `go build ./...` succeeds
- [x] Invalid inputs print exactly `ERROR`
- [x] Valid inputs produce a solved square
- [x] Output uses uppercase labels by input order
- [x] Solver returns a minimal square size
- [x] Parser, validator, solver, and CLI tests are present

## Project Constraints

- [x] Go standard library only
- [x] No external module dependencies in `go.mod`
- [x] Modular package structure in place:
  - `internal/parser`
  - `internal/model`
  - `internal/solver`
  - `internal/output`

## Quality Checks

- [x] `gofmt -w` run on Go files
- [x] `go test ./...` passes
- [x] `go build ./...` passes
- [x] Deterministic output checks added
- [x] Regression tests added for previously fixed parser/shape issues

## Output and Error Policy

- [x] Success path renders solved board text with trailing newline
- [x] Failure path prints only `ERROR`
- [x] No debug-print scan findings from `fmt.Print*`, `log.*`, or `panic(...)`

## Examples and Fixtures

- [x] Top-level `README.md` includes usage, input rules, error behavior, and sample output
- [x] `testdata/README.md` catalogs reusable fixtures
- [x] Valid, invalid-format, invalid-shape, and multi-piece fixtures are present

## Notes

- Current verification is based on the repository state at this checklist date.
- The repository is ready for submission based on the PRD and completed task set.
