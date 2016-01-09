package main

import (
  "io/ioutil"
  "strconv"
)

func main() {
  input, err := ioutil.ReadFile("./day-20-input.txt")
  if err != nil { panic(err) }

  target, _ := strconv.Atoi(string(input))
  houses := make([]int, target)

  for elf := 1; elf <= target / 10; elf++ {
    upto := target / 10 / elf
    if upto > 50 {
      upto = 50
    }
    for number := 1; number <= upto; number++ {
      houses[number * elf] += elf * 11
    }
  }

  for i := 1; i < len(houses); i++ {
    presents := houses[i]
    if presents >= target {
      println(i)
      break
    }
  }
}
