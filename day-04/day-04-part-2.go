package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	input, err := ioutil.ReadFile("./day-04-input.txt")
	if err != nil {
		panic(err)
	}

	number := 0
	regex := regexp.MustCompile("\\A000000")
	digest := ""

	for !regex.MatchString(digest) {
		number++

		data := []byte(input)
		data = append(data, fmt.Sprintf("%d", number)...)

		digest = fmt.Sprintf("%x", md5.Sum(data))
	}

	println(number)
}
