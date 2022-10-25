package main

func sumDivs(n int) int {
	sum := 1
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func sumLimitedDivs(n int) int {
	sum := 0
	from := n/50 + 1
	for i := from; i <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum * 11
}

func main() {
	lookupSum := 36000000

	//for n := lookupSum / 5; n <= lookupSum; n++ {
	//	sd := sumDivs(n)
	//	printf("inspecting n: %d, sd: %d/%d", n, sd, lookupSum/10)
	//	if sd >= lookupSum/10 {
	//		printf("house number: %d", n)
	//		break
	//	}
	//}

	for n := lookupSum / 50; n < lookupSum; n++ {
		sd := sumLimitedDivs(n)
		printf("inspecting n: %d, sd: %d/%d", n, sd, lookupSum)
		if sd >= lookupSum {
			printf("house number: %d", n)
			break
		}
	}
}

/*

1 -> 11
2 -> 11 + 22
3 -> 11 + 33
4 -> 11 + 22 + 44
5 -> 11 + 55
...
50 -> 11 + 22 + 55 + 110 + 25*11 + 50*11
      11 * (1 + 2 + 5 + 10 + 25 + 50)

51 -> 11 * (3 + 17 + 51)
100 -> 11 * (2 + 5 + 10 + 20 + 25 + 50 + 100)
102 -> 11 * (3 + 6 + 17 + 51 + 102)
150 -> 11 * (3 + 5 + 10 + 15 + 25 + 30 + 50 + 75 + 150)

*/
