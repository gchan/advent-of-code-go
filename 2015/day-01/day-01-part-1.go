package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./day-01-input.txt")
	if err != nil {
		panic(err)
	}

	string := string(input)

	up := strings.Count(string, "(")
	down := strings.Count(string, ")")

	floor := up - down

	println(floor)
}
