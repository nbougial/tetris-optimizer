# Task 18: Backtracking Solver Core

## Goal
Implement recursive solver that places pieces in input order.

## Steps
1. Base case: all pieces placed => success.
2. Iterate candidate cells for current piece.
3. Place, recurse, and rollback on failure.

## Done Criteria
- Solver finds a valid placement when one exists.

## Depends On
- Task 17.
