package day01

import (
	"adventofcode/helper"
	"container/heap"
	"sort"
	"strconv"
	"strings"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var left, right = make([]int, 0, 1000), make([]int, 0, 1000)

	for line := range input.IterateLines {
		nums := strings.Split(line, "   ")
		left = append(left, helper.MustConvNum(nums[0]))
		right = append(right, helper.MustConvNum(nums[1]))
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := range left {
		res += absDist(left[i], right[i])
	}

	return strconv.Itoa(res), nil
}

func Star1Heap(input *helper.InputReader) (string, error) {
	var res int

	var left, right = helper.MakeHeap[int](helper.MinHeap), helper.MakeHeap[int](helper.MinHeap)
	heap.Init(left)
	heap.Init(right)

	for line := range input.IterateLines {
		nums := strings.Split(line, "   ")
		heap.Push(left, helper.MustConvNum(nums[0]))
		heap.Push(right, helper.MustConvNum(nums[1]))
	}

	for left.Len() > 0 {
		res += absDist(heap.Pop(left).(int), heap.Pop(right).(int))
	}

	return strconv.Itoa(res), nil
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var left = make([]int, 0, 1000)
	var right = make(map[int]int, 1000)

	for line := range input.IterateLines {
		nums := strings.Split(line, "   ")
		left = append(left, helper.MustConvNum(nums[0]))

		right[helper.MustConvNum(nums[1])]++
	}

	for _, num := range left {
		res += num * (right[num])
	}

	return strconv.Itoa(res), nil
}

func absDist(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
