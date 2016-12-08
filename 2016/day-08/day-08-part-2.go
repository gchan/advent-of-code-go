// Same as day-08-part-1.go

package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day-08-input.txt")
	lines := strings.Split(string(input), "\n")

	pixels := [6][50]int{}

	for _, line := range lines {
		if line[0:4] == "rect" {
			matches := regexp.MustCompile("(\\d*)x(\\d*)").FindStringSubmatch(line)
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])

			for r := 0; r < x; r++ {
				for c := 0; c < y; c++ {
					pixels[c][r] = 1
				}
			}
		} else { // rotate
			by, _ := strconv.Atoi(regexp.MustCompile("by (\\d*)").FindStringSubmatch(line)[1])
			rowColRegex := regexp.MustCompile("=(\\d*)")

			if regexp.MustCompile("row").MatchString(line) {
				row, _ := strconv.Atoi(rowColRegex.FindStringSubmatch(line)[1])

				copy(pixels[row][:], append(pixels[row][50-by:], pixels[row][:50-by]...))
			} else { // column
				col, _ := strconv.Atoi(rowColRegex.FindStringSubmatch(line)[1])

				colValues := [6]int{}
				for r := 0; r < 6; r++ {
					colValues[r] = pixels[r][col]
				}

				copy(colValues[:], append(colValues[6-by:], colValues[:6-by]...))
				for r := 0; r < 6; r++ {
					pixels[r][col] = colValues[r]
				}
			}
		}
	}

	var count int

	// upojflbcez
	for _, row := range pixels {
		for _, char := range row {
			if char == 1 {
				print("#")
				count++
			} else {
				print(" ")
			}
		}
		println()
	}

	println(count)
}
