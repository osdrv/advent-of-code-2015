package main

import (
	"bytes"
	"os"
)

func runelen(s string) int {
	ll := 0
	ptr := 0
	ptr = consume(s, ptr, '"')
	for ptr < len(s) {
		if match(s, ptr, '\\') {
			ptr++
			if match(s, ptr, '\\') || match(s, ptr, '"') {
				ptr++
			} else if match(s, ptr, 'x') {
				ptr += 3
			}
			ll++
			continue
		} else if match(s, ptr, '"') {
			break
		} else {
			ll++
			ptr++
		}
	}
	consume(s, ptr, '"')
	return ll
}

func encoderune(s string) string {
	var b bytes.Buffer
	b.WriteByte('"')
	ptr := 0
	for ptr < len(s) {
		if s[ptr] == '"' || s[ptr] == '\\' {
			b.WriteByte('\\')
		}
		b.WriteByte(s[ptr])
		ptr++
	}
	b.WriteByte('"')
	return b.String()
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	code, inmem, enclen := 0, 0, 0

	for _, line := range lines {
		linecode, lineinmem := len(line), runelen(line)
		code += linecode
		inmem += lineinmem
		enc := encoderune(line)
		debugf("encoded: %s", enc)
		enclen += len(enc)
	}

	printf("part 1: %d - %d = %d", code, inmem, code-inmem)
	printf("part 2: %d - %d = %d", enclen, code, enclen-code)
}
