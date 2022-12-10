package day10

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/10

func CycleSum(input string) (string, error) {
	var res int
	var x, nxt, cycle int = 1, 20, 1

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		var val, repeat int
		op := line[:4]
		if op == "noop" { // noop takes 1 turn
			val = 0
			repeat = 1
		} else { // addx takes 2 turns
			val, _ = strconv.Atoi(line[5:])
			repeat = 2
		}
		for repeat > 0 {
			// Summing the res
			if cycle == nxt {
				res += x * cycle
				nxt += 40
			}
			cycle++
			repeat--
		}
		x += val
	}

	return strconv.Itoa(res), nil
}

func DrawScreen(input string) (string, error) {
	var res strings.Builder

	lines := strings.Split(input, "\n")
	var curLine = make([]byte, 40)
	var x, nxt, cycle int = 1, 39, 0

	for _, line := range lines {
		var val, repeat int
		op := line[:4]
		if op == "noop" { // noop takes 1 turn
			val = 0
			repeat = 1
		} else { // addx takes 2 turns
			val, _ = strconv.Atoi(line[5:])
			repeat = 2
		}
		for repeat > 0 {
			// Drawing
			curLine[cycle] = '.'
			if abs(x-cycle) < 2 {
				curLine[cycle] = '#'
			}
			// Add after draw
			if repeat == 1 {
				x += val
			}
			// Going down a line
			if cycle == nxt {
				cycle = -1
				res.WriteString(string(curLine) + "\n")
				curLine = make([]byte, 40)
			}
			cycle++
			repeat--
		}
	}

	return res.String(), nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
