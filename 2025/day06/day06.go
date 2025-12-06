package day06

import (
	"adventofcode/helper"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int
	var operators []string
	var numbers [][]int
	for line := range input.IterateLines {
		var numLine []int
		if line[0] == '+' || line[0] == '*' {
			operators = strings.Fields(line)
			break
		}
		for _, num := range strings.Fields(line) {
			numLine = append(numLine, helper.MustConvNum(num))
		}
		numbers = append(numbers, numLine)
	}

	for idx, op := range operators {
		var cur int
		if op == "*" {
			cur = 1
		}
		for i := 0; i < len(numbers); i++ {
			if op == "+" {
				cur += numbers[i][idx]
			} else {
				cur *= numbers[i][idx]
			}
		}
		res += cur
	}

	return strconv.Itoa(res), nil
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	grid, _ := input.ReadByteLines()

	var curNums []int
	for col := len(grid[0]) - 1; col >= 0; col-- {
		var curNum int
		for row := 0; row < len(grid); row++ {
			switch grid[row][col] {
			case '*', '+':
				curNums = append(curNums, curNum)
				res += DoMath(curNums, grid[row][col])
				curNums = curNums[:0]
				curNum = 0
				continue
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				curNum = curNum*10 + int(grid[row][col]-'0')
			}
		}
		if curNum != 0 {
			curNums = append(curNums, curNum)
		}
	}

	return strconv.Itoa(res), nil
}

func DoMath(nums []int, op byte) int {
	var res int
	if op == '*' {
		res = 1
	}
	for _, num := range nums {
		if op == '*' {
			res *= num
		} else {
			res += num
		}
	}
	return res
}
