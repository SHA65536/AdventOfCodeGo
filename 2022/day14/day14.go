package day14

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/14

func HowMuchSand(input string) (string, error) {
	var res int
	var board = map[[2]int]bool{}

	// Loading lines
	var lines, floor = LoadLines(input)

	// Drawing lines
	for i := range lines {
		DrawLine(board, lines[i])
	}

MajorLoop:
	for res = 0; true; res++ {
		var cont = true
		var cur = [2]int{0, 500}
		board[cur] = true
		for cont {
			// If we reached the bottom
			if cur[0] == floor {
				break MajorLoop
			}
			next := [2]int{cur[0] + 1, cur[1]}
			if !board[next] {
				// Going down
				board[cur] = false
				cur = next
				board[cur] = true
				continue
			}
			next = [2]int{cur[0] + 1, cur[1] - 1}
			if !board[next] {
				// Going left
				board[cur] = false
				cur = next
				board[cur] = true
				continue
			}
			next = [2]int{cur[0] + 1, cur[1] + 1}
			if !board[next] {
				// Going right
				board[cur] = false
				cur = next
				board[cur] = true
				continue
			}
			cont = false
		}
	}

	return strconv.Itoa(res), nil
}

func SandWithFloor(input string) (string, error) {
	var res int
	var board = map[[2]int]bool{}
	var origin = [2]int{0, 500}

	// Loading lines
	var lines, floor = LoadLines(input)

	floor += 2

	// Drawing lines
	for i := range lines {
		DrawLine(board, lines[i])
	}

	DrawLine(board, [][2]int{{floor, -(floor * 5)}, {floor, floor * 5}})

	for res = 0; !board[origin]; res++ {
		var cont = true
		var cur = origin
		board[cur] = true
		for cont {
			next := [2]int{cur[0] + 1, cur[1]}
			if !board[next] {
				// Going down
				board[cur] = false
				cur = next
				board[cur] = true
				continue
			}
			next = [2]int{cur[0] + 1, cur[1] - 1}
			if !board[next] {
				// Going left
				board[cur] = false
				cur = next
				board[cur] = true
				continue
			}
			next = [2]int{cur[0] + 1, cur[1] + 1}
			if !board[next] {
				// Going right
				board[cur] = false
				cur = next
				board[cur] = true
				continue
			}
			cont = false
		}
	}

	return strconv.Itoa(res), nil
}

func LoadLines(input string) ([][][2]int, int) {
	var maxY int
	var lines [][][2]int
	// Loading lines
	for _, line := range strings.Split(input, "\n") {
		var coords [][2]int
		for i := 0; i < len(line); i += 4 {
			var curX, curY int
			for ; i < len(line) && isDigit(line[i]); i++ {
				curX = curX*10 + int(line[i]-'0')
			}
			for i++; i < len(line) && isDigit(line[i]); i++ {
				curY = curY*10 + int(line[i]-'0')
			}
			if curY > maxY {
				maxY = curY
			}
			coords = append(coords, [2]int{curY, curX})
		}
		lines = append(lines, coords)
	}
	return lines, maxY
}

func DrawLine(board map[[2]int]bool, line [][2]int) {
	for i := 1; i < len(line); i++ {
		var direction [2]int
		if line[i-1][0] == line[i][0] {
			// Same X
			direction = [2]int{0, 1}
			if line[i-1][1] > line[i][1] {
				direction[1] = -1
			}
		} else {
			// Same Y
			direction = [2]int{1, 0}
			if line[i-1][0] > line[i][0] {
				direction[0] = -1
			}
		}
		// Plotting
		for line[i-1] != line[i] {
			board[line[i-1]] = true
			line[i-1][0] += direction[0]
			line[i-1][1] += direction[1]
		}
		board[line[i-1]] = true
	}
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
