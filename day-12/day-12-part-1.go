package main

import (
  "io/ioutil"
  "encoding/json"
)

func find_numbers(input interface{}) []int {
  numbers := []int{}

  switch input := input.(type) {
  case []interface{}:
    for _, value := range input {
      numbers = append(numbers, find_numbers(value)...)
    }
  case map[string]interface{}:
    for _, value := range input {
      numbers = append(numbers, find_numbers(value)...)
    }
  case float64:
    numbers = append(numbers, int(input))
  }

  return numbers
}

func main() {
  input, err := ioutil.ReadFile("./day-12-input.txt")
  if err != nil { panic(err) }

  data := make(map[string]interface{}, 0)
  json.Unmarshal(input, &data)

  sum := 0
  for _, num := range find_numbers(data) {
    sum += num
  }

  println(sum)
}
