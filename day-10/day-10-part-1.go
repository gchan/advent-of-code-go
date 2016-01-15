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

	for i := 0; i < 40; i++ {
		new_numbers := make([]byte, 0)

		for j := 0; j < len(numbers); j++ {
			num := numbers[j]

			start := j
			for j < len(numbers)-1 && numbers[j+1] == num {
				j++
			}
			end := j

			length := end - start + 1

			new_numbers = append(new_numbers, strconv.Itoa(length)[0])
			new_numbers = append(new_numbers, numbers[j])
		}

		numbers = new_numbers
	}

	println(len(numbers))
}
