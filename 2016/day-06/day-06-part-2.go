package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day-06-input.txt")

	lines := strings.Split(string(input), "\n")

	var charFreq [8]map[rune]int
	var message string

	for idx := range charFreq {
		charFreq[idx] = make(map[rune]int)
	}

	for _, line := range lines {
		for idx, char := range line {
			charFreq[idx][char]++
		}
	}

	for _, charCount := range charFreq {
		bestCount := len(lines)
		var bestChar rune

		for char, count := range charCount {
			if count < bestCount {
				bestCount = count
				bestChar = char
			}
		}

		message += string(bestChar)
	}

	println(message)
}
