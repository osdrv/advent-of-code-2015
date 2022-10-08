package main

import (
	"os"
)

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	for _, line := range lines {
		lvl := 0
		for ptr := 0; ptr < len(line); ptr++ {
			if line[ptr] == '(' {
				lvl++
			} else {
				lvl--
			}
			if lvl == -1 {
				printf("lvl -1 at %d", ptr+1)
			}
		}
		printf("lvl: %d", lvl)
	}
}
