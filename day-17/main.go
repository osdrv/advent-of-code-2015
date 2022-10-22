package main

import (
	"os"
)

var (
	mem  = make(map[[2]int]int)
	cnts = make(map[int]int)
)

func computeCombs(inventory []int, ix int, rem int, cnt int) int {
	if rem == 0 {
		cnts[cnt]++
		return 1
	}
	if ix >= len(inventory) {
		return 0
	}
	//if v, ok := mem[[2]int{ix, rem}]; ok {
	//	return v
	//}
	res := computeCombs(inventory, ix+1, rem, cnt)
	if inventory[ix] <= rem {
		res += computeCombs(inventory, ix+1, rem-inventory[ix], cnt+1)
	}
	//mem[[2]int{ix, rem}] = res
	return res
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)
	inventory := make([]int, 0, len(lines))
	for _, line := range lines {
		inventory = append(inventory, parseInt(line))
	}

	printf("inventory: %+v", inventory)

	combs := computeCombs(inventory, 0, 150, 0)
	printf("combs: %d", combs)

	minN := ALOT
	minCnt := 0
	for n, cnt := range cnts {
		if n < minN {
			minN = n
			minCnt = cnt
		}
	}

	printf("cnts: %+v", cnts)

	printf("mincnt: %d", minCnt)
}
