package main

import (
  "io/ioutil"
  "strings"
)

func neighbours_on(grid [][]bool, x, y int) int {
  sum := 0

  for i := x - 1; i <= x + 1 && i < 100; i++ {
    for j := y - 1; j <= y + 1 && j < 100; j++ {
      if i >= 0 && j >= 0 && !(x == i && y == j) && grid[i][j] {
        sum++
      }
    }
  }

  return sum
}

func make_grid() [][]bool {
  grid := make([][]bool, 100)
  for i := 0; i < 100; i++ {
    grid[i] = make([]bool, 100)
  }
  return grid
}

func main() {
  input, err := ioutil.ReadFile("./day-18-input.txt")
  if err != nil { panic(err) }

  grid := make_grid()
  rows := strings.Split(string(input), "\n")

  for y, row := range rows {
    for x, state := range strings.Split(row, "") {
      grid[x][y] = state == "#"
    }
  }

  for i := 0; i < 100; i++ {
    new_grid := make_grid()

    for x := 0; x < 100; x++ {
      for y := 0; y < 100; y++ {
        neighbours := neighbours_on(grid, x, y)

        if grid[x][y] {
          new_grid[x][y] = neighbours == 2 || neighbours == 3
        } else {
          new_grid[x][y] = neighbours == 3
        }
      }
    }

    grid = new_grid
  }

  sum := 0

  for x := 0; x < 100; x++ {
    for y := 0; y < 100; y++ {
      if grid[x][y] {
        sum++
      }
    }
  }

  println(sum)
}
