package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	expenseReport, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Unable to open the file: %v", err)
	}
	defer expenseReport.Close()

	diffMap := getDiffMap(expenseReport, 2020)

	answer := findAnswer(diffMap)

	fmt.Println(answer)

	os.Exit(0)
}

func getDiffMap(f *os.File, year int) map[int]int {

	diffMap := make(map[int]int)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		diffMap[year - num] = num
	}

	return diffMap
}

func findAnswer(diffMap map[int]int) int {

	for _, val := range diffMap {
		val2, match := diffMap[val]
		if match == true {
			return val * val2
		}
	}

	return 0
}