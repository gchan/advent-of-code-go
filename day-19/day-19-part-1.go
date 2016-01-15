package main

import (
	"io/ioutil"
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

	replacements := make(map[string][]string)
	newMolecules := make(map[string]bool)

	elementRegex := regexp.MustCompile("\\A\\w*")
	newElementRegex := regexp.MustCompile("\\w*\\z")

	for _, replacement := range lines {
		element := elementRegex.FindString(replacement)
		newElement := newElementRegex.FindString(replacement)

		if replacements[element] == nil {
			replacements[element] = []string{}
		}

		replacements[element] = append(replacements[element], newElement)
	}

	elementRegex = regexp.MustCompile("[A-Z][a-z]*")
	for _, index := range elementRegex.FindAllStringSubmatchIndex(molecule, -1) {
		element := molecule[index[0]:index[1]]

		for _, replacement := range replacements[element] {
			newMolecule := make([]byte, len(molecule))
			copy(newMolecule, molecule)
			newMolecule = append(newMolecule[0:index[0]], append([]byte(replacement), newMolecule[index[1]:]...)...)
			newMolecules[string(newMolecule)] = true
		}
	}

	println(len(newMolecules))
}
