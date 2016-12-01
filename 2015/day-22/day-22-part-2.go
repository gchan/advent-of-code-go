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
	hp, mana, armor                         int
	shieldTimer, poisonTimer, rechargeTimer int
	spentMana                               int
}

func (wizard *Wizard) damaged(damage int) {
	if damage-wizard.armor > 0 {
		wizard.hp -= damage - wizard.armor
	} else {
		wizard.hp--
	}
}

func (wizard *Wizard) dead() bool {
	return wizard.hp <= 0 || wizard.mana < 53
}

func (wizard *Wizard) cast(spell *Spell, boss *Boss) {
	wizard.spentMana += spell.cost
	wizard.mana -= spell.cost

	switch spell.name {
	case "missile":
		boss.damaged(4)
	case "drain":
		boss.damaged(2)
		wizard.hp += 2
	case "shield":
		wizard.armor = 7
		wizard.shieldTimer = 6
	case "poison":
		wizard.poisonTimer = 6
	case "recharge":
		wizard.rechargeTimer = 5
	}
}

func (wizard *Wizard) canCast(spell *Spell) bool {
	if spell.cost > wizard.mana {
		return false
	}

	switch spell.name {
	case "shield":
		return wizard.shieldTimer == 0
	case "poison":
		return wizard.poisonTimer == 0
	case "recharge":
		return wizard.rechargeTimer == 0
	default:
		return true
	}
}

func (wizard *Wizard) tick(boss *Boss) {
	if wizard.shieldTimer > 0 {
		wizard.shieldTimer--
	} else {
		wizard.armor = 0
	}

	if wizard.poisonTimer > 0 {
		boss.damaged(3)
		wizard.poisonTimer--
	}

	if wizard.rechargeTimer > 0 {
		wizard.mana += 101
		wizard.rechargeTimer--
	}
}

func main() {
	input, err := ioutil.ReadFile("./day-22-input.txt")
	if err != nil {
		panic(err)
	}

	numberRegexp := regexp.MustCompile("\\d+")

	numbers := numberRegexp.FindAllStringSubmatch(string(input), -1)
	bossHp, _ := strconv.Atoi(numbers[0][0])
	bossDamage, _ := strconv.Atoi(numbers[1][0])

	spells := []Spell{
		Spell{name: "missile", cost: 53},
		Spell{name: "drain", cost: 73},
		Spell{name: "shield", cost: 113},
		Spell{name: "poison", cost: 173},
		Spell{name: "recharge", cost: 229},
	}

	minMana := 9000

	// A translation of my Ruby solution
	// This could *definitely* be improved

	for i := 0; i < 200000; i++ {
		boss := &Boss{hp: bossHp, damage: bossDamage}
		wizard := &Wizard{hp: 50, mana: 500}

		wizard.damaged(1)

		for !wizard.dead() {
			wizard.tick(boss)

			spell := &spells[rand.Intn(len(spells))]
			for !wizard.canCast(spell) {
				spell = &spells[rand.Intn(len(spells))]
			}

			wizard.cast(spell, boss)
			wizard.tick(boss)

			if boss.dead() {
				if wizard.spentMana <= minMana {
					minMana = wizard.spentMana
					println(minMana)
				}
				break
			}

			wizard.damaged(boss.damage)
			wizard.damaged(1)
		}
	}
}
