package main

import (
  "io/ioutil"
  "strings"
  "strconv"
)

func main() {
  input, err := ioutil.ReadFile("./day-08-input.txt")
  if err != nil { panic(err) }

  lines := strings.Split(string(input), "\n")

  code_characters   := 0
  memory_characters := 0

  for _, line := range lines {
    code_characters += len(line)
    unescaped, _ := strconv.Unquote(line)
    memory_characters += len(unescaped)
  }

  println(code_characters - memory_characters)
}
