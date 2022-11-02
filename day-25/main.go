package main

func main() {
	n := uint64(20151125)
	mult := uint64(252533)
	mod := uint64(33554393)

	i, j := 1, 1

	ii, jj := 3010, 3019
	//ii, jj := 4, 3

	cur := 1

	for {
		debugf("next number is: %d", n)
		if i == ii && j == jj {
			printf("the code is: %d", n)
			break
		}
		if i == 1 {
			cur++
			i = cur
			j = 1
		} else {
			i--
			j++
		}
		n = (n * mult) % mod
	}
}
