package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func abba(text string) bool {
	for i := 0; i < len(text)-3; i++ {
		str := text[i:(i + 4)]

		if str[0] == str[1] {
			continue
		}

		if str[0] == str[3] && str[1] == str[2] {
			return true
		}
	}

	return false
}

func main() {
	input, _ := ioutil.ReadFile("./day-07-input.txt")

	lines := strings.Split(string(input), "\n")
	sequenceRegexFirst := regexp.MustCompile("(\\w+)\\[")
	sequenceRegexRest := regexp.MustCompile("\\](\\w+)")
	hypernetRegex := regexp.MustCompile("\\[(\\w+)\\]")
	var count int

	for _, line := range lines {
		sequences := sequenceRegexFirst.FindAllStringSubmatch(line, -1)
		sequences = append(sequences, sequenceRegexRest.FindAllStringSubmatch(line, -1)...)

		hypernets := hypernetRegex.FindAllStringSubmatch(line, -1)

		sequenceAbba := false
		for _, sequence := range sequences {
			if abba(sequence[1]) {
				sequenceAbba = true
				break
			}
		}

		hypernetAbba := false
		for _, hypernet := range hypernets {
			if abba(hypernet[1]) {
				hypernetAbba = true
				break
			}
		}

		if sequenceAbba && !hypernetAbba {
			count++
		}
	}

	println(count)
}
