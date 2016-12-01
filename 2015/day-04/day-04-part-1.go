package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
)

func main() {
	input, err := ioutil.ReadFile("./day-04-input.txt")
	if err != nil {
		panic(err)
	}

	number := 0
	regex := regexp.MustCompile("\\A00000")
	digest := ""

	for !regex.MatchString(digest) {
		number++

		hash := md5.New()
		io.WriteString(hash, string(input))
		io.WriteString(hash, fmt.Sprintf("%d", number))

		digest = fmt.Sprintf("%x", hash.Sum(nil))
	}

	println(number)
}
