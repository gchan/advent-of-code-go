package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day-06-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [1000][1000]int

	fromRegex := regexp.MustCompile("\\d+,\\d+")
	toRegex := regexp.MustCompile("\\d+,\\d+$")

	for scanner.Scan() {
		instruction := scanner.Text()

		fromStr := fromRegex.FindString(instruction)
		toStr := toRegex.FindString(instruction)
		from := strings.Split(fromStr, ",")
		to := strings.Split(toStr, ",")

		fromX, _ := strconv.Atoi(from[0])
		fromY, _ := strconv.Atoi(from[1])
		toX, _ := strconv.Atoi(to[0])
		toY, _ := strconv.Atoi(to[1])

		for x := fromX; x <= toX; x++ {
			for y := fromY; y <= toY; y++ {
				if strings.Contains(instruction, "off") {
					grid[x][y]--
					if grid[x][y] < 0 {
						grid[x][y] = 0
					}
				} else if strings.Contains(instruction, "on") {
					grid[x][y]++
				} else {
					grid[x][y] += 2
				}
			}
		}
	}

	total := 0

	for _, row := range grid {
		for _, col := range row {
			total += col
		}
	}

	println(total)
}
