package main

import (
  "io/ioutil"
  "strings"
)

type Coord struct {
  x, y int
}

type Santa struct {
  coord Coord
}

func (santa *Santa) deliver(homes map[Coord]bool) {
  homes[santa.coord] = true
}

func (santa *Santa) right() {
  coord := santa.coord
  santa.coord = Coord{coord.x + 1, coord.y}
}

func (santa *Santa) left() {
  coord := santa.coord
  santa.coord = Coord{coord.x - 1, coord.y}
}

func (santa *Santa) up() {
  coord := santa.coord
  santa.coord = Coord{coord.x, coord.y - 1}
}

func (santa *Santa) down() {
  coord := santa.coord
  santa.coord = Coord{coord.x, coord.y + 1}
}

func main() {
  input, err := ioutil.ReadFile("./day-03-input.txt")
  if err != nil { panic(err) }

  santas := []*Santa{new(Santa), new(Santa)}

  homes := make(map[Coord]bool)
  homes[Coord{0, 0}] = true

  for index, direction := range strings.Split(string(input), "") {
    santa := santas[index % 2]

    switch direction {
    case ">":
      santa.right()
    case "<":
      santa.left()
    case "^":
      santa.up()
    case "v":
      santa.down()
    }

    santa.deliver(homes)
  }

  println(len(homes))
}
