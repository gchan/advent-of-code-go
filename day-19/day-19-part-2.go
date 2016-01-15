package main

import (
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./day-19-input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	molecule := lines[len(lines)-1]
	lines = lines[:len(lines)-2]

	replacements := make(map[string]string)

	element_regex := regexp.MustCompile("\\A\\w*")
	new_element_regex := regexp.MustCompile("\\w*\\z")

	for _, replacement := range lines {
		element := element_regex.FindString(replacement)
		new_element := new_element_regex.FindString(replacement)

		replacements[new_element] = element
	}

	molecule_dup := make([]byte, len(molecule))
	electron_regex := regexp.MustCompile("e+")
	steps := 0

	for !electron_regex.Match(molecule_dup) {
		steps = 0
		exhausted := false
		molecule_dup = make([]byte, len(molecule))
		copy(molecule_dup, molecule)

		for !exhausted {
			exhausted = true

			keys := make([]string, len(replacements))

			i := 0
			for k := range replacements {
				keys[i] = k
				i++
			}

			list := rand.Perm(len(replacements))
			for _, index := range list {
				element := keys[index]
				new_element := replacements[element]
				if strings.Contains(string(molecule_dup), element) {
					molecule_dup = []byte(strings.Replace(string(molecule_dup), element, new_element, 1))
					steps++
					exhausted = false
					break
				}
			}
		}
	}

	println(steps)
}
