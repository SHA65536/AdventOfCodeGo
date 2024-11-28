package day03

import (
	"adventofcode/helper"
	"strconv"
)

var adjacents = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func Star1(input *helper.InputReader) (string, error) {
	var res, cur int
	var adj bool

	var schematic, err = input.ReadLines()
	if err != nil {
		return "", err
	}

	for i := range schematic {
		for j := range schematic[i] {
			if schematic[i][j] < '0' || schematic[i][j] > '9' {
				if adj {
					res += cur
					adj = false
				}
				cur = 0
				continue
			}
			cur = cur*10 + int(schematic[i][j]-'0')
			for _, ndif := range adjacents {
				ni, nj := i+ndif[0], j+ndif[1]
				if ni < 0 || ni >= len(schematic) || nj < 0 || nj >= len(schematic[0]) {
					continue
				}
				if isSymbol(schematic[ni][nj]) {
					adj = true
					break
				}
			}
		}
		if adj {
			res += cur
			adj = false
		}
		cur = 0
	}

	return strconv.Itoa(res), nil
}

func Star2(input *helper.InputReader) (string, error) {
	var res, cur int
	var adj = map[[2]int]bool{}
	var gears = map[[2]int][]int{}

	var schematic, err = input.ReadLines()
	if err != nil {
		return "", err
	}

	for i := range schematic {
		for j := range schematic[i] {
			// If not a number
			if schematic[i][j] < '0' || schematic[i][j] > '9' {
				// Check if a number ended and had any adhacent
				if cur > 0 && len(adj) > 0 {
					// Add adjacent gears
					for g := range adj {
						gears[g] = append(gears[g], cur)
					}
					clear(adj) // Empty gears
				}
				cur = 0 // Prepare for new number
				continue
			}
			cur = cur*10 + int(schematic[i][j]-'0')
			for _, ndif := range adjacents {
				ni, nj := i+ndif[0], j+ndif[1]
				if ni < 0 || ni >= len(schematic) || nj < 0 || nj >= len(schematic[0]) {
					continue
				}
				if schematic[ni][nj] == '*' {
					adj[[2]int{ni, nj}] = true
					break
				}
			}
		}
		if cur > 0 && len(adj) > 0 {
			for g := range adj {
				gears[g] = append(gears[g], cur)
			}
			clear(adj)
		}
		cur = 0
	}

	for _, gear := range gears {
		if len(gear) == 2 {
			res += gear[0] * gear[1]
		}
	}

	return strconv.Itoa(res), nil
}

func isSymbol(in byte) bool {
	return (in < '0' || in > '9') && in != '.' && in != '\n'
}
