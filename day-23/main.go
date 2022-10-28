package main

import (
	"os"
)

type Instr struct {
	code string
	reg  byte
	arg  int
}

func parseInstr(s string) Instr {
	var i Instr
	i.code = s[0:3]
	debugf("parsing instr: %s", s)
	if isAlpha(s[4]) {
		i.reg = s[4]
		if len(s) > 5 {
			i.arg = parseInt(s[7:])
		}
	} else {
		i.arg = parseInt(s[4:])
	}
	return i
}

func execInstr(instrs []Instr, pc int, regs map[byte]int) int {
	i := instrs[pc]
	switch i.code {
	case "hlf":
		regs[i.reg] /= 2
		return pc + 1
	case "tpl":
		regs[i.reg] *= 3
		return pc + 1
	case "inc":
		regs[i.reg]++
		return pc + 1
	case "jmp":
		return pc + i.arg
	case "jie":
		if regs[i.reg]%2 == 0 {
			return pc + i.arg
		}
		return pc + 1
	case "jio":
		if regs[i.reg] == 1 {
			return pc + i.arg
		}
		return pc + 1
	default:
		panic("unsupported instruction" + i.code)
	}
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)
	instrs := make([]Instr, 0, len(lines))

	for _, line := range lines {
		instrs = append(instrs, parseInstr(line))
	}

	pc := 0
	regs := make(map[byte]int)

	printf("instrs: %+v", instrs)

	for pc < len(instrs) {
		pc = execInstr(instrs, pc, regs)
	}

	printf("regs(part1): %+v", regs)

	pc = 0
	regs = map[byte]int{'a': 1}
	for pc < len(instrs) {
		pc = execInstr(instrs, pc, regs)
	}

	printf("regs(part2): %+v", regs)
}
