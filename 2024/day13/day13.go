package day13

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
)

type Machine struct {
	Ax, Ay int
	Bx, By int
	Px, Py int
}

func getMachines(input *helper.InputReader) []Machine {
	var res []Machine

	for line, _ := input.ReadLine(); line != ""; line, _ = input.ReadLine() {
		var cur Machine
		fmt.Sscanf(line, "Button A: X+%d, Y+%d", &cur.Ax, &cur.Ay)
		line, _ = input.ReadLine()
		fmt.Sscanf(line, "Button B: X+%d, Y+%d", &cur.Bx, &cur.By)
		line, _ = input.ReadLine()
		fmt.Sscanf(line, "Prize: X=%d, Y=%d", &cur.Px, &cur.Py)
		line, _ = input.ReadLine()

		res = append(res, cur)
	}

	return res
}

func Star1(input *helper.InputReader) (string, error) {
	var res int

	for _, m := range getMachines(input) {
		a, b := numTokens(m)
		res += a*3 + b
	}

	return strconv.Itoa(res), nil
}

func numTokens(m Machine) (a, b int) {
	var den = m.Ax*m.By - m.Ay*m.Bx
	var num_a = m.By*m.Px - m.Bx*m.Py
	var num_b = m.Ay*m.Px - m.Ax*m.Py

	if num_a%den == 0 && num_b%den == 0 {
		return num_a / den, num_b / (-den)
	}
	return 0, 0
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	for _, m := range getMachines(input) {
		m.Px += 10000000000000
		m.Py += 10000000000000
		a, b := numTokens(m)
		res += a*3 + b
	}

	return strconv.Itoa(res), nil
}
