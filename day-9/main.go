package main

import (
	"os"
)

//London to Dublin = 464

func parseDist(s string) (string, string, int) {
	var from, to string
	var dist int
	ptr := 0
	from, ptr = readWord(s, ptr)
	ptr = eatWhitespace(s, ptr)
	_, ptr = readStr(s, ptr, "to")
	ptr = eatWhitespace(s, ptr)
	to, ptr = readWord(s, ptr)
	ptr = eatWhitespace(s, ptr)
	ptr = consume(s, ptr, '=')
	ptr = eatWhitespace(s, ptr)
	dist, ptr = readInt(s, ptr)

	return from, to, dist
}

func addEdge(graph map[string]map[string]int, from string, to string, dist int) {
	if _, ok := graph[from]; !ok {
		graph[from] = make(map[string]int)
	}
	graph[from][to] = dist
}

func findShortestPath(graph map[string]map[string]int) int {

	var traverse func(string, map[string]bool) (int, bool)
	traverse = func(cur string, visited map[string]bool) (int, bool) {
		md, found := ALOT, false
		visited[cur] = true
		if len(visited) == len(graph) {
			delete(visited, cur)
			return 0, true
		}
		for next := range graph[cur] {
			if !visited[next] {
				dd, ok := traverse(next, visited)
				if ok {
					md = min(md, graph[cur][next]+dd)
					found = true
				}
			}
		}
		delete(visited, cur)
		return md, found
	}

	mindist := ALOT
	for vx := range graph {
		if dist, ok := traverse(vx, make(map[string]bool)); ok {
			mindist = min(mindist, dist)
		}
	}

	return mindist
}

func findLongestPath(graph map[string]map[string]int) int {

	var traverse func(string, map[string]bool) (int, bool)
	traverse = func(cur string, visited map[string]bool) (int, bool) {
		md, found := -ALOT, false
		visited[cur] = true
		if len(visited) == len(graph) {
			delete(visited, cur)
			return 0, true
		}
		for next := range graph[cur] {
			if !visited[next] {
				dd, ok := traverse(next, visited)
				if ok {
					md = max(md, graph[cur][next]+dd)
					found = true
				}
			}
		}
		delete(visited, cur)
		return md, found
	}

	maxdist := -ALOT
	for vx := range graph {
		if dist, ok := traverse(vx, make(map[string]bool)); ok {
			maxdist = max(maxdist, dist)
		}
	}

	return maxdist
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	graph := make(map[string]map[string]int)

	for _, line := range lines {
		from, to, dist := parseDist(line)
		printf("from: %s, to: %s, dist: %d", from, to, dist)
		addEdge(graph, from, to, dist)
		addEdge(graph, to, from, dist)
	}

	dmin := findShortestPath(graph)
	printf("the shortest path: %d", dmin)

	dmax := findLongestPath(graph)
	printf("the longest path: %d", dmax)
}
