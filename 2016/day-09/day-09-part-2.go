package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func decompress(str string) (length int) {
	var idx int
	markerRegexp := regexp.MustCompile("\\(([0-9x]*)\\)")

	for idx < len(str) {
		if str[idx] == '(' {
			marker := markerRegexp.FindStringSubmatch(str[idx:])[1]
			nums := strings.Split(marker, "x")

			chars, _ := strconv.Atoi(nums[0])
			repeat, _ := strconv.Atoi(nums[1])

			subStart := idx + len(marker) + 2
			substring := str[subStart : subStart+chars]

			length += decompress(substring) * repeat
			idx += chars + len(marker) + 2
		} else {
			length++
			idx++
		}
	}

	return length
}

func main() {
	input, _ := ioutil.ReadFile("./day-09-input.txt")

	println(decompress(string(input)))
}
