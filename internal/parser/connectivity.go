package parser

import "errors"

var errDisconnectedBlock = errors.New("disconnected block")

type cell struct {
	row int
	col int
}

func ValidateBlockConnectivity(blocks [][]string) error {
	for _, block := range blocks {
		if !isBlockConnected(block) {
			return errDisconnectedBlock
		}
	}

	return nil
}

func isBlockConnected(block []string) bool {
	start, found := findFirstHash(block)
	if !found {
		return false
	}

	visited := map[cell]bool{start: true}
	queue := []cell{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range neighbors(current) {
			if next.row < 0 || next.row >= len(block) {
				continue
			}
			if next.col < 0 || next.col >= len(block[next.row]) {
				continue
			}
			if block[next.row][next.col] != '#' || visited[next] {
				continue
			}

			visited[next] = true
			queue = append(queue, next)
		}
	}

	return len(visited) == 4
}

func findFirstHash(block []string) (cell, bool) {
	for row, line := range block {
		for col, current := range line {
			if current == '#' {
				return cell{row: row, col: col}, true
			}
		}
	}

	return cell{}, false
}

func neighbors(current cell) []cell {
	return []cell{
		{row: current.row - 1, col: current.col},
		{row: current.row + 1, col: current.col},
		{row: current.row, col: current.col - 1},
		{row: current.row, col: current.col + 1},
	}
}
