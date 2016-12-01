package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day-01-input.txt")

	var dir int
	var x, y float64

	string := strings.Replace(string(input), ",", "", -1)

	for _, command := range strings.Split(string, " ") {
		if command[0] == 'L' {
			dir += -1 + 4
		} else {
			dir++
		}

		dir %= 4

		blocks, _ := strconv.Atoi(command[1:])
		steps := float64(blocks)

		switch dir {
		case 0:
			y += steps
		case 1:
			x += steps
		case 2:
			y -= steps
		case 3:
			x -= steps
		}
	}

	println(int(math.Abs(x) + math.Abs(y)))
}
