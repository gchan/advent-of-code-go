package main

import (
  "crypto/md5"
  "io/ioutil"
  "regexp"
  "fmt"
)

func main() {
  input, err := ioutil.ReadFile("./day-04-input.txt")
  if err != nil { panic(err) }

  number := 0
  regex  := regexp.MustCompile("\\A000000")
  digest := ""

  for !regex.MatchString(digest) {
    number += 1

    data := []byte(input)
    data = append(data, fmt.Sprintf("%d", number)...)

    digest = fmt.Sprintf("%x", md5.Sum(data))
  }

  println(number)
}
