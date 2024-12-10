package helper

var around = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {-1, 1},
}

var adjacent = [4][2]int{
	{-1, 0}, {0, -1}, {0, 1}, {1, 0},
}

func InsideC[T any](board [][]T, i, j int) bool {
	return i >= 0 && i < len(board) && j >= 0 && j < len(board[0])
}

func InsideP[T any](board [][]T, pos [2]int) bool {
	return pos[0] >= 0 && pos[0] < len(board) && pos[1] >= 0 && pos[1] < len(board[0])
}

func IterateAroundC(i, j int) func(yield func(ni, nj int) bool) {
	return func(yield func(ni, nj int) bool) {
		for _, dir := range around {
			if !yield(i+dir[0], j+dir[1]) {
				return
			}
		}
	}
}

func IterateAroundP(i, j int) func(yield func(pos [2]int) bool) {
	return func(yield func(pos [2]int) bool) {
		for _, dir := range around {
			if !yield([2]int{i + dir[0], j + dir[1]}) {
				return
			}
		}
	}
}

func IterateAdjacentC(i, j int) func(yield func(ni, nj int) bool) {
	return func(yield func(ni, nj int) bool) {
		for _, dir := range adjacent {
			if !yield(i+dir[0], j+dir[1]) {
				return
			}
		}
	}
}

func IterateAdjacentP(i, j int) func(yield func(pos [2]int) bool) {
	return func(yield func(pos [2]int) bool) {
		for _, dir := range adjacent {
			if !yield([2]int{i + dir[0], j + dir[1]}) {
				return
			}
		}
	}
}
