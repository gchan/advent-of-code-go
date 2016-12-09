package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./day-09-input.txt")
	str := string(input)

	var idx, length int
	markerRegexp := regexp.MustCompile("\\(([0-9x]*)\\)")

	for idx < len(str) {
		if str[idx] == '(' {
			marker := markerRegexp.FindStringSubmatch(str[idx:])[1]
			nums := strings.Split(marker, "x")

			chars, _ := strconv.Atoi(nums[0])
			repeat, _ := strconv.Atoi(nums[1])

			length += chars * repeat
			idx += chars + len(marker) + 2
		} else {
			length++
			idx++
		}
	}

	println(length)
}
