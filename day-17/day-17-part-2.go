package main

import (
  "io/ioutil"
  "strings"
  "strconv"
)

func main() {
  input, err := ioutil.ReadFile("./day-17-input.txt")
  if err != nil { panic(err) }

  descriptions := strings.Split(string(input), "\n")
  containers   := []int{}
  liters       := 150
  combinations := map[int]int{}

  for _, description := range descriptions {
    size, _ := strconv.Atoi(description)
    containers = append(containers, size)
  }

  test_combo := func(combo []int) {
    sum := 0
    for _, container := range combo {
      sum += container
    }
    if sum == liters {
      combinations[len(combo)]++
    }
  }

  for num := 1; num <= len(containers); num++ {
    combo := make([]int, num)

    var next_combo func(int, int)
    last := len(combo) - 1
    next_combo = func(i, from int) {
      for j := from; j < len(containers); j++ {
        combo[i] = containers[j]
        if i == last {
          test_combo(combo)
        } else {
          next_combo(i + 1, j + 1)
        }
      }
    }

    next_combo(0, 0)

    if len(combinations) > 0 {
      println(combinations[num])
      break
    }
  }
}
