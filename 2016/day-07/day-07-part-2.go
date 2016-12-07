package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func aba(text string) (result []string) {
	for i := 0; i < len(text)-2; i++ {
		str := text[i : i+3]

		if str[0] == str[2] {
			result = append(result, str)
		}
	}

	return result
}

func bab(text string) (result []string) {
	for _, str := range aba(text) {
		bab := append([]byte(str[1:2]), str[0], str[1])
		result = append(result, string(bab))
	}

	return result
}

func main() {
	input, _ := ioutil.ReadFile("./day-07-input.txt")

	lines := strings.Split(string(input), "\n")
	sequenceRegexFirst := regexp.MustCompile("(\\w+)\\[")
	sequenceRegexRest := regexp.MustCompile("\\](\\w+)")
	hypernetRegex := regexp.MustCompile("\\[(\\w+)\\]")
	var count int
	var match bool

	for _, line := range lines {
		sequences := sequenceRegexFirst.FindAllStringSubmatch(line, -1)
		sequences = append(sequences, sequenceRegexRest.FindAllStringSubmatch(line, -1)...)

		hypernets := hypernetRegex.FindAllStringSubmatch(line, -1)

		var abas []string
		for _, sequence := range sequences {
			abas = append(abas, aba(sequence[1])...)
		}

		var babs []string
		for _, hypernet := range hypernets {
			babs = append(babs, bab(hypernet[1])...)
		}

		match = false
		for _, aba := range abas {
			for _, bab := range babs {
				if aba == bab {
					match = true
				}
			}
		}

		if match {
			count++
		}
	}

	println(count)
}
