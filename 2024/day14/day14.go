package day14

import (
	"adventofcode/helper"
	"fmt"
	"strconv"
)

// I added a line at the start of the input that specifies how big the area is

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var xMax, yMax int
	var line, _ = input.ReadLine()
	fmt.Sscanf(line, "%d,%d", &xMax, &yMax)

	var board [][]int = make([][]int, yMax)
	for i := range board {
		board[i] = make([]int, xMax)
	}

	for line, _ = input.ReadLine(); line != ""; line, _ = input.ReadLine() {
		var px, py, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

		px = (px + vx*100) % len(board[0])
		py = (py + vy*100) % len(board)

		if px < 0 {
			px = xMax + px
		}
		if py < 0 {
			py = yMax + py
		}

		board[py][px]++
	}

	var a, b, c, d int
	for i := range board {
		for j := range board[i] {
			if i < len(board)/2 && j < len(board[0])/2 {
				a += board[i][j]
			}
			if i > len(board)/2 && j < len(board[0])/2 {
				b += board[i][j]
			}
			if i < len(board)/2 && j > len(board[0])/2 {
				c += board[i][j]
			}
			if i > len(board)/2 && j > len(board[0])/2 {
				d += board[i][j]
			}
		}
	}

	res = a * b * c * d

	return strconv.Itoa(res), nil
}

type Bot struct {
	X, Y, Vx, Vy int
}

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var xMax, yMax int
	var line, _ = input.ReadLine()
	fmt.Sscanf(line, "%d,%d", &xMax, &yMax)

	if xMax != 101 || yMax != 103 { // Only large input has easter egg
		return "0", nil
	}

	var board [][]int = make([][]int, yMax)
	for i := range board {
		board[i] = make([]int, xMax)
	}

	var bots []*Bot

	for line, _ = input.ReadLine(); line != ""; line, _ = input.ReadLine() {
		var px, py, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

		bots = append(bots, &Bot{px, py, vx, vy})
	}

	var botMap = make(map[[2]int]bool, len(bots))
	for {
		for _, bot := range bots {
			bot.X = (bot.X + bot.Vx) % xMax
			bot.Y = (bot.Y + bot.Vy) % yMax
			if bot.X < 0 {
				bot.X += xMax
			}
			if bot.Y < 0 {
				bot.Y += yMax
			}
			botMap[[2]int{bot.Y, bot.X}] = true
		}

		res++

		if len(botMap) == len(bots) { // apparetly when they don't overlap it's the easter egg
			break
		}
		clear(botMap)
	}

	return strconv.Itoa(res), nil
}
