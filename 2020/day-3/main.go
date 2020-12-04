package main

import (
	"./input"
	"fmt"
)

func main() {

	// Right 1, down 1 -> 66
	// Right 3, down 1 -> 200
	// Right 5, down 1 -> 76
	// Right 7, down 1 -> 81
	// Right 1, down 2 -> 46

	run := 1
	rise := 2

	grid := input.Input
	fmt.Println("grid length:", len(grid))
	var trees int
	var col int
	cols := len(grid[0])
	fmt.Println("cols:", cols)

	for i := rise; i < len(grid); i += rise {
		col += run
		if col >= cols {
			col -= cols
		}
		row := grid[i]
		//fmt.Printf("row: %v, col: %v: \n", i, position)
		if row[col] == '#' {
			trees++
		}
	}
	fmt.Println("answer: ", trees)
}