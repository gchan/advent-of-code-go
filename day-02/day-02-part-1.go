package main

import (
  "io/ioutil"
  "strconv"
  "strings"
  "sort"
)

func main() {
  input, err := ioutil.ReadFile("./day-02-input.txt")
  if err != nil { panic(err) }

  presents := strings.Split(string(input), "\n")

  total := 0

  for _, present := range presents {
    sides := []int{}

    for _, side_string := range strings.Split(present, "x") {
      side, _ := strconv.Atoi(side_string)
      sides = append(sides, side)
    }

    sort.Ints(sides)

    total += sides[0] * sides[1] * 3
    total += sides[0] * sides[2] * 2
    total += sides[1] * sides[2] * 2
  }

  println(total)
}
