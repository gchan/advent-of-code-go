package main

import (
  "math"
  "io/ioutil"
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
  hero_net_damage := hero.damage - boss.armor
  boss_net_damage := boss.damage - hero.armor

  if hero_net_damage < 1 { hero_net_damage = 1 }
  if boss_net_damage < 1 { boss_net_damage = 1 }

  turns_to_defeat := math.Ceil(float64(boss.hp) / float64(hero_net_damage))
  turns_to_lose   := math.Ceil(float64(hero.hp) / float64(boss_net_damage))

  return turns_to_defeat <= turns_to_lose
}

func main() {
  input, err := ioutil.ReadFile("./day-21-input.txt")
  if err != nil { panic(err) }

  number_regexp := regexp.MustCompile("\\d+")

  numbers := number_regexp.FindAllStringSubmatch(string(input), -1)
  boss_hp, _     := strconv.Atoi(numbers[0][0])
  boss_damage, _ := strconv.Atoi(numbers[1][0])
  boss_armor, _  := strconv.Atoi(numbers[2][0])

  boss := Person{hp: boss_hp, damage: boss_damage, armor: boss_armor}

  weapons := []Item{
    Item{cost: 8,  damage: 4},
    Item{cost: 10, damage: 5},
    Item{cost: 25, damage: 6},
    Item{cost: 40, damage: 7},
    Item{cost: 74, damage: 8},
  }

  shields := []Item{
    Item{cost: 0,   armor: 0},
    Item{cost: 13,  armor: 1},
    Item{cost: 31,  armor: 2},
    Item{cost: 53,  armor: 3},
    Item{cost: 75,  armor: 4},
    Item{cost: 102, armor: 5},
  }

  rings := []Item{
    Item{cost: 0,   armor: 0},
    Item{cost: 0,   armor: 0},
    Item{cost: 25,  damage: 1},
    Item{cost: 50,  damage: 2},
    Item{cost: 100, damage: 3},
    Item{cost: 20,  armor: 1},
    Item{cost: 40,  armor: 2},
    Item{cost: 80,  armor: 3},
  }

  _ = weapons
  _ = shields
  _ = rings

  minimum_cost := 9000

  for _, weapon := range weapons {
    for _, shield := range shields {
      for _, ring_one := range rings {
        for _, ring_two := range rings {
          if ring_one == ring_two {
            continue
          }

          cost   := weapon.cost + shield.cost + ring_one.cost + ring_two.cost
          damage := weapon.damage + shield.damage + ring_one.damage + ring_two.damage
          armor  := weapon.armor + shield.armor + ring_one.armor + ring_two.armor

          hero := Person{hp: 100, damage: damage, armor: armor}

          if win(hero, boss) && cost < minimum_cost {
            minimum_cost = cost
          }
        }
      }
    }
  }

  println(minimum_cost)
}
