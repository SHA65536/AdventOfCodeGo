package day11

import (
	"adventofcode/helper"
	"container/list"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var stones = makeStones(input)

	for i := 25; i > 0; i-- {
		for e := stones.Front(); e != nil; e = e.Next() {
			if e.Value.(int) == 0 {
				e.Value = 1
				continue
			}
			if left, right, ok := splitInteger(e.Value.(int)); ok {
				e.Value = right
				stones.InsertBefore(left, e)
				continue
			}
			e.Value = e.Value.(int) * 2024
		}
	}

	res = stones.Len()

	return strconv.Itoa(res), nil
}

func makeStones(input *helper.InputReader) *list.List {
	var res = list.New()

	var _line, _ = input.ReadAll()
	for _, num := range strings.Fields(_line) {
		res.PushBack(helper.MustConvNum(num))
	}

	return res
}

func splitInteger(n int) (int, int, bool) {
	nStr := strconv.Itoa(n)

	if len(nStr)%2 != 0 {
		return 0, 0, false
	}

	mid := (len(nStr)) / 2

	return helper.MustConvNum(nStr[:mid]), helper.MustConvNum(nStr[mid:]), true
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var cache = make(map[[2]int]int, 0)

	var calcStone func(st int, it int) int
	calcStone = func(st, it int) int {
		if v, ok := cache[[2]int{st, it}]; ok {
			return v
		}
		if it == 0 {
			return 1
		}

		if st == 0 {
			ret := calcStone(1, it-1)
			cache[[2]int{st, it}] = ret
			return ret
		}

		if l, r, ok := splitInteger(st); ok {
			ret := calcStone(l, it-1) + calcStone(r, it-1)
			cache[[2]int{st, it}] = ret
			return ret
		}

		ret := calcStone(st*2024, it-1)
		cache[[2]int{st, it}] = ret
		return ret
	}

	var _line, _ = input.ReadAll()
	for _, num := range strings.Fields(_line) {
		res += (calcStone(helper.MustConvNum(num), 75))
	}

	return strconv.Itoa(res), nil
}
