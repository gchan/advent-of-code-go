package main

import (
  "os"
  "bufio"
  "regexp"
  "strings"
  "strconv"
)

func main() {
  file, err := os.Open("./day-06-input.txt")
  if err != nil { panic(err) }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  var grid [1000][1000]bool

  from_regex := regexp.MustCompile("\\d+,\\d+")
  to_regex   := regexp.MustCompile("\\d+,\\d+$")

  for scanner.Scan() {
    instruction := scanner.Text()

    from_str := from_regex.FindString(instruction)
    to_str   := to_regex.FindString(instruction)
    from := strings.Split(from_str, ",")
    to   := strings.Split(to_str, ",")

    from_x, _ := strconv.Atoi(from[0])
    from_y, _ := strconv.Atoi(from[1])
    to_x, _   := strconv.Atoi(to[0])
    to_y, _   := strconv.Atoi(to[1])

    for x := from_x; x <= to_x; x++ {
      for y:= from_y; y <= to_y; y++ {
        if strings.Contains(instruction, "off") {
          grid[x][y] = false
        } else if strings.Contains(instruction, "on") {
          grid[x][y] = true
        } else {
          grid[x][y] = !grid[x][y]
        }
      }
    }
  }

  total := 0

  for _, row := range grid {
    for _, col := range row {
      if col {
        total++
      }
    }
  }

  println(total)
}
