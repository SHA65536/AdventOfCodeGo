package day01

import (
	"strconv"
)

// https://adventofcode.com/2022/day/1

func MaxCalories(calories string) (string, error) {
	var curTotal, curLine int
	var prev byte
	var max int
	// Looping over characters
	for i := range calories {
		if calories[i] >= '0' && calories[i] <= '9' {
			// If it's a digit, add to current line
			curLine = 10*curLine + int(calories[i]-'0')
		} else if prev != '\n' {
			// If it's one linebreak, add current line to elf total
			curTotal += curLine
			curLine = 0
		} else {
			// If it's double line break, end current elf and check for top
			if curTotal > max {
				max = curTotal
			}
			curTotal = 0
		}
		prev = calories[i]
	}
	curTotal += curLine
	if curTotal > max {
		max = curTotal
	}
	return strconv.Itoa(max), nil
}

func MaxCaloriesThree(calories string) (string, error) {
	var curTotal, curLine int
	var prev byte
	var maxes [3]int
	// Looping over characters
	for i := range calories {
		if calories[i] >= '0' && calories[i] <= '9' {
			// If it's a digit, add to current line
			curLine = 10*curLine + int(calories[i]-'0')
		} else if prev != '\n' {
			// If it's one linebreak, add current line to elf total
			curTotal += curLine
			curLine = 0
		} else {
			// If it's double line break, end current elf and check for top
			UpdateMaxThree(&maxes, curTotal)
			curTotal = 0
		}
		prev = calories[i]
	}
	// Last elf isn't inside the loop
	curTotal += curLine
	UpdateMaxThree(&maxes, curTotal)
	return strconv.Itoa(maxes[0] + maxes[1] + maxes[2]), nil
}

func UpdateMaxThree(arr *[3]int, val int) {
	// Bubbling the value down
	for i := len(arr) - 1; i >= 0; i-- {
		if val > arr[i] {
			val, arr[i] = arr[i], val
		}
	}
}
