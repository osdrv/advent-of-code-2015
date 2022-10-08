package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func main() {
	input := "iwrupvqb"
	//input := "abcdef"
	//input := "pqrstuv"
	i := 0
	for {
		data := []byte(input + strconv.Itoa(i))
		hh := md5.Sum(data)
		if hex.EncodeToString(hh[:])[:6] == "000000" {
			printf("found hash %x at iteration %d", hh, i)
			break
		}
		i++
	}
}
