// 8 guests so we can brute force
// !8 = 40320 possible seating arrangements
// Some seating arrangements are essentially the same for the purpose of this problem

package main

import (
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Heap's Algorithm
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func generate_permutations(n int, strs []string, perms *[][]string) {
	if n == 1 {
		strs_copy := make([]string, len(strs))
		copy(strs_copy, strs)

		*perms = append(*perms, strs_copy)
	} else {
		for i := 0; i < n-1; i++ {
			generate_permutations(n-1, strs, perms)
			if n%2 == 0 {
				swap(strs, i, n-1)
			} else {
				swap(strs, 0, n-1)
			}
		}
		generate_permutations(n-1, strs, perms)
	}
}

func swap(strs []string, i, j int) {
	strs[i], strs[j] = strs[j], strs[i]
}

func main() {
	input, err := ioutil.ReadFile("./day-13-input.txt")
	if err != nil {
		panic(err)
	}

	rules := strings.Split(string(input), "\n")
	happiness_rules := make(map[string]map[string]int)

	regex := regexp.MustCompile("\\A(\\w+) .* (gain|lose) (\\d+) .* to (\\w+)")

	for _, rule := range rules {
		matches := regex.FindStringSubmatch(rule)[1:]

		person := matches[0]
		negative := strings.Contains(matches[1], "lose")
		happiness, _ := strconv.Atoi(matches[2])
		next_to := matches[3]

		if negative {
			happiness *= -1
		}

		if _, present := happiness_rules[person]; !present {
			happiness_rules[person] = make(map[string]int)
		}

		happiness_rules[person][next_to] = happiness
	}

	guests := []string{}
	for guest, _ := range happiness_rules {
		guests = append(guests, guest)
	}

	max_happiness := math.MinInt32
	perms := [][]string{}
	generate_permutations(len(guests), guests, &perms)

	for _, seating := range perms {
		seating = append(seating, seating[0])

		sum := 0
		for i := 0; i < len(seating)-1; i++ {
			person := seating[i]
			next_to := seating[i+1]
			sum += happiness_rules[person][next_to]
			sum += happiness_rules[next_to][person]
		}

		if sum > max_happiness {
			max_happiness = sum
		}
	}

	println(max_happiness)
}
