package main

import (
	"os"
	"strings"
)

type Ingridient struct {
	name                    string
	cpt, dur, flv, txt, cal int
}

func parseIngr(s string) *Ingridient {
	ingr := &Ingridient{}
	ss := strings.SplitN(s, ": ", 2)
	ingr.name = ss[0]
	cs := strings.Split(ss[1], ", ")
	for _, c := range cs {
		cc := strings.SplitN(c, " ", 2)
		cn, amnt := cc[0], parseInt(cc[1])
		switch cn {
		case "capacity":
			ingr.cpt = amnt
		case "durability":
			ingr.dur = amnt
		case "flavor":
			ingr.flv = amnt
		case "texture":
			ingr.txt = amnt
		case "calories":
			ingr.cal = amnt
		default:
			panic("wtf")
		}
	}
	debugf("ingr: %+v", ingr)
	return ingr
}

func computeScore(ingrs []*Ingridient, shares []int) int {
	assert(len(ingrs) == len(shares), "the number of ingridients must match the number of shares")
	var cpt, dur, flv, txt int
	for ix, ingr := range ingrs {
		cpt += shares[ix] * ingr.cpt
		dur += shares[ix] * ingr.dur
		flv += shares[ix] * ingr.flv
		txt += shares[ix] * ingr.txt
	}
	return max(0, cpt) * max(0, dur) * max(0, flv) * max(0, txt)
}

func computeCalories(ingrs []*Ingridient, shares []int) int {
	var cal int
	for ix, ingr := range ingrs {
		cal += shares[ix] * ingr.cal
	}
	return cal
}

func computeShares(n, maxN int) [][]int {
	if n == 1 {
		return [][]int{{maxN}}
	}
	res := make([][]int, 0, 1)
	for i := 1; i <= maxN-(n-1); i++ {
		for _, sh := range computeShares(n-1, maxN-i) {
			res = append(res, append([]int{i}, sh...))
		}
	}
	return res
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)
	ingrs := make([]*Ingridient, 0, len(lines))

	for _, line := range lines {
		ingrs = append(ingrs, parseIngr(line))
	}

	shs := computeShares(len(ingrs), 100)
	debugf("%+v", shs)

	maxscore := -ALOT
	for _, sh := range shs {
		if score := computeScore(ingrs, sh); score > maxscore {
			maxscore = score
		}
	}

	printf("maxscore: %d", maxscore)

	maxscore2 := -ALOT
	for _, sh := range shs {
		if computeCalories(ingrs, sh) != 500 {
			continue
		}
		if score := computeScore(ingrs, sh); score > maxscore2 {
			maxscore2 = score
		}
	}

	printf("maxscore2: %d", maxscore2)
}
