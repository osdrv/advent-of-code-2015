package main

import (
	"os"
)

func traverse(s string) map[Point2]int {
	pth := make(map[Point2]int)

	pos := Point2{0, 0}
	pth[pos] = 1
	for _, ch := range s {
		switch ch {
		case '>':
			pos.x++
		case '<':
			pos.x--
		case '^':
			pos.y--
		case 'v':
			pos.y++
		}
		pth[pos]++
	}
	return pth
}

func traverse2(s string) map[Point2]int {
	pth := make(map[Point2]int)
	poss := map[bool]Point2{
		true:  {0, 0},
		false: {0, 0},
	}
	pth[Point2{0, 0}] = 2
	turn := true
	for _, ch := range s {
		pos := poss[turn]
		switch ch {
		case '>':
			pos.x++
		case '<':
			pos.x--
		case '^':
			pos.y--
		case 'v':
			pos.y++
		}
		pth[pos]++
		poss[turn] = pos
		turn = !turn
	}
	return pth
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	for _, line := range lines {
		prs := traverse(line)
		prs2 := traverse2(line)
		printf("line: %s, unique houses: %d", line, len(prs))
		printf("traverse2: %d", len(prs2))
	}
}
