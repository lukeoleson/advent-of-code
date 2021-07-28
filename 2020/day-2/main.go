package main

import (
	"fmt"
	"strconv"
	"strings"
	"./input"
)

func main() {

	//part1()
	part2()
}

func part1() {

	var validPassports int
	for _, str := range input.Input {

		subStrings := strings.Split(str, " ")

		ranges := strings.Split(subStrings[0], "-")
		low, err := strconv.Atoi(ranges[0])
		if err != nil {
		}
		high, err := strconv.Atoi(ranges[1])
		if err != nil {
		}

		requiredChar := subStrings[1][0:1]

		pass := subStrings[2]

		count := strings.Count(pass, requiredChar)

		if count >= low && count <= high {
			validPassports++
		}
	}
	fmt.Println("Part 1: ", validPassports)

}

/*
	for every slice of strings
		split on space
			for the first element
				split on hyphen
				convert these two to integers
			for the second element
				capture it
			Use these indices to get the chars at these positions (non-zero indexed)
			If exactly one of these is our character
				count this as valid
	print the count
*/

func part2() {

	var validPassports int
	for _, str := range input.Input {

		subStrings := strings.Split(str, " ")

		indices := strings.Split(subStrings[0], "-")
		first, _ := strconv.Atoi(indices[0])
		second, _ := strconv.Atoi(indices[1])

		requiredChar := subStrings[1]

		pass := subStrings[2]

			if pass[first - 1] != pass[second - 1] {
				if pass[first - 1] == requiredChar[0]  && pass[second - 2] == requiredChar[0] {
					validPassports++
				}
			}


	}
	fmt.Println("Part 2: ", validPassports)

}