package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//part1()
	part2()
	os.Exit(0)
}

func part1() {
	expenseReport, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Unable to open the file: %v", err)
	}
	defer expenseReport.Close()

	diffMap := getDiffMap(expenseReport, 2020)

	answer := findAnswer(diffMap)

	fmt.Println(answer)
}

func part2() {

	expenseReport, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Unable to open the file: %v", err)
	}
	defer expenseReport.Close()

	var nums []int
	scanner := bufio.NewScanner(expenseReport)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		nums = append(nums, num)
	}

	var answer int
	for i, val := range nums {
		for j := i + 1; j < len(nums); j++ {
			sum := val + nums[j]
			if sum < 2020 {
				for k := j + 1; k < len(nums); k++ {
					total := sum + nums[k]
					if total == 2020 {
						answer = val * nums[j] * nums[k]
						break
					}
				}
			}
		}
	}
	fmt.Println("answer: ", answer)

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