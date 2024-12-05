package day05

import (
	"adventofcode/helper"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var after = map[int]map[int]bool{}

	for line := range input.IterateLines {
		if strings.Contains(line, "|") {
			var bef, aft int
			fmt.Sscanf(line, "%d|%d", &bef, &aft)
			if after[bef] == nil {
				after[bef] = map[int]bool{aft: true}
			} else {
				after[bef][aft] = true
			}
		} else {
			var nums, mid = parsePages(line)
			if checkOrder(nums, after) {
				res += mid
			}
		}
	}

	return strconv.Itoa(res), nil
}

func parsePages(line string) ([]int, int) {
	var pages []int
	for _, num := range strings.Split(line, ",") {
		pages = append(pages, helper.MustConvNum(num))
	}
	return pages, pages[len(pages)/2]
}

func checkOrder(nums []int, after map[int]map[int]bool) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j != i; j++ {
			if after[nums[i]][nums[j]] {
				return false
			}
		}
	}
	return true
}

func Star1Sort(input *helper.InputReader) (string, error) {
	var res int

	var after = map[int]map[int]bool{}

	for line := range input.IterateLines {
		if strings.Contains(line, "|") {
			var bef, aft int
			fmt.Sscanf(line, "%d|%d", &bef, &aft)
			if after[bef] == nil {
				after[bef] = map[int]bool{aft: true}
			} else {
				after[bef][aft] = true
			}
		} else {
			var nums, mid = parsePages(line)
			var sorted = make([]int, len(nums))
			copy(sorted, nums)
			sort.Slice(nums, func(i, j int) bool {
				return after[nums[i]][nums[j]]
			})
			if slices.Equal(nums, sorted) {
				res += mid
			}
		}
	}

	return strconv.Itoa(res), nil
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var after = map[int]map[int]bool{}

	for line := range input.IterateLines {
		if strings.Contains(line, "|") {
			var bef, aft int
			fmt.Sscanf(line, "%d|%d", &bef, &aft)
			if after[bef] == nil {
				after[bef] = map[int]bool{aft: true}
			} else {
				after[bef][aft] = true
			}
		} else {
			var nums, _ = parsePages(line)
			var sorted = make([]int, len(nums))
			copy(sorted, nums)
			sort.Slice(sorted, func(i, j int) bool {
				return after[sorted[i]][sorted[j]]
			})
			if !slices.Equal(nums, sorted) {
				res += sorted[len(sorted)/2]
			}
		}
	}

	return strconv.Itoa(res), nil
}
