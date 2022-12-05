package day05

import (
	"strings"
)

// https://adventofcode.com/2022/day/5

func TopCrates(crates string) (string, error) {
	var res strings.Builder
	var crateStacks []CrateStack
	sections := strings.Split(crates, "\n\n")
	crateLines := strings.Split(sections[0], "\n")
	crateStacks = make([]CrateStack, (len(crateLines[0])+1)/4)
	// Loading crates
	for i := len(crateLines) - 2; i >= 0; i-- {
		var cur int
		for j := 1; j < len(crateLines[i]); j += 4 {
			if crateLines[i][j] >= 'A' && crateLines[i][j] <= 'Z' {
				crateStacks[cur].Push(crateLines[i][j])
			}
			cur++
		}
	}
	// Moving crates
	for i := 0; i < len(sections[1]); i++ {
		var size, src, dst int
		i += 5
		// Loading quantity
		for size = 0; sections[1][i] != ' '; i++ {
			size = size*10 + int(sections[1][i]-'0')
		}
		i += 6
		// Loading source
		for src = 0; sections[1][i] != ' '; i++ {
			src = src*10 + int(sections[1][i]-'0')
		}
		i += 4
		// Loading destination
		for dst = 0; i < len(sections[1]) && sections[1][i] != '\n'; i++ {
			dst = dst*10 + int(sections[1][i]-'0')
		}
		// Moving the crates one by one
		for j := 0; j < size; j++ {
			val := crateStacks[src-1].Pop()
			crateStacks[dst-1].Push(val)
		}
	}
	// Writing the top crates
	for i := range crateStacks {
		res.WriteByte(crateStacks[i].Pop())
	}
	return res.String(), nil
}

func TopCratesMulti(crates string) (string, error) {
	var res strings.Builder
	var crateStacks []CrateStack
	sections := strings.Split(crates, "\n\n")
	crateLines := strings.Split(sections[0], "\n")
	crateStacks = make([]CrateStack, (len(crateLines[0])+1)/4)
	// Loading crates
	for i := len(crateLines) - 2; i >= 0; i-- {
		var cur int
		for j := 1; j < len(crateLines[i]); j += 4 {
			if crateLines[i][j] >= 'A' && crateLines[i][j] <= 'Z' {
				crateStacks[cur].Push(crateLines[i][j])
			}
			cur++
		}
	}
	// Moving crates
	for i := 0; i < len(sections[1]); i++ {
		var size, src, dst int
		var temp []byte
		i += 5
		// Loading quantity
		for size = 0; sections[1][i] != ' '; i++ {
			size = size*10 + int(sections[1][i]-'0')
		}
		i += 6
		// Loading source
		for src = 0; sections[1][i] != ' '; i++ {
			src = src*10 + int(sections[1][i]-'0')
		}
		i += 4
		// Loading destination
		for dst = 0; i < len(sections[1]) && sections[1][i] != '\n'; i++ {
			dst = dst*10 + int(sections[1][i]-'0')
		}
		// Moving crates to temporary memory
		for j := 0; j < size; j++ {
			val := crateStacks[src-1].Pop()
			temp = append(temp, val)
		}
		// Moving crates from temporary memory to destination
		for j := len(temp) - 1; j >= 0; j-- {
			crateStacks[dst-1].Push(temp[j])
		}
	}
	// Writing the top crates
	for i := range crateStacks {
		res.WriteByte(crateStacks[i].Pop())
	}
	return res.String(), nil
}

type CrateStack []byte

func (s *CrateStack) Push(val byte) {
	(*s) = append((*s), val)
}

func (s *CrateStack) Pop() byte {
	ret := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return ret
}
