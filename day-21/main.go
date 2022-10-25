package main

type Player struct {
	hp, damage, armor int
}

func NewPlayer(hp, damage, armor int) *Player {
	return &Player{hp, damage, armor}
}

func (p *Player) IsAlive() bool {
	return p.hp > 0
}

func (p *Player) HitBy(dmg int) {
	p.hp -= dmg
}

// returns true if p1 wins, false otherwise
func play(p1, p2 *Player) bool {
	at, df := p1, p2
	for at.IsAlive() && df.IsAlive() {
		dmg := max(1, at.damage-df.armor)
		df.HitBy(dmg)
		at, df = df, at
	}
	return p1.IsAlive()
}

type StoreItem struct {
	name                string
	cost, damage, armor int
}

var (
	Weapons = []StoreItem{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	}

	Armors = []StoreItem{
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
	}

	Rings = []StoreItem{
		{"Damage +1", 25, 1, 0},
		{"Damage +2", 50, 2, 0},
		{"Damage +3", 100, 3, 0},
		{"Defense +1", 20, 0, 1},
		{"Defense +2", 40, 0, 2},
		{"Defense +3", 80, 0, 3},
	}
)

func computeExtra(weapons, armors, rings []StoreItem) StoreItem {
	var extra StoreItem
	items := make([]StoreItem, 0, len(weapons)+len(armors)+len(rings))
	items = append(items, weapons...)
	items = append(items, armors...)
	items = append(items, rings...)
	for _, item := range items {
		extra.cost += item.cost
		extra.damage += item.damage
		extra.armor += item.armor
	}
	return extra
}

func chooseAny(items []StoreItem, low, high int) [][]StoreItem {
	return computeSubsetsOfLenRange(items, low, high)
}

func main() {
	minCost := ALOT
	maxCost := -ALOT
	for _, weapons := range chooseAny(Weapons, 1, 1) {
		for _, armors := range chooseAny(Armors, 0, 1) {
			for _, rings := range chooseAny(Rings, 0, 2) {
				extra := computeExtra(weapons, armors, rings)
				player := NewPlayer(100, extra.damage, extra.armor)
				boss := NewPlayer(104, 8, 1)
				if play(player, boss) {
					if extra.cost < minCost {
						printf("new min cost %d for %+v %+v %+v", extra.cost, weapons, armors, rings)
						minCost = extra.cost
					}
				} else {
					if extra.cost > maxCost {
						printf("new max cost %d for %+v %+v %+v", extra.cost, weapons, armors, rings)
						maxCost = extra.cost
					}
				}
			}
		}
	}

	printf("min cost: %d", minCost)
	printf("max cost: %d", maxCost)
}
