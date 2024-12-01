package day01

import (
	"adventofcode/helper"
	"sort"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var left, right []int

	for line := range input.IterateLines {
		nums := strings.Split(line, "   ")
		left = append(left, helper.MustConvNum(nums[0]))
		right = append(right, helper.MustConvNum(nums[1]))
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := range left {
		res += absDist(left[i], right[i])
	}

	return strconv.Itoa(res), nil
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var left []int
	var right = map[int]int{}

	for line := range input.IterateLines {
		nums := strings.Split(line, "   ")
		left = append(left, helper.MustConvNum(nums[0]))

		right[helper.MustConvNum(nums[1])]++
	}

	for _, num := range left {
		res += num * (right[num])
	}

	return strconv.Itoa(res), nil
}

func absDist(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
