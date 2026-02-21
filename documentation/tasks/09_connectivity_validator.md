# Task 09: Connectivity Validator

## Goal
Ensure tetromino `#` cells are edge-connected.

## Steps
1. Build adjacency via up/down/left/right.
2. Run DFS/BFS from first `#`.
3. Reject if visited `#` count is not 4.

## Done Criteria
- Disconnected shapes produce `ERROR`.

## Depends On
- Task 08.
