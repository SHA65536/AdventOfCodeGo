package day08

import "strconv"

// https://adventofcode.com/2022/day/8

func VisibleTrees(heights string) (string, error) {
	var mat [][]int
	var visible = map[[2]int]bool{}

	// Loading trees
	for i := 0; i < len(heights); i++ {
		var cur []int
		for ; i < len(heights) && heights[i] != '\n'; i++ {
			cur = append(cur, int(heights[i]-'0'))
		}
		mat = append(mat, cur)
	}
	// Checking rows
	for i := range mat {
		var max int = -1
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] > max {
				max = mat[i][j]
				visible[[2]int{i, j}] = true
			}
		}
		max = -1
		for j := len(mat) - 1; j >= 0; j-- {
			if mat[i][j] > max {
				max = mat[i][j]
				visible[[2]int{i, j}] = true
			}
		}
	}
	// Checking cols
	for j := range mat[0] {
		var max int = -1
		for i := 0; i < len(mat)-1; i++ {
			if mat[i][j] > max {
				max = mat[i][j]
				visible[[2]int{i, j}] = true
			}
		}
		max = -1
		for i := len(mat) - 1; i >= 0; i-- {
			if mat[i][j] > max {
				max = mat[i][j]
				visible[[2]int{i, j}] = true
			}
		}

	}
	return strconv.Itoa(len(visible)), nil
}

func ScenicScore(heights string) (string, error) {
	var mat [][]int
	var maxScore int

	// Loading trees
	for i := 0; i < len(heights); i++ {
		var cur []int
		for ; i < len(heights) && heights[i] != '\n'; i++ {
			cur = append(cur, int(heights[i]-'0'))
		}
		mat = append(mat, cur)
	}

	for i := range mat {
		for j := range mat {
			var score, temp int = 1, 0
			var cur int = mat[i][j]
			// Checking right
			for r := j + 1; r < len(mat[0]); r++ {
				temp++
				if mat[i][r] >= cur {
					break
				}
			}
			score *= temp
			temp = 0
			// Checking left
			for l := j - 1; l >= 0; l-- {
				temp++
				if mat[i][l] >= cur {
					break
				}
			}
			score *= temp
			temp = 0
			// Checking down
			for d := i + 1; d < len(mat); d++ {
				temp++
				if mat[d][j] >= cur {
					break
				}
			}
			score *= temp
			temp = 0
			// Checking up
			for u := i - 1; u >= 0; u-- {
				temp++
				if mat[u][j] >= cur {
					break
				}
			}
			score *= temp
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return strconv.Itoa(maxScore), nil
}
