package day25

import (
	"adventofcode/helper"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	keys, locks := parseSchematic(input)

	for _, key := range keys {
		for _, lock := range locks {
			if tryKey(key, lock) {
				res++
			}
		}
	}

	return strconv.Itoa(res), nil
}

func tryKey(key, lock [5]int) bool {
	for pin := range key {
		if key[pin]+lock[pin] > 5 {
			return false
		}
	}
	return true
}

func parseSchematic(input *helper.InputReader) (locks, keys [][5]int) {
	var raw, _ = input.ReadAll()
	var schematics = strings.Split(raw, "\n\n")

	for _, schematic := range schematics {
		if strings.HasPrefix(schematic, ".....") {
			keys = append(keys, parseKey(schematic))
		}
		if strings.HasPrefix(schematic, "#####") {
			locks = append(locks, parseLock(schematic))
		}
	}
	return locks, keys
}

func parseLock(lock string) [5]int {
	var res [5]int
	lines := strings.Split(lock, "\n")

	for pin := range res {
		for i := range lines {
			if lines[i][pin] == '.' {
				res[pin] = i - 1
				break
			}
		}
	}

	return res
}

func parseKey(key string) [5]int {
	var res [5]int
	lines := strings.Split(key, "\n")

	for pin := range res {
		for i := len(lines) - 1; i >= 0; i-- {
			if lines[i][pin] == '.' {
				res[pin] = len(lines[i]) - i
				break
			}
		}
	}

	return res
}
