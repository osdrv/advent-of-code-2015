package main

import (
	"os"
	"strings"
)

func isContains3Vovels(s string) bool {
	v := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'u', 'o':
			v++
		}
		if v >= 3 {
			return true
		}
	}
	debugf("%s contains %d vovels", s, v)
	return false
}

func isContainsDup(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	debugf("%s contains no duplicates", s)
	return false
}

func isContains(s string, fracts []string) bool {
	for _, fr := range fracts {
		if strings.Contains(s, fr) {
			debugf("%s contains %s", s, fr)
			return true
		}
	}
	return false
}

func isNice(s string) bool {
	return isContains3Vovels(s) && isContainsDup(s) && (!isContains(s, []string{"ab", "cd", "pq", "xy"}))
}

func isContainsBiChar(s string) bool {
	reps := make(map[string]bool)
	ptr := 0
	for ptr < len(s)-1 {
		ss := s[ptr : ptr+2]
		if reps[ss] {
			return true
		}
		reps[ss] = true
		if ptr < len(s)-2 && s[ptr] == s[ptr+1] && s[ptr+1] == s[ptr+2] {
			ptr++
		}
		ptr++
	}
	debugf("%s contains no bi chars", s)
	return false
}

func isContainsRepWithSkip1(s string) bool {
	ptr := 0
	for ptr < len(s)-2 {
		if s[ptr] == s[ptr+2] {
			return true
		}
		ptr++
	}
	debugf("%s contains no reps with skip 1", s)
	return false
}

func isNice2(s string) bool {
	return isContainsBiChar(s) && isContainsRepWithSkip1(s)
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	nice := 0
	for _, line := range lines {
		if isNice2(line) {
			debugf("line %s is nice", line)
			nice++
		}
	}

	printf("nice strings: %d", nice)
}
