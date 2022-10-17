package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
)

// Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
var (
	re = regexp.MustCompile("(\\w+)\\scan\\sfly\\s(\\d+)\\skm/s\\sfor\\s(\\d+)\\sseconds\\,\\sbut\\sthen\\smust\\srest\\sfor\\s(\\d+)\\sseconds.")
)

type Reindeer struct {
	name                     string
	speed, runtime, resttime int
}

func NewReindeer(name string, speed, runtime, resttime int) *Reindeer {
	return &Reindeer{name, speed, runtime, resttime}
}

func (d *Reindeer) String() string {
	return fmt.Sprintf("deer %s[speed: %d, run time: %d, rest time: %d]", d.name, d.speed, d.runtime, d.resttime)
}

func (d *Reindeer) GetDistanceAt(t int) int {
	timespan := d.runtime + d.resttime
	dist := 0
	dist += d.speed * (t / timespan) * d.runtime
	dist += d.speed * min(t%(timespan), d.runtime)
	return dist
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	deers := make([]*Reindeer, 0, len(lines))

	for _, line := range lines {
		mtch := re.FindStringSubmatch(line)
		deers = append(deers, NewReindeer(mtch[1], parseInt(mtch[2]), parseInt(mtch[3]), parseInt(mtch[4])))
		printf("mtch: %+v", mtch)
	}

	TIME := 2503
	//TIME := 1000
	//TIME := 1000

	sort.Slice(deers, func(i, j int) bool {
		return deers[i].GetDistanceAt(TIME) > deers[j].GetDistanceAt(TIME)
	})

	printf("deers: %+v", deers)

	printf("the champion deer is %s at distance %d", deers[0].name, deers[0].GetDistanceAt(TIME))

	for _, deer := range deers {
		printf("deer: %s, dist: %d", deer.name, deer.GetDistanceAt(TIME))
	}

	// part 2

	scores := make(map[string]int)
	for t := 1; t <= TIME; t++ {
		maxdist := -ALOT
		champs := make([]string, 0, 1)
		for _, deer := range deers {
			dist := deer.GetDistanceAt(t)
			printf("At t:%d %s is at %d", t, deer.name, dist)
			if dist > maxdist {
				maxdist = dist
				champs = []string{deer.name}
			} else if dist == maxdist {
				champs = append(champs, deer.name)
			}
		}
		for _, champ := range champs {
			scores[champ]++
		}
	}

	printf("scores: %+v", scores)

	maxscore := -ALOT
	for _, score := range scores {
		if score > maxscore {
			maxscore = score
		}
	}

	printf("maxscore: %d", maxscore)
}
