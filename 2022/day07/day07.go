package day07

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/7

func ParseTerminal(term string) (string, error) {
	var res uint64
	var root = CreateStructure(term)
	calcSizes(root, &res, 100000)

	return strconv.FormatUint(res, 10), nil
}

func FindDelete(term string) (string, error) {
	var res uint64
	var root = CreateStructure(term)

	used := calcSizes(root, &res, 0)
	available := 70000000 - used
	needed := 30000000 - available

	res = ^uint64(0)

	calcForDelete(root, &res, needed)

	return strconv.FormatUint(res, 10), nil
}

func calcSizes(obj *Object, res *uint64, limit uint64) uint64 {
	var sum uint64
	if obj.File {
		return obj.Size
	}
	for k := range obj.Below {
		sum += calcSizes(obj.Below[k], res, limit)
	}
	obj.Size = sum
	if sum <= limit {
		*res += sum
	}
	return sum
}

func calcForDelete(obj *Object, res *uint64, limit uint64) {
	if obj.File {
		return
	}
	if obj.Size >= limit && obj.Size < *res {
		*res = obj.Size
	}
	for k := range obj.Below {
		calcForDelete(obj.Below[k], res, limit)
	}
}

func CreateStructure(term string) *Object {
	var root = &Object{Below: map[string]*Object{}}

	lines := strings.Split(term, "\n")[2:]
	var cur = root
	for i := range lines {
		if strings.HasPrefix(lines[i], "dir") {
			dir := &Object{
				Name:   lines[i][4:],
				Below:  map[string]*Object{},
				Parent: cur,
			}
			cur.Below[dir.Name] = dir
		} else if lines[i][0] >= '0' && lines[i][0] <= '9' {
			sections := strings.Split(lines[i], " ")
			size, _ := strconv.ParseUint(sections[0], 10, 64)
			file := &Object{
				File:   true,
				Size:   size,
				Name:   sections[1],
				Parent: cur,
			}
			cur.Below[file.Name] = file
		} else if strings.HasPrefix(lines[i], "$ cd") {
			sections := strings.Split(lines[i], " ")
			if sections[2] == ".." {
				cur = cur.Parent
			} else {
				cur = cur.Below[sections[2]]
			}
		}
	}
	return root
}

type Object struct {
	File   bool
	Size   uint64
	Name   string
	Below  map[string]*Object
	Parent *Object
}
