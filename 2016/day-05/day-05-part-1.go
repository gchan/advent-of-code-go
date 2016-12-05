package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	input, _ := ioutil.ReadFile("./day-05-input.txt")

	id := string(input)
	var pw string
	var i int

	zeroesRegexp := regexp.MustCompile("\\A0{5}")

	for len(pw) < 8 {
		digest := md5.Sum([]byte(id + strconv.Itoa(i)))

		digestStr := fmt.Sprintf("%x", digest)

		if zeroesRegexp.MatchString(digestStr) {
			pw += string(digestStr[5])
			println(pw)
		}

		i++
	}
}
