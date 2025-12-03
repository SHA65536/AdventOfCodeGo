package day03

import (
	"adventofcode/helper"
	"strconv"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		first, idx := findMax(line[:len(line)-1])
		second, _ := findMax(line[idx+1:])
		res += first*10 + second
	}

	return strconv.Itoa(res), nil
}

func findMax(input string) (res, idx int) {
	for i, val := range input {
		if int(val-'0') > res {
			res = int(val - '0')
			idx = i
		}
	}
	return res, idx
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	for line := range input.IterateLines {
		var curSum, prevIdx int
		for i := 11; i >= 0; i-- {
			val, idx := findMax(line[prevIdx : len(line)-i])
			curSum = curSum*10 + val
			prevIdx += idx + 1
		}
		res += curSum
	}

	return strconv.Itoa(res), nil
}
