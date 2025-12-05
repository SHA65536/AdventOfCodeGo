package day05

import (
	"adventofcode/helper"
	"sort"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int
	var ranges [][2]int

	for line := range input.IterateLines {
		switch {
		case line == "":
			continue
		case strings.Contains(line, "-"):
			parts := strings.Split(line, "-")
			ranges = append(ranges, [2]int{helper.MustConvNum(parts[0]), helper.MustConvNum(parts[1])})
		default:
			if checkFresh(ranges, helper.MustConvNum(line)) {
				res++
			}
		}
	}

	return strconv.Itoa(res), nil
}

func checkFresh(ranges [][2]int, id int) bool {
	for _, r := range ranges {
		if id >= r[0] && id <= r[1] {
			return true
		}
	}
	return false
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var ranges [][2]int

	for line := range input.IterateLines {
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		ranges = append(ranges, [2]int{helper.MustConvNum(parts[0]), helper.MustConvNum(parts[1])})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	var cur = ranges[0]
	for _, r := range ranges[1:] {
		if r[0] <= cur[1] {
			if r[1] > cur[1] {
				cur[1] = r[1]
			}
			continue
		}
		res += cur[1] - cur[0] + 1
		cur = r
	}

	res += cur[1] - cur[0] + 1

	return strconv.Itoa(res), nil
}
