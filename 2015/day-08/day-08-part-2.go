package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day-08-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	codeCharacters := 0
	encodedCharacters := 0

	for scanner.Scan() {
		line := scanner.Text()

		codeCharacters += len(line)
		escaped := strconv.Quote(line)
		encodedCharacters += len(escaped)
	}

	println(encodedCharacters - codeCharacters)
}
