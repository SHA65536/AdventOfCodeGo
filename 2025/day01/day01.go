package day01

import (
	"adventofcode/helper"
	"strconv"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int
	var cur int = 50

	for line := range input.IterateLines {
		val := helper.MustConvNum(line[1:]) % 100

		if line[0] == 'L' {
			val = -val
		}

		cur = helper.Mod(cur+val, 100)

		if cur == 0 {
			res++
		}
	}

	return strconv.Itoa(res), nil
}

func Star2(input *helper.InputReader) (string, error) {
	var res int
	var cur int = 50
	var prev int = 1

	for line := range input.IterateLines {
		val := helper.MustConvNum(line[1:])

		res += val / 100 // Full rotations
		val %= 100
		if line[0] == 'L' {
			val = -val
		}

		cur = helper.Mod(cur+val, 100)

		if cur == 0 {
			// Ends on 0
			res++
		} else if prev != 0 && (cur >= prev && line[0] == 'L' || cur <= prev && line[0] == 'R') {
			// Goes through 0
			res++
		}
		prev = cur
	}

	return strconv.Itoa(res), nil
}
