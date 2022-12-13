package day13

import (
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/13

func InRightOrder(input string) (string, error) {
	var res int
	pairs := strings.Split(input, "\n\n")

	// Looping over pairs
	for i := range pairs {
		packets := strings.Split(pairs[i], "\n")
		// Creating items
		left, right := CreateItem(packets[0]), CreateItem(packets[1])
		// Checking if in the right order
		if left.IsEqual(right) == 1 {
			res += i + 1
		}
	}

	return strconv.Itoa(res), nil
}

func SortPackets(input string) (string, error) {
	var res int = 1
	var divider1, divider2 = CreateItem("[[2]]"), CreateItem("[[6]]")
	var items = []*Item{divider1, divider2}

	// Looping over lines
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		items = append(items, CreateItem(line))
	}

	// Sorting items
	sort.Slice(items, func(i, j int) bool {
		return items[i].IsEqual(items[j]) == 1
	})

	// Summing items
	for i := range items {
		if items[i].IsEqual(divider1) == 0 || items[i].IsEqual(divider2) == 0 {
			res *= (i + 1)
		}
	}
	return strconv.Itoa(res), nil
}

type Item struct {
	Val  int
	List []*Item
}

// IsEqual returns -1 if left is bigger, 0 if equal, 1 if right is bigger
func (l *Item) IsEqual(r *Item) int {
	// Comparing value vs value
	if l.List == nil && r.List == nil {
		if l.Val == r.Val {
			return 0
		} else if l.Val > r.Val {
			return -1
		} else {
			return 1
		}
	}
	// If type mismatch, convert to list
	if l.List == nil && r.List != nil {
		l = &Item{Val: 0, List: []*Item{l}}
	}
	if r.List == nil && l.List != nil {
		r = &Item{Val: 0, List: []*Item{r}}
	}
	// Comparing list vs list
	for i := 0; i <= len(l.List); i++ {
		// If r ran out of items
		if i < len(l.List) && i >= len(r.List) {
			return -1
		}
		// If l ran out of items
		if i < len(r.List) && i >= len(l.List) {
			return 1
		}
		// If both are out of items
		if i >= len(r.List) && i >= len(l.List) {
			return 0
		}
		// Comparing the current value
		res := l.List[i].IsEqual(r.List[i])
		if res != 0 {
			return res
		}
	}
	return 0
}

// Loads an item
func CreateItem(input string) *Item {
	var res = &Item{Val: -1}
	var stack []*Item
	for i := 1; i < len(input)-1; i++ {
		if isDigit(input[i]) { // If current is digit
			var cur int
			for ; i < len(input)-1 && isDigit(input[i]); i++ {
				cur = cur*10 + int(input[i]-'0')
			}
			i--
			res.List = append(res.List, &Item{Val: cur})
		} else if input[i] == '[' {
			stack = append(stack, res)
			res = &Item{Val: -1}
		} else if input[i] == ']' {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prev.List = append(prev.List, res)
			res = prev
		}
	}
	return res
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
