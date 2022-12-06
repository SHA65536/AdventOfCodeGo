package day15

import (
	"container/heap"
	"strconv"
)

// https://adventofcode.com/2021/day/15

var dirs = [][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func LowestPath(risks string, tiles int) (string, error) {
	var board [][]int

	// Loading original board
	for i := 0; i < len(risks); i++ {
		var cur []int
		for ; i < len(risks) && risks[i] != '\n'; i++ {
			cur = append(cur, int(risks[i]-'0'))
		}
		board = append(board, cur)
	}

	// Places seen
	var seen = map[[2]int]bool{{0, 0}: true}
	// Min heap for calculations
	var minHeap = MakeHeap(MinHeap)

	// Starting Node
	heap.Push(minHeap, Node{0, 0, 0})

	for minHeap.Len() > 0 {
		prev := heap.Pop(minHeap).(Node)
		// If we reached the end
		if prev.X == tiles*len(board)-1 && prev.Y == tiles*len(board[0])-1 {
			return strconv.Itoa(prev.Value), nil
		}
		// Checking each direcion
		for _, dir := range dirs {
			// New coords
			nx, ny := prev.X+dir[0], prev.Y+dir[1]
			// Checking if outside the board
			if nx < 0 || ny < 0 || nx >= tiles*len(board) || ny >= tiles*len(board[0]) {
				continue
			}
			// Calculating added risk level
			a, am := nx/len(board), nx%len(board)
			b, bm := ny/len(board[0]), ny%len(board[0])
			n := ((board[am][bm]+a+b)-1)%9 + 1
			// If we didn't calculate this already, add to nodes to consider
			if !seen[[2]int{nx, ny}] {
				seen[[2]int{nx, ny}] = true
				// New risk level
				heap.Push(minHeap, Node{prev.Value + n, nx, ny})
			}
		}
	}

	return "-1", nil
}

type Node struct {
	Value int
	X, Y  int
}

type Heap struct {
	Values   []Node
	LessFunc func(int, int) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i].Value, h.Values[j].Value) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() Node         { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(Node)) }

func MakeHeap(less func(int, int) bool) *Heap {
	return &Heap{LessFunc: less}
}

func MaxHeap(i, j int) bool { return i > j }
func MinHeap(i, j int) bool { return i < j }
