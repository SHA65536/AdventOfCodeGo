package day01

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/1

func MaxCalories(calories string) (string, error) {
	var max, cur int
	lines := strings.Split(calories, "\n")
	for i := range lines {
		if len(lines[i]) == 0 {
			if cur > max {
				max = cur
			}
			cur = 0
		} else {
			num, err := strconv.Atoi(lines[i])
			if err != nil {
				return "", err
			}
			cur += num
		}
	}
	if cur > max {
		max = cur
	}
	return strconv.Itoa(max), nil
}

func MaxCaloriesThree(calories string) (string, error) {
	var cur int
	var maxes [3]int
	lines := strings.Split(calories, "\n")
	for i := range lines {
		if len(lines[i]) == 0 {
			UpdateMaxThree(&maxes, cur)
			cur = 0
		} else {
			num, err := strconv.Atoi(lines[i])
			if err != nil {
				return "", err
			}
			cur += num
		}
	}
	UpdateMaxThree(&maxes, cur)
	return strconv.Itoa(maxes[0] + maxes[1] + maxes[2]), nil
}

func UpdateMaxThree(arr *[3]int, val int) {
	if val > arr[2] {
		arr[0] = arr[1]
		arr[1] = arr[2]
		arr[2] = val
	} else if val > arr[1] {
		arr[0] = arr[1]
		arr[1] = val
	} else if val > arr[0] {
		arr[0] = val
	}
}
