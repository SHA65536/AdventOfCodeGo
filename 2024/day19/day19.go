package day19

import (
	"adventofcode/helper"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var available = parseAvailable(input)

	var cache = map[string]bool{}

	for line := range input.IterateLines {
		if isPossible(line, cache, available) {
			res++
		}
	}

	return strconv.Itoa(res), nil
}

func isPossible(pattern string, cache map[string]bool, available []string) bool {
	if v, ok := cache[pattern]; ok {
		return v
	}

	if len(pattern) == 0 {
		return true
	}

	for _, towel := range available {
		if strings.HasPrefix(pattern, towel) && isPossible(pattern[len(towel):], cache, available) {
			cache[pattern] = true
			return true
		}
	}

	cache[pattern] = false
	return false
}

func parseAvailable(input *helper.InputReader) []string {
	line, _ := input.ReadLine()
	input.ReadLine()
	return strings.Split(line, ", ")
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var available = parseAvailable(input)

	var cache = map[string]int{}

	for line := range input.IterateLines {
		res += isPossibleCount(line, cache, available)
	}

	return strconv.Itoa(res), nil
}

func isPossibleCount(pattern string, cache map[string]int, available []string) int {
	if v, ok := cache[pattern]; ok {
		return v
	}

	if len(pattern) == 0 {
		return 1
	}

	var ways int
	for _, towel := range available {
		if strings.HasPrefix(pattern, towel) {
			ways += isPossibleCount(pattern[len(towel):], cache, available)
		}
	}

	cache[pattern] = ways
	return ways
}

func Star1Map(input *helper.InputReader) (string, error) {
	var res int

	var available, maxLen = parseAvailableMap(input)

	var cache = map[string]bool{}

	for line := range input.IterateLines {
		if isPossibleMap(line, cache, available, maxLen) {
			res++
		}
	}

	return strconv.Itoa(res), nil
}

func isPossibleMap(pattern string, cache map[string]bool, available map[string]struct{}, maxLen int) bool {
	if v, ok := cache[pattern]; ok {
		return v
	}

	if len(pattern) == 0 {
		return true
	}

	for i := 1; i <= maxLen && i <= len(pattern); i++ {
		if _, ok := available[pattern[:i]]; ok && isPossibleMap(pattern[i:], cache, available, maxLen) {
			cache[pattern] = true
			return true
		}
	}

	cache[pattern] = false
	return false
}

func Star2Map(input *helper.InputReader) (string, error) {
	var res int

	var available, maxLen = parseAvailableMap(input)

	var cache = map[string]int{}

	for line := range input.IterateLines {
		res += isPossibleCountMap(line, cache, available, maxLen)
	}

	return strconv.Itoa(res), nil
}

func isPossibleCountMap(pattern string, cache map[string]int, available map[string]struct{}, maxLen int) int {
	if v, ok := cache[pattern]; ok {
		return v
	}

	if len(pattern) == 0 {
		return 1
	}

	var ways int
	for i := 1; i <= maxLen && i <= len(pattern); i++ {
		if _, ok := available[pattern[:i]]; ok {
			ways += isPossibleCountMap(pattern[i:], cache, available, maxLen)
		}
	}

	cache[pattern] = ways
	return ways
}

func parseAvailableMap(input *helper.InputReader) (map[string]struct{}, int) {
	var res = map[string]struct{}{}
	line, _ := input.ReadLine()
	input.ReadLine()

	var maxLen int
	for _, word := range strings.Split(line, ", ") {
		res[word] = struct{}{}
		maxLen = max(maxLen, len(word))
	}

	return res, maxLen
}
