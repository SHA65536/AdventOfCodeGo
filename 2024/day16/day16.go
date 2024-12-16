package day16

import (
	"adventofcode/helper"
	"strconv"
)

var dirs = [4][2]int{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1},
}

func Star1(input *helper.InputReader) (string, error) {
	var res int = -1

	var maze, dI, dJ = parseMaze(input)
	var cache = make([][][4]int, len(maze))
	for i := range maze {
		cache[i] = make([][4]int, len(maze[i]))
		for j := range cache[i] {
			cache[i][j] = [4]int{-1, -1, -1, -1}
		}
	}

	var dfs func(i, j, d, sc int)
	dfs = func(i, j, d, sc int) {
		// Outside or obstacle
		if !helper.InsideC(maze, i, j) || maze[i][j] == '#' {
			return
		}

		// Better path already with this dir
		if t := cache[i][j]; t[d] != -1 && t[d] < sc {
			return
		}

		// Mark cur score as best
		cache[i][j][d] = sc

		// Stop at end
		if maze[i][j] == 'E' {
			if res == -1 {
				res = sc
			}
			res = min(res, sc)
			return
		}

		// Try move forward
		dfs(i+dirs[d][0], j+dirs[d][1], d, sc+1)
		// Try right
		dfs(i, j, mod4(d+1), sc+1000)
		// Try left
		dfs(i, j, mod4(d-1), sc+1000)
	}

	dfs(dI, dJ, 1, 0)

	return strconv.Itoa(res), nil
}

func parseMaze(input *helper.InputReader) ([][]byte, int, int) {
	var maze [][]byte
	var x, y int
	var i int
	for line := range input.IterateLines {
		maze = append(maze, []byte(line))
		for j := range line {
			if line[j] == 'S' {
				x = i
				y = j
			}
		}
		i++
	}
	return maze, x, y
}

func mod4(i int) int {
	i = i % 4
	if i < 0 {
		i += 4
	}
	return i
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var bestScore int = -1

	var maze, dI, dJ = parseMaze(input)
	var cache = make([][][4]int, len(maze))
	for i := range maze {
		cache[i] = make([][4]int, len(maze[i]))
		for j := range cache[i] {
			cache[i][j] = [4]int{-1, -1, -1, -1}
		}
	}

	// Find best score
	var dfs func(i, j, d, sc int)
	dfs = func(i, j, d, sc int) {
		if !helper.InsideC(maze, i, j) || maze[i][j] == '#' {
			return
		}
		if t := cache[i][j]; t[d] != -1 && t[d] < sc {
			return
		}
		cache[i][j][d] = sc
		if maze[i][j] == 'E' {
			if bestScore == -1 {
				bestScore = sc
			}
			bestScore = min(bestScore, sc)
			return
		}
		dfs(i+dirs[d][0], j+dirs[d][1], d, sc+1)
		dfs(i, j, mod4(d+1), sc+1000)
		dfs(i, j, mod4(d-1), sc+1000)
	}
	dfs(dI, dJ, 1, 0)

	cache = make([][][4]int, len(maze))
	var lastSpot = make([][][2]int, len(maze))
	for i := range maze {
		cache[i] = make([][4]int, len(maze[i]))
		lastSpot[i] = make([][2]int, len(maze[i]))
		for j := range cache[i] {
			lastSpot[i][j] = [2]int{-1, -1}
			cache[i][j] = [4]int{-1, -1, -1, -1}
		}
	}

	// Find all spot with best score
	var sits = map[[2]int]struct{}{}
	var dfsEnd func(i, j, d, sc int) bool
	dfsEnd = func(i, j, d, sc int) bool {
		if sc > bestScore {
			return false
		}
		if !helper.InsideC(maze, i, j) || maze[i][j] == '#' {
			return false
		}
		if t := cache[i][j]; t[d] != -1 && t[d] < sc {
			return false
		}
		cache[i][j][d] = sc
		if maze[i][j] == 'E' {
			return sc == bestScore
		}
		var t [2]int
		if helper.InsideC(maze, i+dirs[d][0], j+dirs[d][1]) {
			t = lastSpot[i+dirs[d][0]][j+dirs[d][1]]
			lastSpot[i+dirs[d][0]][j+dirs[d][1]] = [2]int{i, j}
		}
		if dfsEnd(i+dirs[d][0], j+dirs[d][1], d, sc+1) {
			backTrack(i+dirs[d][0], j+dirs[d][1], sits, maze, lastSpot)
		}
		if helper.InsideC(maze, i+dirs[d][0], j+dirs[d][1]) {
			lastSpot[i+dirs[d][0]][j+dirs[d][1]] = t
		}
		dfsEnd(i, j, mod4(d+1), sc+1000)
		dfsEnd(i, j, mod4(d-1), sc+1000)
		return false
	}
	dfsEnd(dI, dJ, 1, 0)

	res = len(sits) + 1

	return strconv.Itoa(res), nil
}

func backTrack(i, j int, sits map[[2]int]struct{}, maze [][]byte, lastSpot [][][2]int) {
	for maze[i][j] != 'S' {
		sits[[2]int{i, j}] = struct{}{}
		i, j = lastSpot[i][j][0], lastSpot[i][j][1]
	}
}
