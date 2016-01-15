package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func matches_reading(compound, sue_value string, compound_value int) bool {
	value, _ := strconv.Atoi(sue_value)

	switch compound {
	case "cats":
		return value > compound_value
	case "trees":
		return value > compound_value
	case "pomeranians":
		return value < compound_value
	case "goldfish":
		return value < compound_value
	default:
		return value == compound_value
	}
}

func main() {
	input, err := ioutil.ReadFile("./day-16-input.txt")
	if err != nil {
		panic(err)
	}

	sues := strings.Split(string(input), "\n")

	compounds := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	regexes := map[string]*regexp.Regexp{
		"number":      regexp.MustCompile("Sue (\\d+)"),
		"children":    regexp.MustCompile("children: (\\d+)"),
		"cats":        regexp.MustCompile("cats: (\\d+)"),
		"samoyeds":    regexp.MustCompile("samoyeds: (\\d+)"),
		"pomeranians": regexp.MustCompile("pomeranians: (\\d+)"),
		"akitas":      regexp.MustCompile("akitas: (\\d+)"),
		"vizslas":     regexp.MustCompile("vizslas: (\\d+)"),
		"goldfish":    regexp.MustCompile("goldfish: (\\d+)"),
		"trees":       regexp.MustCompile("trees: (\\d+)"),
		"cars":        regexp.MustCompile("cars: (\\d+)"),
		"perfumes":    regexp.MustCompile("perfumes: (\\d+)"),
	}

	for _, sue := range sues {
		real_sue := true

		for compound, amount := range compounds {
			match := regexes[compound].FindStringSubmatch(sue)
			if len(match) > 0 && !matches_reading(compound, match[1], amount) {
				real_sue = false
				break
			}
		}

		if real_sue {
			println(regexes["number"].FindStringSubmatch(sue)[1])
			break
		}
	}
}
