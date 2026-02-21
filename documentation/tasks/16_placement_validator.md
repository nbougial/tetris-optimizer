# Task 16: Placement Validator

## Goal
Check whether a piece can be placed at board coordinate `(r, c)`.

## Steps
1. Translate normalized piece coordinates by anchor `(r, c)`.
2. Ensure all translated cells are in bounds.
3. Ensure all target cells are currently `.`.

## Done Criteria
- Placement check is correct and deterministic.

## Depends On
- Task 14.
