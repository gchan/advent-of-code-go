package main

import (
  "io/ioutil"
  "strings"
)

func valid_password(password string) bool {
  return increasing_straight(password) && two_pairs(password) &&
    excludes_confusing_characters(password)
}

func increasing_straight(password string) bool {
  for i := 0; i < len(password) - 2; i++ {
    if password[i] + 1 == password[i + 1] &&
     password[i] + 2 == password[i + 2] {
      return true
     }
  }

  return false
}

func two_pairs(password string) bool {
  pairs := 0
  for i := 0; i < len(password) - 1; i++ {
    if password[i] == password[i + 1] {
      pairs++
      i++
    }
  }

  return pairs >= 2
}

func excludes_confusing_characters(password string) bool {
  return !strings.ContainsAny(password, "iol")
}

func increment_password(password *string) {
  new_password := []byte(*password)

  for i := len(new_password) - 1; i >= 0; i-- {
    if new_password[i] != 'z' {
      new_password[i] += 1
      break
    } else {
      new_password[i] = 'a'
    }
  }

  *password = string(new_password)
}

func main() {
  input, err := ioutil.ReadFile("./day-11-input.txt")
  if err != nil { panic(err) }

  password := string(input)

  next_password := 1

  for next_password != 0 {
    increment_password(&password)

    if valid_password(password) {
      next_password--
    }
  }

  println(password)
}
