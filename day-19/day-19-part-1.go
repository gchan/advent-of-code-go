package main

import (
  "io/ioutil"
  "strings"
  "regexp"
)

func main() {
  input, err := ioutil.ReadFile("./day-19-input.txt")
  if err != nil { panic(err) }

  lines    := strings.Split(string(input), "\n")
  molecule := lines[len(lines) - 1]
  lines    = lines[:len(lines) - 2]

  replacements := make(map[string][]string)
  new_molecules := make(map[string]bool)

  element_regex := regexp.MustCompile("\\A\\w*")
  new_element_regex := regexp.MustCompile("\\w*\\z")

  for _, replacement := range lines {
    element     := element_regex.FindString(replacement)
    new_element := new_element_regex.FindString(replacement)

    if replacements[element] == nil {
      replacements[element] = []string{}
    }

    replacements[element] = append(replacements[element], new_element)
  }

  element_regex = regexp.MustCompile("[A-Z][a-z]*")
  for _, index := range element_regex.FindAllStringSubmatchIndex(molecule, -1) {
    element := molecule[index[0]:index[1]]

    for _, replacement := range replacements[element] {
      new_molecule := make([]byte, len(molecule))
      copy(new_molecule, molecule)
      new_molecule = append(new_molecule[0:index[0]], append([]byte(replacement), new_molecule[index[1]:]...)...)
      new_molecules[string(new_molecule)] = true
    }
  }

  println(len(new_molecules))
}
