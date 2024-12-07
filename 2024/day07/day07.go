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

func concat(a, b int) int {
	var big = 1
	for b >= big {
		big *= 10
	}
	return (a * big) + b
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		var nums []int
		words := strings.Fields(line)
		var end = helper.MustConvNum(words[0][:len(words[0])-1])
		for _, word := range words[1:] {
			nums = append(nums, helper.MustConvNum(word))
		}

		if isPossible(end, nums, []byte{'|', '*', '+'}) {
			res += end
		}
	}

	return strconv.Itoa(res), nil
}

func Star1Iter(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		var nums []int
		words := strings.Fields(line)
		var end = helper.MustConvNum(words[0][:len(words[0])-1])
		for _, word := range words[1:] {
			nums = append(nums, helper.MustConvNum(word))
		}

		if isPossibleIterator(end, nums, []byte{'*', '+'}) {
			res += end
		}
	}

	return strconv.Itoa(res), nil
}

func Star2Iter(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		var nums []int
		words := strings.Fields(line)
		var end = helper.MustConvNum(words[0][:len(words[0])-1])
		for _, word := range words[1:] {
			nums = append(nums, helper.MustConvNum(word))
		}

		if isPossibleIterator(end, nums, []byte{'|', '*', '+'}) {
			res += end
		}
	}

	return strconv.Itoa(res), nil
}

func isPossibleIterator(res int, nums []int, ops []byte) bool {
	for comb := range helper.IteratePermutations(ops, len(nums)-1) {
		if res == Calc(nums, comb, res) {
			return true
		}
	}
	return false
}

func Star1Rec(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		var nums []int
		words := strings.Fields(line)
		var end = helper.MustConvNum(words[0][:len(words[0])-1])
		for _, word := range words[1:] {
			nums = append(nums, helper.MustConvNum(word))
		}

		if isPossibleRecursive(end, nums[0], nums[1:], []byte{'*', '+'}) {
			res += end
		}
	}

	return strconv.Itoa(res), nil
}

func Star2Rec(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		var nums []int
		words := strings.Fields(line)
		var end = helper.MustConvNum(words[0][:len(words[0])-1])
		for _, word := range words[1:] {
			nums = append(nums, helper.MustConvNum(word))
		}

		if isPossibleRecursive(end, nums[0], nums[1:], []byte{'|', '*', '+'}) {
			res += end
		}
	}

	return strconv.Itoa(res), nil
}

func isPossibleRecursive(end, acc int, nums []int, ops []byte) bool {
	if len(nums) == 0 {
		return acc == end
	}

	for _, op := range ops {
		var temp = Op(acc, nums[0], op)
		if isPossibleRecursive(end, temp, nums[1:], ops) {
			return true
		}
	}
	return false
}
