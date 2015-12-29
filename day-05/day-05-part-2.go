package main

import (
  "os"
  "bufio"
  "strings"
)

func repeating_letter_with_gap(str string) bool {
  for i := 0; i < len(str) - 2; i++ {
    if str[i] == str[i + 2] {
      return true
    }
  }

  return false
}

func two_pairs(str string) bool {
  for i := 0; i < len(str) - 2; i++ {
    if strings.Count(str, str[i:i+2]) >= 2 {
      return true
    }
  }

  return false
}

func main() {
  file, err := os.Open("./day-05-input.txt")
  if err != nil { panic(err) }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  count   := 0

  for scanner.Scan() {
    str := scanner.Text()
    if repeating_letter_with_gap(str) && two_pairs(str) {
      count++
    }
  }

  println(count)
}
