package main

import (
	"io/ioutil"
	"strconv"
)

func main() {
	numbers, err := ioutil.ReadFile("./day-10-input.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 50; i++ {
		var newNumbers []byte

		for j := 0; j < len(numbers); j++ {
			num := numbers[j]

			start := j
			for j < len(numbers)-1 && numbers[j+1] == num {
				j++
			}
			end := j

			length := end - start + 1

			newNumbers = append(newNumbers, strconv.Itoa(length)[0])
			newNumbers = append(newNumbers, numbers[j])
		}

		numbers = newNumbers
	}

	println(len(numbers))
}
