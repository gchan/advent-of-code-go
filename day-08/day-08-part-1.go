package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./day-08-input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	codeCharacters := 0
	memoryCharacters := 0

	for _, line := range lines {
		codeCharacters += len(line)
		unescaped, _ := strconv.Unquote(line)
		memoryCharacters += len(unescaped)
	}

	println(codeCharacters - memoryCharacters)
}
