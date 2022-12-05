package day12

import (
	"strconv"
)

// https://adventofcode.com/2021/day/12

func UniquePaths(caves string) (string, error) {
	var res int
	var graph = map[string][]string{}
	// Loading the graph
	for i := 0; i < len(caves); i++ {
		var src, dst string
		var start int
		for start = i; caves[i] != '-'; i++ {
		}
		src = caves[start:i]
		i++
		for start = i; i < len(caves) && caves[i] != '\n'; i++ {
		}
		dst = caves[start:i]
		// Connections are both ways
		graph[src] = InsertIfNotIn(graph[src], dst)
		graph[dst] = InsertIfNotIn(graph[dst], src)
	}
	TraverseGraph(graph, map[string]bool{}, "start", &res)
	return strconv.Itoa(res), nil
}

func TraverseGraph(graph map[string][]string, visited map[string]bool, cur string, res *int) {
	if !IsUppercase(cur) {
		visited[cur] = true
	}
	options := graph[cur]
	for i := range options {
		// If we got to the end
		if options[i] == "end" {
			(*res)++
			continue
		}
		// Traverse if not visited
		if !visited[options[i]] {
			TraverseGraph(graph, visited, options[i], res)
		}
	}
	// Unmark this for others
	visited[cur] = false
}

func UniquePathsTwo(caves string) (string, error) {
	var res int
	var graph = map[string][]string{}
	// Loading the graph
	for i := 0; i < len(caves); i++ {
		var src, dst string
		var start int
		for start = i; caves[i] != '-'; i++ {
		}
		src = caves[start:i]
		i++
		for start = i; i < len(caves) && caves[i] != '\n'; i++ {
		}
		dst = caves[start:i]
		// Connections are both ways
		graph[src] = InsertIfNotIn(graph[src], dst)
		graph[dst] = InsertIfNotIn(graph[dst], src)
	}
	TraverseGraphTwo(graph, map[string]int{}, "start", &res)
	return strconv.Itoa(res), nil
}

func TraverseGraphTwo(graph map[string][]string, visited map[string]int, cur string, res *int) {
	if !IsUppercase(cur) {
		visited[cur]++
	}
	options := graph[cur]
	for i := range options {
		// If we got to the end
		if options[i] == "end" {
			(*res)++
			continue
		}
		// Can't go back to the start
		if options[i] == "start" {
			continue
		}
		// We can only visit one cave twice
		if HasTwo(visited) {
			if visited[options[i]] < 1 {
				TraverseGraphTwo(graph, visited, options[i], res)
			}
		} else {
			if visited[options[i]] < 2 {
				TraverseGraphTwo(graph, visited, options[i], res)
			}
		}
	}
	// Unmark this for others
	visited[cur]--
}

// Inserts a string into a list if it's not already in there
func InsertIfNotIn(in []string, val string) []string {
	for i := range in {
		if in[i] == val {
			return in
		}
	}
	in = append(in, val)
	return in
}

// Checks if given string is all uppercase
func IsUppercase(str string) bool {
	for i := range str {
		if str[i] < 'A' || str[i] > 'Z' {
			return false
		}
	}
	return true
}

// Checks if a map has a two
func HasTwo(m map[string]int) bool {
	for _, v := range m {
		if v == 2 {
			return true
		}
	}
	return false
}
