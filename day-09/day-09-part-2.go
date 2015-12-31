// It's only 8 nodes (locations), so we can brute force with no problem :)
// !8 = 40320 possible paths

package main

import (
  "io/ioutil"
  "strings"
  "regexp"
  "strconv"
  "math"
)

// Heap's Algorithm
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func generate_permutations(n int, strs []string, perms *[][]string) {
  if n == 1 {
    strs_copy := make([]string, len(strs))
    copy(strs_copy, strs)

    *perms = append(*perms, strs_copy)
  } else {
    for i := 0; i < n - 1; i++ {
      generate_permutations(n - 1, strs, perms)
      if n % 2 == 0 {
        swap(strs, i, n - 1)
      } else {
        swap(strs, 0, n - 1)
      }
    }
    generate_permutations(n - 1, strs, perms)
  }
}

func swap(strs []string, i, j int) {
  strs[i], strs[j] = strs[j], strs[i]
}

func main() {
  input, err := ioutil.ReadFile("./day-09-input.txt")
  if err != nil { panic(err) }

  inputs := strings.Split(string(input), "\n")
  distances := make(map[string]map[string]int)

  destination_regexp := regexp.MustCompile("[A-Z][A-z]+")
  length_regexp      := regexp.MustCompile("\\d+")

  for _, distance := range inputs {
    desintations := destination_regexp.FindAllStringSubmatch(distance, -1)
    from := desintations[0][0]
    to   := desintations[1][0]

    length, _ := strconv.Atoi(length_regexp.FindString(distance))

    if _, present := distances[from]; !present {
      distances[from] = make(map[string]int)
    }

    if _, present := distances[to]; !present {
      distances[to] = make(map[string]int)
    }

    distances[from][to] = length
    distances[to][from] = length
  }

  longest := math.MinInt32
  locations := make([]string, 0)
  for location, _ := range distances {
    locations = append(locations, location)
  }

  routes := make([][]string, 0)
  generate_permutations(len(locations), locations, &routes)

  for _, route := range routes {
    total := 0

    for i := 0; i < len(route) - 1; i++ {
      from := route[i]
      to   := route[i + 1]
      total += distances[from][to]
    }

    if total > longest {
      longest = total
    }
  }

  println(longest)
}
