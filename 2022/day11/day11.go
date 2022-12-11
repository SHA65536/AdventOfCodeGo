package day11

import (
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/11

func MonkeyBusiness(input string, steps, worry uint64) (string, error) {
	var monkeys []*Monkey
	// Monkey div is the product of all the div tests
	// It is the highest a worry level can get
	var monkeyDiv uint64 = 1
	lines := strings.Split(input, "\n\n")
	// Loading monkeys
	for i := range lines {
		monkeys = append(monkeys, ParseMonkey(lines[i]))
		// Adding current div test to monkeyDiv
		monkeyDiv *= monkeys[len(monkeys)-1].TestDiv
	}
	// Acting monkeys
	for ; steps > 0; steps-- {
		for _, m := range monkeys {
			for len(m.Items) > 0 {
				cur := <-m.Items
				cur = m.Operation(cur) / worry
				cur %= monkeyDiv // Limiting the worry
				if cur%m.TestDiv == 0 {
					monkeys[m.ThrowTrue].Items <- cur
				} else {
					monkeys[m.ThrowFalse].Items <- cur
				}
				m.NumChecks++
			}
		}
	}
	// Finding highest
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].NumChecks > monkeys[j].NumChecks
	})
	return strconv.Itoa(monkeys[0].NumChecks * monkeys[1].NumChecks), nil
}

func ParseMonkey(input string) *Monkey {
	var res = &Monkey{Items: make(chan uint64, 256)}
	var lines = strings.Split(input, "\n")
	// Monkey number
	res.Num = int(lines[0][7] - '0')
	// Initial list
	for i := 18; i < len(lines[1]); i++ {
		var cur uint64
		for cur = 0; i < len(lines[1]) && lines[1][i] != ','; i++ {
			cur = cur*10 + uint64(lines[1][i]-'0')
		}
		i++
		res.Items <- cur
	}
	// Operation
	op := lines[2][23]
	if lines[2][25] == 'o' { // Referencing old
		res.Operation = func(a uint64) uint64 {
			if op == '*' {
				return a * a
			}
			return a + a
		}
	} else { // Referencing a const
		val, _ := strconv.ParseUint(lines[2][25:], 10, 64)
		res.Operation = func(a uint64) uint64 {
			if op == '*' {
				return a * val
			}
			return a + val
		}
	}
	// Test
	res.TestDiv, _ = strconv.ParseUint(lines[3][21:], 10, 64)
	res.ThrowTrue, res.ThrowFalse = int(lines[4][29]-'0'), int(lines[5][30]-'0')
	return res
}

type Monkey struct {
	Num        int
	Items      chan uint64
	Operation  func(uint64) uint64
	TestDiv    uint64
	ThrowTrue  int
	ThrowFalse int
	NumChecks  int
}
