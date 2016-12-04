package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day-04-input.txt")

	nameRegexp := regexp.MustCompile("[a-z-]*")
	idRegexp := regexp.MustCompile("\\d{3}")
	targetRegexp := regexp.MustCompile("north")

	for _, room := range strings.Fields(string(input)) {
		name := nameRegexp.FindString(room)
		id, _ := strconv.Atoi(idRegexp.FindString(room))

		cipher := map[rune]rune{
			'-': '-',
		}

		var decryptedName string

		for i := 0; i < 26; i++ {
			cipher[rune(int('a')+i)] = rune((int('a') + (i+id)%26))
		}

		for _, char := range name {
			decryptedName += string(cipher[char])
		}

		if targetRegexp.MatchString(decryptedName) {
			println(fmt.Sprintf("%s%d", decryptedName, id))
		}
	}
}
