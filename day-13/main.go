package main

import (
	"os"
	"regexp"
)

var (
	re = regexp.MustCompile("(\\w+)\\swould\\s(gain|lose)\\s(\\d+)\\shappiness\\sunits\\sby\\ssitting\\snext\\sto\\s(\\w+)\\.")
)

func addEdge(graph map[string]map[string]int, edge []string) {
	from, gain, val, to := edge[0], edge[1], parseInt(edge[2]), edge[3]
	if _, ok := graph[from]; !ok {
		graph[from] = make(map[string]int)
	}
	if gain == "lose" {
		val *= -1
	}
	graph[from][to] = val
}

func findBestPlacement(graph map[string]map[string]int) int {
	var traverse func(cur string, visited map[string]bool) int
	var start string
	for vrtx := range graph {
		start = vrtx
		break
	}
	traverse = func(cur string, visited map[string]bool) int {
		visited[cur] = true
		defer delete(visited, cur)
		if len(visited) == len(graph) {
			return graph[cur][start] + graph[start][cur]
		}
		maxval := -ALOT
		for next := range graph[cur] {
			if visited[next] {
				continue
			}
			maxval = max(maxval, traverse(next, visited)+graph[cur][next]+graph[next][cur])
		}
		return maxval
	}

	return traverse(start, make(map[string]bool))
}

func addMyself(graph map[string]map[string]int, score int) {
	mslf := make(map[string]int)
	for vrtx := range graph {
		mslf[vrtx] = score
		graph[vrtx]["myself"] = score
	}
	graph["myself"] = mslf
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	graph := make(map[string]map[string]int)

	for _, line := range lines {
		mtch := re.FindStringSubmatch(line)
		addEdge(graph, mtch[1:])
	}

	addMyself(graph, 0)

	printf("graph: %+v", graph)

	score := findBestPlacement(graph)
	printf("best placement score: %d", score)
}
