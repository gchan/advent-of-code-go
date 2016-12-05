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
	pw := "........"
	var i int

	digestRegexp := regexp.MustCompile("\\A0{5}[0-7]")

	for {
		digest := md5.Sum([]byte(id + strconv.Itoa(i)))

		digestStr := fmt.Sprintf("%x", digest)

		if digestRegexp.MatchString(digestStr) {
			pos, _ := strconv.Atoi(digestStr[5:6])

			if pw[pos] == '.' {
				pw = pw[:pos] + digestStr[6:7] + pw[pos+1:]
				println(pw)
			}

			if !regexp.MustCompile("\\.").MatchString(pw) {
				break
			}
		}

		i++
	}
}
