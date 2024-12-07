package day07

import (
	"adventofcode/helper"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		var nums []int
		words := strings.Fields(line)
		var end = helper.MustConvNum(words[0][:len(words[0])-1])
		for _, word := range words[1:] {
			nums = append(nums, helper.MustConvNum(word))
		}

		if isPossible(end, nums, []byte{'*', '+'}) {
			res += end
		}
	}

	return strconv.Itoa(res), nil
}

func isPossible(res int, nums []int, ops []byte) bool {
	var combs = GenerateCombs(ops, len(nums)-1)
	for _, comb := range combs {
		if res == Calc(nums, comb, res) {
			return true
		}
	}
	return false
}

func Calc(nums []int, ops []byte, upper int) int {
	var res int = nums[0]
	for i := range ops {
		res = Op(res, nums[i+1], ops[i])
		if res > upper {
			return -1
		}
	}
	return res
}

func GenerateCombs(chars []byte, size int) [][]byte {
	if size == 0 {
		return [][]byte{{}}
	}

	var results [][]byte
	for _, c := range chars {
		for _, suffix := range GenerateCombs(chars, size-1) {
			results = append(results, append([]byte{c}, suffix...))
		}
	}

	return results
}

func Op(a, b int, op byte) int {
	switch op {
	case '+':
		return a + b
	case '*':
		return a * b
	case '|':
		return concat(a, b)
	}
	return 0
}

func Star2(input *helper.InputReader) (string, error) {
	var res uint64

	for line := range input.IterateLines {
		var nums []int
		words := strings.Fields(line)
		var end = helper.MustConvNum(words[0][:len(words[0])-1])
		for _, word := range words[1:] {
			nums = append(nums, helper.MustConvNum(word))
		}

		if isPossible(end, nums, []byte{'|', '*', '+'}) {
			res += uint64(end)
		}
	}

	return strconv.FormatUint(res, 10), nil
}

func concat(a, b int) int {
	var big = 1
	for b >= big {
		big *= 10
	}
	return (a * big) + b
}
