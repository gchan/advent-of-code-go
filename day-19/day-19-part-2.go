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

	elementRegex := regexp.MustCompile("\\A\\w*")
	newElementRegex := regexp.MustCompile("\\w*\\z")

	for _, replacement := range lines {
		element := elementRegex.FindString(replacement)
		newElement := newElementRegex.FindString(replacement)

		replacements[newElement] = element
	}

	moleculeDup := make([]byte, len(molecule))
	electronRegex := regexp.MustCompile("e+")
	steps := 0

	for !electronRegex.Match(moleculeDup) {
		steps = 0
		exhausted := false
		moleculeDup = make([]byte, len(molecule))
		copy(moleculeDup, molecule)

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
				newElement := replacements[element]
				if strings.Contains(string(moleculeDup), element) {
					moleculeDup = []byte(strings.Replace(string(moleculeDup), element, newElement, 1))
					steps++
					exhausted = false
					break
				}
			}
		}
	}

	println(steps)
}
