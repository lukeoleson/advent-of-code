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

	diffMap := getValuesMap(expenseReport)

	answer := findAnswer(diffMap)

	fmt.Println(answer)

	os.Exit(0)
}

func getValuesMap(f *os.File) map[int]int {

	valMap := make(map[int]int)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		valMap[2020 - num] = num
	}

	return valMap
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