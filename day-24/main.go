package main

import (
	"os"
)

func weightsEqual(weights []int) bool {
	for i := 1; i < len(weights); i++ {
		if weights[i] != weights[0] {
			return false
		}
	}
	return true
}

func SplitEqlGrps(nums []int, nGrps int) [][]uint64 {
	weights := make([]int, nGrps)
	groups := make([]uint64, nGrps)

	totalWeight := 0
	for _, num := range nums {
		totalWeight += num
	}
	maxGroupWeight := totalWeight / nGrps

	var recurse func(groups []uint64, weights []int, ix int) [][]uint64
	recurse = func(groups []uint64, weights []int, ix int) [][]uint64 {
		if ix >= len(nums) {
			if weightsEqual(weights) {
				cp := make([]uint64, len(groups))
				copy(cp, groups)
				return [][]uint64{cp}
			}
			return nil
		}
		res := make([][]uint64, 0, 1)
		for i := 0; i < nGrps; i++ {
			if weights[i]+nums[ix] <= maxGroupWeight {
				groups[i] |= (1 << ix)
				weights[i] += nums[ix]
				res = append(res, recurse(groups, weights, ix+1)...)
				weights[i] -= nums[ix]
				groups[i] &= ^(1 << ix)
			}
		}
		return res
	}

	return recurse(groups, weights, 0)
}

func findBestGrp(nums []int, nGrps int) []int {
	totalWeight := 0
	for _, num := range nums {
		totalWeight += num
	}
	maxGroupWeight := totalWeight / nGrps

	var bestGrp uint64
	bestSize := ALOT
	bestEnt := ALOT

	// [group, sum, size, ent, ix]
	q := make([][5]int, 0, 1)
	q = append(q, [5]int{0, 0, 0, 1, len(nums) - 1})

	var head [5]int
	var group, sum, size, ent, ix int
	for len(q) > 0 {
		head, q = q[0], q[1:]
		group, sum, size, ent, ix = head[0], head[1], head[2], head[3], head[4]
		if ix < 0 {
			continue
		}
		ns := sum + nums[ix]
		gg := uint64(group) | (1 << ix)
		debugf("gg: %012b, ns: %d", gg, ns)
		if ns == maxGroupWeight {
			debugf("group: %012b", gg)
			sum += nums[ix]
			ent *= nums[ix]
			if (size < bestSize) || ((size == bestSize) && (ent < bestEnt)) {
				bestGrp = gg
				bestSize = size
				bestEnt = ent
			}
		} else if ns < maxGroupWeight {
			q = append(q, [5]int{int(gg), ns, size + 1, ent * nums[ix], ix - 1})
		}
		q = append(q, [5]int{group, sum, size, ent, ix - 1})
	}

	res := make([]int, 0, popCnt(bestGrp))

	for i := 0; i < 64; i++ {
		if bestGrp&(1<<i) > 0 {
			res = append(res, nums[i])
		}
	}

	return res
}

func popCnt(n uint64) int {
	cnt := 0
	for n > 0 {
		cnt++
		n &= n - 1
	}
	return cnt
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	nums := make([]int, 0, len(lines))
	for _, line := range lines {
		nums = append(nums, parseInt(line))
	}

	bestGrp := findBestGrp(nums, 3)
	printf("best group: %+v", bestGrp)

	ent := 1
	for _, num := range bestGrp {
		ent *= num
	}

	printf("quantum entanglement is: %d", ent)

	bestGrp2 := findBestGrp(nums, 4)
	printf("best group2: %+v", bestGrp2)

	ent2 := 1
	for _, num := range bestGrp2 {
		ent2 *= num
	}

	printf("quantum entanglement 2 is: %d", ent2)
}
