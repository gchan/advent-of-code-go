package main

import (
  "os"
  "bufio"
  "strconv"
)

func main() {
  file, err := os.Open("./day-08-input.txt")
  if err != nil { panic(err) }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  code_characters    := 0
  encoded_characters := 0

  for scanner.Scan() {
    line := scanner.Text()

    code_characters += len(line)
    escaped := strconv.Quote(line)
    encoded_characters += len(escaped)
  }

  println(encoded_characters - code_characters)
}
