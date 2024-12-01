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

	var left, right []int

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

	var left, right = &MinHeap{}, &MinHeap{}
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

	var left []int
	var right = map[int]int{}

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

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
