package main

import (
	"os"
	"sort"
	"strings"
)

func parseDims(s string) (int, int, int) {
	ss := strings.SplitN(s, "x", 3)
	return parseInt(ss[0]), parseInt(ss[1]), parseInt(ss[2])
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	square := 0
	rib := 0

	lines := readLines(f)
	for _, line := range lines {
		l, w, h := parseDims(line)
		dims := []int{l, w, h}
		sort.Ints(dims)
		s1, s2, s3 := l*w, w*h, l*h

		sq := 2*s1 + 2*s2 + 2*s3 + min(s1, min(s2, s3))
		rb := 2*(dims[0]+dims[1]) + l*w*h

		printf("box: %s, sq: %d, rb: %d", line, sq, rb)

		square += sq
		rib += rb
	}

	printf("total square: %d", square)
	printf("total ribbon: %d", rib)
}
