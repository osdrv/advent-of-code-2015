package main

import (
	"os"
)

func nextGen(field [][]int, static map[[2]int]int) [][]int {
	height, width := sizeNumField(field)
	fieldcp := makeIntField(height, width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if v, ok := static[[2]int{y, x}]; ok {
				fieldcp[y][x] = v
				continue
			}
			nbs := 0
			for _, step := range STEPS8 {
				ny, nx := y+step[0], x+step[1]
				if ny < 0 || nx < 0 || ny >= height || nx >= width {
					continue
				}
				nbs += field[ny][nx]
			}
			if field[y][x] == 1 {
				if nbs == 2 || nbs == 3 {
					fieldcp[y][x] = 1
				}
			} else {
				if nbs == 3 {
					fieldcp[y][x] = 1
				}
			}
		}
	}
	return fieldcp
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	field := makeNumField[int](len(lines), len(lines[0]))
	for y, line := range lines {
		for x := 0; x < len(lines[y]); x++ {
			if line[x] == '#' {
				field[y][x] = 1
			}
		}
	}

	field2 := copyNumField(field)

	static := make(map[[2]int]int)

	for i := 0; i < 100; i++ {
		field = nextGen(field, static)
	}

	debugf(printNumFieldWithSubs(field, "", map[int]string{
		0: ".",
		1: "#",
	}))

	printf("field count: %d", countNumField(field))

	height, width := sizeNumField(field2)
	static2 := map[[2]int]int{
		{0, 0}:                  1,
		{0, width - 1}:          1,
		{height - 1, 0}:         1,
		{height - 1, width - 1}: 1,
	}

	for p, v := range static2 {
		field2[p[0]][p[1]] = v
	}

	for i := 0; i < 100; i++ {
		field2 = nextGen(field2, static2)
	}

	debugf(printNumFieldWithSubs(field2, "", map[int]string{
		0: ".",
		1: "#",
	}))

	printf("field2 count: %d", countNumField(field2))
}
