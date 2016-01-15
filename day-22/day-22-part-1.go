package main

import (
	"io/ioutil"
	"math/rand"
	"regexp"
	"strconv"
)

type Spell struct {
	name string
	cost int
}

type Boss struct {
	hp, damage int
}

func (boss *Boss) damaged(damage int) {
	boss.hp -= damage
}

func (boss *Boss) dead() bool {
	return boss.hp <= 0
}

type Wizard struct {
	hp, mana, armor                            int
	shield_timer, poison_timer, recharge_timer int
	spent_mana                                 int
}

func (wizard *Wizard) damaged(damage int) {
	if damage-wizard.armor > 0 {
		wizard.hp -= damage - wizard.armor
	} else {
		wizard.hp -= 1
	}
}

func (wizard *Wizard) dead() bool {
	return wizard.hp <= 0 || wizard.mana < 53
}

func (wizard *Wizard) cast(spell *Spell, boss *Boss) {
	wizard.spent_mana += spell.cost
	wizard.mana -= spell.cost

	switch spell.name {
	case "missile":
		boss.damaged(4)
	case "drain":
		boss.damaged(2)
		wizard.hp += 2
	case "shield":
		wizard.armor = 7
		wizard.shield_timer = 6
	case "poison":
		wizard.poison_timer = 6
	case "recharge":
		wizard.recharge_timer = 5
	}
}

func (wizard *Wizard) can_cast(spell *Spell) bool {
	if spell.cost > wizard.mana {
		return false
	}

	switch spell.name {
	case "shield":
		return wizard.shield_timer == 0
	case "poison":
		return wizard.poison_timer == 0
	case "recharge":
		return wizard.recharge_timer == 0
	default:
		return true
	}
}

func (wizard *Wizard) tick(boss *Boss) {
	if wizard.shield_timer > 0 {
		wizard.shield_timer -= 1
	} else {
		wizard.armor = 0
	}

	if wizard.poison_timer > 0 {
		boss.damaged(3)
		wizard.poison_timer -= 1
	}

	if wizard.recharge_timer > 0 {
		wizard.mana += 101
		wizard.recharge_timer -= 1
	}
}

func main() {
	input, err := ioutil.ReadFile("./day-22-input.txt")
	if err != nil {
		panic(err)
	}

	number_regexp := regexp.MustCompile("\\d+")

	numbers := number_regexp.FindAllStringSubmatch(string(input), -1)
	boss_hp, _ := strconv.Atoi(numbers[0][0])
	boss_damage, _ := strconv.Atoi(numbers[1][0])

	spells := []Spell{
		Spell{name: "missile", cost: 53},
		Spell{name: "drain", cost: 73},
		Spell{name: "shield", cost: 113},
		Spell{name: "poison", cost: 173},
		Spell{name: "recharge", cost: 229},
	}

	min_mana := 9000

	// A translation of my Ruby solution
	// This could *definitely* be improved

	for i := 0; i < 10000; i++ {
		boss := &Boss{hp: boss_hp, damage: boss_damage}
		wizard := &Wizard{hp: 50, mana: 500}

		for !wizard.dead() {
			wizard.tick(boss)

			spell := &spells[rand.Intn(len(spells))]
			for !wizard.can_cast(spell) {
				spell = &spells[rand.Intn(len(spells))]
			}

			wizard.cast(spell, boss)
			wizard.tick(boss)

			if boss.dead() {
				if wizard.spent_mana <= min_mana {
					min_mana = wizard.spent_mana
					println(min_mana)
				}
				break
			}

			wizard.damaged(boss.damage)
		}
	}
}
