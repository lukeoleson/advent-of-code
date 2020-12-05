package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	part1()
	part2()
}

func part1() {

	// open the file
	expenseReport, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Unable to open the file: %v", err)
	}
	defer expenseReport.Close()

	//
	// Build a map where the value is the value pulled from the list of
	// values, and the key is the value we need to make this equal 2020
	//
	diffMap := make(map[int]int)
	scanner := bufio.NewScanner(expenseReport)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		diffMap[2020 - num] = num
	}

	//
	// Plug the values into the map as keys if you find a key that works,
	// then those two numbers must add to 2020
	//
	var answer int
	for _, key := range diffMap {
		val, match := diffMap[key]
		if match == true {
			answer = key * val
			break
		}
	}

	fmt.Println("Part I: ", answer)
}

func part2() {

	// open the file
	expenseReport, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Unable to open the file: %v", err)
	}
	defer expenseReport.Close()

	// Get all the numbers
	var nums []int
	scanner := bufio.NewScanner(expenseReport)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		nums = append(nums, num)
	}

	// loop through everything until you hit the one you want
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
	fmt.Println("Part II: ", answer)

}