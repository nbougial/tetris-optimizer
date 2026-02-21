# Task 21: Unified Error Output

## Goal
Centralize failure path to print exactly `ERROR`.

## Steps
1. Define application-level error policy.
2. Avoid leaking internal details on stdout.
3. Ensure newline behavior is consistent.

## Done Criteria
- Any invalid input/path scenario prints only `ERROR`.

## Depends On
- Task 03.
