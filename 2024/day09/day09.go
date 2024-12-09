package day09

import (
	"adventofcode/helper"
	"slices"
	"strconv"
)

func Star1(input *helper.InputReader) (string, error) {
	var res int

	var disk, _ = input.ReadAll()

	var full_disk = decompress([]byte(disk))
	defragment(full_disk)

	for i := range full_disk {
		if full_disk[i] == -1 {
			break
		}
		res += i * int(full_disk[i])
	}

	return strconv.Itoa(res), nil
}

func defragment(in []int) {
	var l, r int = 0, len(in) - 1
	for l < r {
		if in[l] == -1 && in[r] != -1 {
			in[l] = in[r]
			in[r] = -1
		}
		if in[l] != -1 {
			l++
		}
		if in[r] == -1 {
			r--
		}
	}
}

func decompress(in []byte) (out []int) {
	out = make([]int, 0, len(in)*8)
	var free bool
	for idx, char := range in {
		if free {
			for ; char > '0'; char-- {
				out = append(out, -1)
			}
		} else {
			for ; char > '0'; char-- {
				out = append(out, int(idx/2))
			}
		}
		free = !free
	}
	return out
}

type File struct {
	Id   int
	Size int
}

func (f File) isSpace() bool { return f.Id == -1 }

func Star2(input *helper.InputReader) (string, error) {
	var res int

	var disk, _ = input.ReadAll()
	var fs = make([]File, 0, len(disk))

	var free bool
	for i := range disk {
		if free {
			fs = append(fs, File{-1, int(disk[i] - '0')})
		} else {
			fs = append(fs, File{i / 2, int(disk[i] - '0')})
		}
		free = !free
	}

	for i := len(fs) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			var rNum, lSpace = fs[i], fs[j]
			if rNum.isSpace() || !lSpace.isSpace() {
				continue
			}
			if rNum.Size <= lSpace.Size {
				fs[i] = File{-1, rNum.Size}
				fs[j] = File{-1, lSpace.Size - rNum.Size}
				fs = slices.Insert(fs, j, File{rNum.Id, rNum.Size})
			}
		}
	}

	var pos int
	for _, f := range fs {
		if f.isSpace() {
			pos += f.Size
			continue
		}
		for t := f.Size; t > 0; t-- {
			res += (pos * f.Id)
			pos++
		}
	}

	return strconv.Itoa(res), nil
}

// Failed approach doing one list for spaces and one for files
/*
func Star2(input *helper.InputReader) (string, error) {
	var res int

	var disk, _ = input.ReadAll()

	var nums = make([][2]int, 0, len(disk))
	var spaces = make([]int, 0, len(disk)/2)

	var free bool
	for i := range disk {
		if free {
			spaces = append(spaces, int(disk[i]-'0'))
		} else {
			nums = append(nums, [2]int{int(disk[i] - '0'), i / 2})
		}
		free = !free
	}

	for ni := len(nums) - 1; ni > 0; ni-- {
		for si := range spaces {
			if si < ni && spaces[si] >= nums[ni][0] {
				spaces[si] -= nums[ni][0]
				spaces = slices.Insert(spaces, si, 0)
				spaces[ni] += nums[ni][0]
				if ni+1 < len(spaces) {
					spaces[ni] += spaces[ni+1]
					spaces = slices.Delete(spaces, ni+1, ni+2)
				}
				nums = slices.Insert(nums, si+1, nums[ni])
				nums = slices.Delete(nums, ni+1, ni+2)
				break
			}
		}
	}

	var pos int
	for i, num := range nums {
		for t := num[0]; t > 0; t-- {
			res += pos * num[1]
			pos++
		}
		pos += spaces[i]
	}

	return strconv.Itoa(res), nil
}
*/
