package day10

import (
	"sort"
	"strconv"
)

// https://adventofcode.com/2021/day/10

var Score = map[byte]int{
	')': 3, ']': 57, '}': 1197, '>': 25137, // Scores for corrupt lines
	'(': 1, '[': 2, '{': 3, '<': 4, // Scores for incomplete lines
}

var Closes = map[byte]byte{
	')': '(', ']': '[', '}': '{', '>': '<',
}

func CorruptSyntax(syntax string) (string, error) {
	var res int
	for i := 0; i < len(syntax); i++ {
		var found bool
		var stack = BracketStack{} // Stack with preiviously opened brackets
		for i < len(syntax) && syntax[i] != '\n' {
			if found { // If line is corrupt, skip to next line
				i++
				continue
			}
			if val, ok := Closes[syntax[i]]; ok { // If closing bracket
				// Can't close an empty stack, make sure the last bracket was same type
				if !(len(stack) > 0 && stack.Pop() == val) {
					// Adding to score
					res += Score[syntax[i]]
					found = true
				}
			} else { // If opening bracket, add to stack
				stack.Push(syntax[i])
			}
			i++
		}
	}
	return strconv.Itoa(res), nil
}

func IncompleteSyntax(syntax string) (string, error) {
	var scores []uint64
	for i := 0; i < len(syntax); i++ {
		var found bool
		var stack = BracketStack{} // Stack with preiviously opened brackets
		for i < len(syntax) && syntax[i] != '\n' {
			if found { // If line is corrupt, skip to next line
				i++
				continue
			}
			// Checking if corrupt
			if val, ok := Closes[syntax[i]]; ok {
				if !(len(stack) > 0 && stack.Pop() == val) {
					found = true
				}
			} else {
				stack.Push(syntax[i])
			}
			i++
		}
		// If not corrupt
		if !found {
			var res uint64
			// Sum up the opened brackets according to the score guide
			for j := len(stack) - 1; j >= 0; j-- {
				res = res*5 + uint64(Score[stack[j]])
			}
			scores = append(scores, res)
		}
	}
	// Sort the scores
	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	// Take median
	return strconv.FormatUint(scores[len(scores)/2], 10), nil
}

type BracketStack []byte

func (b *BracketStack) Push(val byte) {
	(*b) = append(*b, val)
}

func (b *BracketStack) Pop() byte {
	ret := (*b)[len(*b)-1]
	(*b) = (*b)[:len(*b)-1]
	return ret
}
