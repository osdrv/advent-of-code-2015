package main

import (
	"os"
)

const (
	_ int = iota
	TURN_ON
	TURN_OFF
	TOGGLE
)

func parseInstr(s string) (int, [2]int, [2]int) {
	ptr := 0
	var op int
	if matchStr(s, ptr, "turn on") {
		_, ptr = readStr(s, ptr, "turn on")
		op = TURN_ON
	} else if matchStr(s, ptr, "turn off") {
		_, ptr = readStr(s, ptr, "turn off")
		op = TURN_OFF
	} else if matchStr(s, ptr, "toggle") {
		_, ptr = readStr(s, ptr, "toggle")
		op = TOGGLE
	} else {
		panic("wtf")
	}
	ptr = eatWhitespace(s, ptr)

	var p1, p2 [2]int

	p1[0], ptr = readInt(s, ptr)
	ptr = consume(s, ptr, ',')
	p1[1], ptr = readInt(s, ptr)
	ptr = eatWhitespace(s, ptr)
	_, ptr = readStr(s, ptr, "through")
	ptr = eatWhitespace(s, ptr)

	p2[0], ptr = readInt(s, ptr)
	ptr = consume(s, ptr, ',')
	p2[1], ptr = readInt(s, ptr)

	return op, p1, p2
}

func apply(field [][]int, instr string) {
	op, p1, p2 := parseInstr(instr)
	for i := p1[0]; i <= p2[0]; i++ {
		for j := p1[1]; j <= p2[1]; j++ {
			if op == TURN_ON {
				field[i][j] = 1
			} else if op == TURN_OFF {
				field[i][j] = 0
			} else if op == TOGGLE {
				field[i][j] ^= 1
			}
		}
	}
}

func apply2(field [][]int, instr string) {
	op, p1, p2 := parseInstr(instr)
	for i := p1[0]; i <= p2[0]; i++ {
		for j := p1[1]; j <= p2[1]; j++ {
			if op == TURN_ON {
				field[i][j]++
			} else if op == TURN_OFF {
				field[i][j] = max(0, field[i][j]-1)
			} else if op == TOGGLE {
				field[i][j] += 2
			}
		}
	}
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	field := makeNumField[int](1000, 1000)

	for _, line := range lines {
		apply(field, line)
	}

	cnt := countNumField(field)

	printf("lights count: %d", cnt)

	field = makeNumField[int](1000, 1000)

	for _, line := range lines {
		apply2(field, line)
	}

	cnt2 := countNumField(field)

	printf("lights count2: %d", cnt2)
}
