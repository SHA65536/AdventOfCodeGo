package day01

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/1

func MaxCalories(calories string) (string, error) {
	var max, cur int
	lines := strings.Split(calories, "\n")
	// Looping over foods
	for i := range lines {
		// Empty line represents end of elf
		if len(lines[i]) == 0 {
			if cur > max { // Check if the current elf is the highest
				max = cur
			}
			cur = 0
		} else {
			num, err := strconv.Atoi(lines[i])
			if err != nil {
				return "", err
			}
			// Adding current line to elf total
			cur += num
		}
	}
	// Last elf is out of the loop
	if cur > max {
		max = cur
	}
	return strconv.Itoa(max), nil
}

func MaxCaloriesThree(calories string) (string, error) {
	var cur int
	var maxes [3]int
	lines := strings.Split(calories, "\n")
	// Looping over foods
	for i := range lines {
		// Empty line represents end of elf
		if len(lines[i]) == 0 {
			// Check if current elf is among the top 3
			UpdateMaxThree(&maxes, cur)
			cur = 0
		} else {
			num, err := strconv.Atoi(lines[i])
			if err != nil {
				return "", err
			}
			// Adding current line to elf total
			cur += num
		}
	}
	// Last elf is out of the loop
	UpdateMaxThree(&maxes, cur)
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
