package main

import (
	"io/ioutil"
	"strings"
)

type Coord struct {
	x, y int
}

func main() {
	input, err := ioutil.ReadFile("./day-03-input.txt")
	if err != nil {
		panic(err)
	}

	homes := make(map[Coord]bool)
	coord := Coord{0, 0}

	homes[coord] = true

	for _, direction := range strings.Split(string(input), "") {
		switch direction {
		case ">":
			coord = Coord{coord.x + 1, coord.y}
		case "<":
			coord = Coord{coord.x - 1, coord.y}
		case "^":
			coord = Coord{coord.x, coord.y - 1}
		case "v":
			coord = Coord{coord.x, coord.y + 1}
		}

		homes[coord] = true
	}

	println(len(homes))
}
