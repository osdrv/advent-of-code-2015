package main

import (
	"os"
	"strings"
)

func parseSue(s string) (string, map[string]int) {
	ss := strings.SplitN(s, ": ", 2)
	sue := ss[0]
	attrs := make(map[string]int)
	sss := strings.Split(ss[1], ", ")
	for _, attr := range sss {
		kv := strings.SplitN(attr, ": ", 2)
		attrs[kv[0]] = parseInt(kv[1])
	}
	return sue, attrs
}

func matchFilters(attrs, filters map[string]int) bool {
	for k, v := range attrs {
		if filters[k] != v {
			return false
		}
	}
	return true
}

func matchFilters2(attrs, filters map[string]int) bool {
	for k, v := range attrs {
		switch k {
		case "cats", "trees":
			if !(v > filters[k]) {
				return false
			}
		case "pomeranians", "goldfish":
			if !(v < filters[k]) {
				return false
			}
		default:
			if filters[k] != v {
				return false
			}
		}
	}
	return true
}

func main() {
	filter := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	for _, line := range lines {
		sue, attrs := parseSue(line)
		debugf("sue: %s, attrs: %+v", sue, attrs)
		if matchFilters(attrs, filter) {
			printf("the Sue: %s", sue)
		}
		if matchFilters2(attrs, filter) {
			printf("the Sue2: %s", sue)
		}
	}
}
