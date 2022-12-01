package day02

import "strconv"

// https://adventofcode.com/2021/day/2

func SubmarinePos(directions string) (string, error) {
	var cmd byte
	var depth, dist, cur int
	// Looping over directions
	for i := 0; i < len(directions); i++ {

		// Cmd is the first letter of the command
		cmd = directions[i]
		// Calculating the value of the command
		for cur = 0; i < len(directions) && directions[i] != '\n'; i++ {
			if directions[i] >= '0' && directions[i] <= '9' {
				cur = cur*10 + int(directions[i]-'0')
			}
		}

		// Adjusting depth and distance according to the command and value
		switch cmd {
		case 'f':
			dist += cur
		case 'd':
			depth += cur
		case 'u':
			depth -= cur
		}
	}
	return strconv.Itoa(depth * dist), nil
}

func SubmarinePosAim(directions string) (string, error) {
	var cmd byte
	var depth, dist, aim, cur int
	// Looping over directions
	for i := 0; i < len(directions); i++ {

		// Cmd is the first letter of the command
		cmd = directions[i]
		// Calculating the value of the command
		for cur = 0; i < len(directions) && directions[i] != '\n'; i++ {
			if directions[i] >= '0' && directions[i] <= '9' {
				cur = cur*10 + int(directions[i]-'0')
			}
		}

		// Adjusting depth, distance and aim according to the command and value
		switch cmd {
		case 'f':
			dist += cur
			depth += aim * cur
		case 'd':
			aim += cur
		case 'u':
			aim -= cur
		}
	}
	return strconv.Itoa(depth * dist), nil
}
