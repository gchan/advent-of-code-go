package main

import (
  "io/ioutil"
  "strings"
  "regexp"
  "strconv"
)

type Ingredient struct {
  capacity, durability, flavor, texture, calories int
}

func main() {
  input, err := ioutil.ReadFile("./day-15-input.txt")
  if err != nil { panic(err) }

  descriptions := strings.Split(string(input), "\n")
  ingredients  := []*Ingredient{}

  regex := regexp.MustCompile("(-*\\d+)")
  for _, description := range descriptions {
    matches := regex.FindAllStringSubmatch(description, -1)

    capacity, _   := strconv.Atoi(matches[0][0])
    durability, _ := strconv.Atoi(matches[1][0])
    flavor, _     := strconv.Atoi(matches[2][0])
    texture, _    := strconv.Atoi(matches[3][0])
    calories, _   := strconv.Atoi(matches[4][0])

    ingredient := &Ingredient{capacity, durability, flavor, texture, calories}
    ingredients = append(ingredients, ingredient)
  }

  best_score := 0

  for a := 0; a <= 100; a++ {
    for b := 0; b <= 100 -a; b++ {
      for c := 0; c <= 100 - a - b; c++ {
        d := 100 - a - b - c

        amounts := []int{a, b, c, d}
        var capacity, durability, flavor, texture, calories int

        for i, ing := range ingredients {
          amount := amounts[i]

          capacity   += ing.capacity * amount
          durability += ing.durability * amount
          flavor     += ing.flavor * amount
          texture    += ing.texture * amount
          calories   += ing.calories * amount
        }

        if calories != 500 {
          continue
        }

        if capacity < 0   { capacity = 0 }
        if durability < 0 { durability = 0 }
        if flavor < 0     { flavor = 0 }
        if texture < 0    { texture = 0 }

        score := capacity * durability * flavor * texture

        if score > best_score {
          best_score = score
        }
      }
    }
  }

  println(best_score)
}
