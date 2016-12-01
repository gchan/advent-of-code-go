package main

import (
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
)

type Person struct {
	hp, damage, armor int
}

type Item struct {
	cost, damage, armor int
}

func win(hero, boss Person) bool {
	heroNetDamage := hero.damage - boss.armor
	bossNetDamage := boss.damage - hero.armor

	if heroNetDamage < 1 {
		heroNetDamage = 1
	}
	if bossNetDamage < 1 {
		bossNetDamage = 1
	}

	turnsToDefeat := math.Ceil(float64(boss.hp) / float64(heroNetDamage))
	turnsToLose := math.Ceil(float64(hero.hp) / float64(bossNetDamage))

	return turnsToDefeat <= turnsToLose
}

func main() {
	input, err := ioutil.ReadFile("./day-21-input.txt")
	if err != nil {
		panic(err)
	}

	numberRegexp := regexp.MustCompile("\\d+")

	numbers := numberRegexp.FindAllStringSubmatch(string(input), -1)
	bossHp, _ := strconv.Atoi(numbers[0][0])
	bossDamage, _ := strconv.Atoi(numbers[1][0])
	bossArmor, _ := strconv.Atoi(numbers[2][0])

	boss := Person{hp: bossHp, damage: bossDamage, armor: bossArmor}

	weapons := []Item{
		Item{cost: 8, damage: 4},
		Item{cost: 10, damage: 5},
		Item{cost: 25, damage: 6},
		Item{cost: 40, damage: 7},
		Item{cost: 74, damage: 8},
	}

	shields := []Item{
		Item{cost: 0, armor: 0},
		Item{cost: 13, armor: 1},
		Item{cost: 31, armor: 2},
		Item{cost: 53, armor: 3},
		Item{cost: 75, armor: 4},
		Item{cost: 102, armor: 5},
	}

	rings := []Item{
		Item{cost: 0, armor: 0},
		Item{cost: 0, armor: 0},
		Item{cost: 25, damage: 1},
		Item{cost: 50, damage: 2},
		Item{cost: 100, damage: 3},
		Item{cost: 20, armor: 1},
		Item{cost: 40, armor: 2},
		Item{cost: 80, armor: 3},
	}

	_ = weapons
	_ = shields
	_ = rings

	maximumCost := 0

	for _, weapon := range weapons {
		for _, shield := range shields {
			for _, ringOne := range rings {
				for _, ringTwo := range rings {
					if ringOne == ringTwo {
						continue
					}

					cost := weapon.cost + shield.cost + ringOne.cost + ringTwo.cost
					damage := weapon.damage + shield.damage + ringOne.damage + ringTwo.damage
					armor := weapon.armor + shield.armor + ringOne.armor + ringTwo.armor

					hero := Person{hp: 100, damage: damage, armor: armor}

					if !win(hero, boss) && cost > maximumCost {
						maximumCost = cost
					}
				}
			}
		}
	}

	println(maximumCost)
}
