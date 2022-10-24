package main

import (
	"os"
	"strings"
)

func getUniqSubs(s string, subs map[string][]string) map[string]bool {
	uniqs := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		s1 := s[i : i+1]
		for _, sub := range subs[s1] {
			uniqs[s[:i]+sub+s[i+1:]] = true
		}
		if i < len(s)-1 {
			s2 := s[i : i+2]
			for _, sub := range subs[s2] {
				uniqs[s[:i]+sub+s[i+2:]] = true
			}
		}
	}
	return uniqs
}

func dist(a, b string) int {
	d := 0
	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			d++
		}
	}
	d += max(len(a), len(b)) - min(len(a), len(b))
	return d
}

// I'll keep it here for the history. It never worked though.
func calcMinSteps(init string, want string, subs map[string][]string) int {

	q := NewBinHeap(func(a, b string) bool {
		return dist(want, a) < dist(want, b)
	})

	q.Push(init)

	steps := make(map[string]int)
	steps[init] = 0

	enq := make(map[string]struct{})
	enq[init] = struct{}{}

	for q.Size() > 0 {
		head := q.Pop()
		ss := steps[head]
		debugf("head: %s, steps: %d", head, ss)
		delete(enq, head)
		for sub := range getUniqSubs(head, subs) {
			if sub == want {
				return ss + 1
			}
			if len(sub) > len(want) {
				continue
			}
			if v, ok := steps[sub]; !ok || v > ss+1 {
				steps[sub] = ss + 1
				if _, kk := enq[sub]; !kk {
					q.Push(sub)
				}
			}
		}
	}

	return -1
}

func isLC(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func countElements(s string) map[string]int {
	ix := 0
	cnts := make(map[string]int)
	for ix < len(s) {
		from := ix
		ix++
		if ix < len(s) && isLC(s[ix]) {
			ix++
		}
		cnts[s[from:ix]]++
		cnts["total"]++
	}
	return cnts
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	subs := make(map[string][]string)

	ix := 0
	for ix < len(lines) {
		if len(lines[ix]) == 0 {
			break
		}
		ss := strings.SplitN(lines[ix], " => ", 2)
		if _, ok := subs[ss[0]]; !ok {
			subs[ss[0]] = make([]string, 0, 1)
		}
		subs[ss[0]] = append(subs[ss[0]], ss[1])
		ix++
	}
	s := lines[len(lines)-1]

	printf("mapping: %+v", subs)

	printf("input: %s", s)

	sbs := getUniqSubs(s, subs)
	printf("unique subs: %+v", sbs)
	printf("cnt: %d", len(sbs))

	cnts := countElements(s)
	printf("counts: %+v", cnts)

	// I took this solution from here:
	// https://www.reddit.com/r/adventofcode/comments/3xflz8/comment/cy4etju/?utm_source=reddit&utm_medium=web2x&context=3
	// It is bloody awesome.
	res2 := cnts["total"] - (cnts["Rn"] + cnts["Ar"]) - 2*cnts["Y"] - 1
	printf("res2: %d", res2)
}
