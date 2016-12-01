// 8 guests so we can brute force
// !9 = 362880 possible seating arrangements
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
func generatePermutations(n int, strs []string, perms *[][]string) {
	if n == 1 {
		strsCopy := make([]string, len(strs))
		copy(strsCopy, strs)

		*perms = append(*perms, strsCopy)
	} else {
		for i := 0; i < n-1; i++ {
			generatePermutations(n-1, strs, perms)
			if n%2 == 0 {
				swap(strs, i, n-1)
			} else {
				swap(strs, 0, n-1)
			}
		}
		generatePermutations(n-1, strs, perms)
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
	happinessRules := make(map[string]map[string]int)

	regex := regexp.MustCompile("\\A(\\w+) .* (gain|lose) (\\d+) .* to (\\w+)")

	for _, rule := range rules {
		matches := regex.FindStringSubmatch(rule)[1:]

		person := matches[0]
		negative := strings.Contains(matches[1], "lose")
		happiness, _ := strconv.Atoi(matches[2])
		nextTo := matches[3]

		if negative {
			happiness *= -1
		}

		if _, present := happinessRules[person]; !present {
			happinessRules[person] = make(map[string]int)
		}

		happinessRules[person][nextTo] = happiness
	}

	guests := []string{"Me"}
	for guest := range happinessRules {
		guests = append(guests, guest)
	}

	maxHappiness := math.MinInt32
	perms := [][]string{}
	generatePermutations(len(guests), guests, &perms)

	for _, seating := range perms {
		seating = append(seating, seating[0])

		sum := 0
		for i := 0; i < len(seating)-1; i++ {
			person := seating[i]
			nextTo := seating[i+1]
			sum += happinessRules[person][nextTo]
			sum += happinessRules[nextTo][person]
		}

		if sum > maxHappiness {
			maxHappiness = sum
		}
	}

	println(maxHappiness)
}
