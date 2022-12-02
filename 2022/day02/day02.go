package day02

import (
	"strconv"
)

// https://adventofcode.com/2022/day/2

func RockPaperScissors(guide string) (string, error) {
	var score int
	// Looping over the guide
	for i := 0; i < len(guide); i += 4 {
		you := int(guide[i+2] - 'X') // Your pick
		him := int(guide[i] - 'A')   // Their pick
		score += you + 1
		if you == him { // Same shape draws
			score += 3
		} else if (you+1)%3 == him { // Shape below loses (looping)
			score += 0
		} else if (you+2)%3 == him { // Shape above wins (looping)
			score += 6
		}
	}
	return strconv.Itoa(score), nil
}

func RockPaperScissorsModified(guide string) (string, error) {
	var score int
	// Looping over the guide
	for i := 0; i < len(guide); i += 4 {
		him := int(guide[i] - 'A')   // Their pick
		res := int(guide[i+2] - 'X') // The result
		if res == 0 {                // Lose
			score += ((him + 2) % 3) + 1 // Shape below loses (looping)
		} else if res == 1 { // Draw
			score += him + 1 + 3 // Same shape draws
		} else if res == 2 { // Win
			score += ((him + 1) % 3) + 1 // Shape above wins (looping)
			score += 6
		}
	}
	return strconv.Itoa(score), nil
}
