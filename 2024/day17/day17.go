package day17

import (
	"adventofcode/helper"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Computer struct {
	RA, RB, RC int
	PC         int
	Program    []int
}

func Star1(input *helper.InputReader) (string, error) {
	var res string

	var computer = parseInput(input)

	var output []string

	for computer.PC < len(computer.Program)-1 {
		if out, ok := computer.DoNext(); ok {
			output = append(output, string(byte(out+'0')))
		}
	}

	res = strings.Join(output, ",")

	return res, nil
}

func (c *Computer) DoNext() (out int, isOut bool) {
	var OpCode, Oper = c.Program[c.PC], c.Program[c.PC+1]
	switch OpCode {
	case 0:
		c.Adv(Oper)
	case 1:
		c.Bxl(Oper)
	case 2:
		c.Bst(Oper)
	case 3:
		c.Jnz(Oper)
	case 4:
		c.Bxc(Oper)
	case 5:
		out = c.Out(Oper)
		isOut = true
	case 6:
		c.Bdv(Oper)
	case 7:
		c.Cdv(Oper)
	}
	return
}

func (c *Computer) Combo(op int) int {
	if op < 4 {
		return op
	}
	switch op {
	case 4:
		return c.RA
	case 5:
		return c.RB
	case 6:
		return c.RC
	}
	return -1
}

func (c *Computer) Adv(op int) {
	var res = c.RA / powInt(2, c.Combo(op))
	c.RA = res
	c.PC += 2
}

func (c *Computer) Bxl(op int) {
	var res = c.RB ^ op
	c.RB = res
	c.PC += 2
}

func (c *Computer) Bst(op int) {
	var res = c.Combo(op) % 8
	c.RB = res
	c.PC += 2
}

func (c *Computer) Jnz(op int) {
	if c.RA == 0 {
		c.PC += 2
	} else {
		c.PC = op
	}
}

func (c *Computer) Bxc(op int) {
	var res = c.RB ^ c.RC
	c.RB = res
	c.PC += 2
}

func (c *Computer) Out(op int) int {
	c.PC += 2
	return c.Combo(op) % 8
}

func (c *Computer) Bdv(op int) {
	var res = c.RA / powInt(2, c.Combo(op))
	c.RB = res
	c.PC += 2
}

func (c *Computer) Cdv(op int) {
	var res = c.RA / powInt(2, c.Combo(op))
	c.RC = res
	c.PC += 2
}

func parseInput(input *helper.InputReader) *Computer {
	var res Computer
	line, _ := input.ReadLine()
	fmt.Sscanf(line, "Register A: %d", &res.RA)
	line, _ = input.ReadLine()
	fmt.Sscanf(line, "Register B: %d", &res.RB)
	line, _ = input.ReadLine()
	fmt.Sscanf(line, "Register C: %d", &res.RC)

	input.ReadLine()
	line, _ = input.ReadLine()
	nums := strings.Split(line[9:], ",")
	for i := 0; i < len(nums); i++ {
		res.Program = append(res.Program, helper.MustConvNum(nums[i]))
	}

	return &res
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var computer = parseInput(input)

	var find func(cur int, prog []int) int
	find = func(cur int, prog []int) int {
		if len(prog) == 0 {
			return cur
		}
		for b := range 8 {
			a := (cur << 3) | b
			b = b ^ 2
			c := a >> b
			b = b ^ 7
			b = b ^ c
			if b%8 == prog[len(prog)-1] {
				if sub := find(a, prog[:len(prog)-1]); sub != -1 {
					return sub
				}
				continue
			}
		}
		return -1
	}

	res = find(0, computer.Program)

	return strconv.Itoa(res), nil
}
