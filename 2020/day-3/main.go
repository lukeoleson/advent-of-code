package main

import (
	"./input"
	"fmt"
)

func main() {

	rise, run := 2, 1
	grid := input.Input
	var trees int
	var col int
	cols := len(grid[0])

	for row := rise; row < len(grid); row += rise {
		col += run
		if col >= cols {
			col -= cols
		}
		row := grid[row]
		if row[col] == '#' {
			trees++
		}
	}
	fmt.Printf("Solution for Rise: %v, Run: %v: %v\n", rise, run, trees)
}