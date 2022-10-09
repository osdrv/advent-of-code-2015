package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ASSIGN = "->"
	NOT    = "NOT"
	OR     = "OR"
	AND    = "AND"
	RSHIFT = "RSHIFT"
	LSHIFT = "LSHIFT"
)

func parseInstr(s string) [4]string {
	var instr [4]string

	ss := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' '
	})

	if ss[0] == NOT {
		instr[0] = ss[0]
		instr[1] = ss[1]
		instr[2] = ss[3]
	} else if ss[1] == OR || ss[1] == AND {
		instr[0] = ss[1]
		instr[1] = ss[0]
		instr[2] = ss[2]
		instr[3] = ss[4]
	} else if ss[1] == RSHIFT || ss[1] == LSHIFT {
		instr[0] = ss[1]
		instr[1] = ss[0]
		instr[2] = ss[2]
		instr[3] = ss[4]
	} else if ss[1] == ASSIGN {
		instr[0] = ss[1]
		instr[1] = ss[0]
		instr[2] = ss[2]
	} else {
		panic(fmt.Sprintf("can not parse instr: %q", s))
	}

	return instr
}

func resolve(regs map[string]uint16, arg string) uint16 {
	if looksLikeNumber(arg) {
		return uint16(parseInt(arg))
	}
	return regs[arg]
}

func evalInstr(regs map[string]uint16, instr [4]string) (string, uint16) {
	switch instr[0] {
	case ASSIGN:
		return instr[2], resolve(regs, instr[1])
	case NOT:
		return instr[2], ^(resolve(regs, instr[1]))
	case AND:
		return instr[3], resolve(regs, instr[1]) & resolve(regs, instr[2])
	case OR:
		return instr[3], resolve(regs, instr[1]) | resolve(regs, instr[2])
	case LSHIFT:
		return instr[3], resolve(regs, instr[1]) << int(resolve(regs, instr[2]))
	case RSHIFT:
		return instr[3], resolve(regs, instr[1]) >> int(resolve(regs, instr[2]))
	default:
		panic(fmt.Sprintf("could not interpret instruction: %+v", instr))
	}
}

func eval(regs map[string]uint16, instrs [][4]string) {
	graph := computeGraph(instrs)
	debugf("graph: %+v", graph)

	assigns := make([][4]string, 0, 1)
	for _, instr := range instrs {
		if instr[0] == ASSIGN {
			assigns = append(assigns, instr)
		}
	}

	var spawnEval func([4]string)
	spawnEval = func(instr [4]string) {
		debugf("execute instr %+v", instr)
		ptr, val := evalInstr(regs, instr)
		if v, ok := regs[ptr]; ok && v == val {
			return
		}
		regs[ptr] = val
		for _, ii := range graph[ptr] {
			spawnEval(ii)
		}
	}

	for _, asgn := range assigns {
		spawnEval(asgn)
	}
}

func computeGraph(instrs [][4]string) map[string][][4]string {
	graph := make(map[string][][4]string)

	for i := 0; i < len(instrs); i++ {
		instr := instrs[i]
		if instr[0] == ASSIGN {
			continue
		}
		for ix := 1; ix <= 2; ix++ {
			if addr := instr[ix]; !looksLikeNumber(addr) {
				if _, ok := graph[addr]; !ok {
					graph[addr] = make([][4]string, 0, 1)
				}
				graph[addr] = append(graph[addr], instr)
			}
		}
	}
	return graph
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	regs := make(map[string]uint16)

	instrs := make([][4]string, 0, len(lines))

	for _, line := range lines {
		instrs = append(instrs, parseInstr(line))
	}

	eval(regs, instrs)

	printf("regs: %+v", regs)
	printf("regs[a]: %d", regs["a"])

	regs2 := make(map[string]uint16)
	// override b
	for i := 0; i < len(instrs); i++ {
		if instrs[i][0] == ASSIGN && instrs[i][2] == "b" {
			instrs[i][1] = strconv.Itoa(int(regs["a"]))
		}
	}

	eval(regs2, instrs)

	printf("regs2: %+v", regs2)
	printf("regs2[a]: %d", regs2["a"])
}
