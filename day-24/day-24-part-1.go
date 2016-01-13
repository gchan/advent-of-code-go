package main

import (
  "io/ioutil"
  "strings"
  "strconv"
  "math"
)

func generate_combinations(elements []int, length int) <-chan []int {
  c := make(chan []int)
  combo := make([]int, length)
  var next_combo func(int, int)
  last := len(combo) - 1

  next_combo = func(i, from int) {
    for j := from; j < len(elements); j++ {
      combo[i] = elements[j]
      if i == last {
        c <- combo
      } else {
        next_combo(i + 1, j + 1)
      }
    }
  }

  go func(c chan []int) {
    defer close(c)
    next_combo(0, 0)
  }(c)

  return c
}

func main() {
  input, err := ioutil.ReadFile("./day-24-input.txt")
  if err != nil { panic(err) }

  presents_strs := strings.Split(string(input), "\n")
  presents := make([]int, len(presents_strs))

  for i, present := range presents_strs {
    presents[i], _ = strconv.Atoi(present)
  }

  groups := 3

  sum := 0
  min_qe := math.MaxInt64
  for _, present := range presents {
    sum += present
  }

  target := sum / groups
  max_group_size := len(presents) / groups

  for group_size := 1; group_size < max_group_size + 1; group_size++ {
    min_eq_found := false
    combos := generate_combinations(presents, group_size)

    for combo := range combos {
      sum := 0
      qe := 1

      for _, present := range combo {
        sum += present
        qe  *= present
      }

      if sum == target && qe < min_qe {
        min_qe = qe
        min_eq_found = true
      }
    }

    if min_eq_found { break }
  }

  println(min_qe)
}
