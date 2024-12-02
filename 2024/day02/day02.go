package day02

import (
	"adventofcode/helper"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		words := strings.Split(line, " ")
		var nums = make([]int, 0, len(words))
		for _, word := range words {
			nums = append(nums, helper.MustConvNum(word))
		}

		if TryLine(nums) {
			res++
		}
	}

	return strconv.Itoa(res), nil
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		words := strings.Split(line, " ")
		var nums = make([]int, 0, len(words))
		for _, word := range words {
			nums = append(nums, helper.MustConvNum(word))
		}

		for i := range nums {
			var without = make([]int, len(nums))
			copy(without, nums)
			if TryLine(append(without[:i], without[i+1:]...)) {
				res++
				break
			}
		}
	}

	return strconv.Itoa(res), nil
}

func TryLine(nums []int) bool {
	var asc bool = nums[0] < nums[1]
	var prev = nums[0]
	for _, num := range nums[1:] {
		if asc && (prev >= num || absDist(num, prev) > 3) {
			return false
		}
		if !asc && (prev <= num || absDist(num, prev) > 3) {
			return false
		}
		prev = num
	}
	return true
}

func absDist(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
