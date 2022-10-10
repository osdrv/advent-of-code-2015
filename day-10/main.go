package main

import (
	"bytes"
	"os"
	"strconv"
)

func lookAndSay(s string, reps int) string {
	for r := 0; r < reps; r++ {
		var b bytes.Buffer
		ptr := 0
		for ptr < len(s) {
			from := ptr
			for ptr < len(s) && s[ptr] == s[from] {
				ptr++
			}
			b.WriteString(strconv.Itoa(ptr - from))
			b.WriteByte(s[from])
		}
		s = b.String()
	}
	return s
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	REPS := 50

	for _, line := range lines {
		las := lookAndSay(line, REPS)
		printf("line: %s, lookAndSay: %s (%d)", line, las, len(las))
	}
}
