package main

func includeIcr(pwd []byte, n int) bool {
NextChar:
	for i := 0; i < len(pwd)-(n-1); i++ {
		for j := 1; j < n; j++ {
			if pwd[i+j]-pwd[i+j-1] != 1 {
				continue NextChar
			}
		}
		return true
	}
	debugf("%s is missing %d incremental", pwd, n)
	return false
}

func contains(pwd []byte, chars ...byte) bool {
	var sa [26]bool
	for _, ch := range chars {
		sa[ch-'a'] = true
	}
	for ptr := 0; ptr < len(pwd); ptr++ {
		if sa[pwd[ptr]-'a'] {
			debugf("%s contains %c", pwd, pwd[ptr])
			return true
		}
	}
	return false
}

func containsPairs(pwd []byte, n int) bool {
	pairs := make(map[[2]byte]bool)
	ptr := 1
	for ptr < len(pwd) {
		if pwd[ptr] == pwd[ptr-1] {
			pairs[[2]byte{pwd[ptr], pwd[ptr]}] = true
			if len(pairs) >= n {
				return true
			}
		}
		ptr++
	}
	debugf("%s is missing %d pairs", pwd, n)
	return false
}

func isValid(pwd []byte) bool {
	return includeIcr(pwd, 3) && !contains(pwd, 'i', 'o', 'l') && containsPairs(pwd, 2)
}

func increment(pwd []byte) {
	ptr := len(pwd) - 1
	for ptr >= 0 {
		pwd[ptr]++
		if pwd[ptr] > 'z' {
			pwd[ptr] = 'a'
			ptr--
		} else {
			break
		}
	}
}

func main() {
	input := []byte("vzbxkghb")
	for i := 0; i < 2; i++ {
		for !isValid(input) {
			increment(input)
		}
		printf("the next password is: %s", string(input))
		increment(input)
	}
}
