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
	var x, y int
	visited := make(map[string]bool)

	string := strings.Replace(string(input), ",", "", -1)

	for _, command := range strings.Split(string, " ") {
		if command[0] == 'L' {
			dir += -1 + 4
		} else {
			dir++
		}

		dir %= 4

		steps, _ := strconv.Atoi(command[1:])

		for i := 0; i < steps; i++ {
			switch dir {
			case 0:
				y++
			case 1:
				x++
			case 2:
				y--
			case 3:
				x--
			}

			location := strconv.Itoa(x) + "," + strconv.Itoa(y)

			if visited[location] {
				println(int(math.Abs(float64(x)) + math.Abs(float64(y))))
				return
			}

			visited[location] = true
		}
	}
}
