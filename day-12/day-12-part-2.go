package main

import (
	"encoding/json"
	"io/ioutil"
)

func find_numbers(input interface{}) []int {
	numbers := []int{}

	switch input := input.(type) {
	case []interface{}:
		for _, value := range input {
			numbers = append(numbers, find_numbers(value)...)
		}
	case map[string]interface{}:
		no_red := true

		for _, value := range input {
			if str, ok := value.(string); ok && str == "red" {
				no_red = false
				break
			}
		}

		if no_red {
			for _, value := range input {
				numbers = append(numbers, find_numbers(value)...)
			}
		}
	case float64:
		numbers = append(numbers, int(input))
	}

	return numbers
}

func main() {
	input, err := ioutil.ReadFile("./day-12-input.txt")
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{}, 0)
	json.Unmarshal(input, &data)

	sum := 0
	for _, num := range find_numbers(data) {
		sum += num
	}

	println(sum)
}
