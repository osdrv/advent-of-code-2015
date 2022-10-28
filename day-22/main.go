package main

import (
	"bytes"
	"fmt"
)

type Spell struct {
	name     string
	cost     int
	duration int

	armor  int
	damage int
	heal   int
	mana   int
}

type Player struct {
	hp    int
	mana  int
	armor int
	sp    [5]int
}

type Boss struct {
	hp     int
	damage int
}

var (
	Spells = []Spell{
		{name: "Magic Missile", cost: 53, damage: 4},
		{name: "Drain", cost: 73, damage: 2, heal: 2},
		{name: "Shield", cost: 113, armor: 7, duration: 6},
		{name: "Poison", cost: 173, damage: 3, duration: 6},
		{name: "Recharge", cost: 229, mana: 101, duration: 5},
	}
)

type QueueItem struct {
	p Player
	b Boss
	m int
}

func applySpell(p *Player, b *Boss, six int) {
	s := Spells[six]
	var buf bytes.Buffer
	buf.WriteString("Spell " + s.name + " deals")
	if s.damage > 0 {
		buf.WriteString(fmt.Sprintf(" %d damage", s.damage))
		b.hp -= s.damage
	}
	if s.heal > 0 {
		buf.WriteString(fmt.Sprintf(" %d healing", s.heal))
		p.hp += s.heal
	}
	if s.armor > 0 {
		buf.WriteString(fmt.Sprintf(" %d armor", s.armor))
		p.armor = s.armor
	}
	if s.mana > 0 {
		buf.WriteString(fmt.Sprintf(" %d mana", s.mana))
		p.mana += s.mana
	}
	p.sp[six] = max(0, p.sp[six]-1)
	if p.sp[six] == 0 {
		if s.armor > 0 {
			p.armor = 0
		}
	}
	if s.duration > 0 {
		buf.WriteString(fmt.Sprintf(" with remaining duration %d turns", p.sp[six]))
	}
	debugf(buf.String())
}

func applySpells(p *Player, b *Boss) {
	for i := 0; i < len(p.sp); i++ {
		if p.sp[i] == 0 {
			continue
		}
		applySpell(p, b, i)
	}
}

func play(p0 Player, b0 Boss, dd int) int {
	q := make([]QueueItem, 0, 1)

	q = append(q, QueueItem{p: p0, b: b0, m: 0})

	minMana := ALOT

	var head QueueItem
	for len(q) > 0 {
		head, q = q[0], q[1:]
		p, b := head.p, head.b
		debugf("------------------")
		debugf("Player has %d hp, %d mana, %d armor, %d mana spent", p.hp, p.mana, p.armor, head.m)
		debugf("Boss has %d hp", b.hp)

		if dd > 0 {
			debugf("Player looses %dhp (part2)", dd)
			p.hp -= dd
			if p.hp <= 0 {
				continue
			}
		}

		applySpells(&p, &b)
		if b.hp <= 0 {
			debugf("Player wins with %d hp, %d mana and %d mana used", p.hp, p.mana, head.m)
			minMana = min(minMana, head.m)
			continue
		}
		// if no spells casted, this branch is over
		for i := 0; i < len(Spells); i++ {
			if p.sp[i] > 0 {
				// can not cast an already lasting spell
				continue
			}
			s := Spells[i]
			if s.cost <= p.mana {
				debugf("~~~~~~~~~~~~~~~~~~")
				p1, b1 := p, b
				debugf("Player casts %s spending %d mana", s.name, s.cost)
				p1.mana -= s.cost
				mana := head.m + s.cost
				if s.duration > 0 {
					p1.sp[i] = s.duration
				} else {
					applySpell(&p1, &b1, i)
				}
				if b1.hp <= 0 {
					debugf("Player wins with %d hp, %d mana, %d armor and %d mana used", p1.hp, p1.mana, p1.armor, mana)
					minMana = min(minMana, mana)
					continue
				}

				// emulate boss's turn
				debugf("------------------")
				debugf("Player has %d hp, %d mana, %d armor, %d mana spent", p1.hp, p1.mana, p1.armor, mana)
				debugf("Boss has %d hp", b1.hp)
				applySpells(&p1, &b1)
				if b1.hp <= 0 {
					debugf("Player wins with %d hp, %d mana and %d mana used", p1.hp, p1.mana, mana)
					minMana = min(minMana, mana)
					continue
				}
				dmg := max(b1.damage-p1.armor, 1)
				debugf("Boss deals %d damage", dmg)
				p1.hp -= dmg
				if p1.hp <= 0 {
					debugf("Boss wins, the game is over")
					continue
				}
				if mana < minMana {
					q = append(q, QueueItem{p: p1, b: b1, m: mana})
				}
			}
		}
	}

	return minMana
}

func main() {
	//p := Player{hp: 10, mana: 250}
	//b := Boss{hp: 14, damage: 8}

	p := Player{hp: 50, mana: 500}
	b := Boss{hp: 55, damage: 8}

	minMana := play(p, b, 0)

	printf("min mana(part1): %d", minMana)

	minMana2 := play(p, b, 1)

	printf("min mana(part2): %d", minMana2)
}
