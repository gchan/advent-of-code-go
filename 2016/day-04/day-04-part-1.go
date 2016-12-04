package main

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type CharCount struct {
	char  rune
	count int
}

type ByCount []*CharCount

func (a ByCount) Len() int      { return len(a) }
func (a ByCount) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool {
	if a[i].count == a[j].count {
		return a[i].char < a[j].char
	}
	return a[i].count > a[j].count
}

func main() {
	input, _ := ioutil.ReadFile("./day-04-input.txt")

	rooms := strings.Fields(strings.Replace(string(input), "-", "", -1))

	nameRegex := regexp.MustCompile("[a-z]*")
	idRegex := regexp.MustCompile("\\d{3}")
	checksumRegex := regexp.MustCompile("\\[([a-z]{5})\\]")

	var roomSum int

	for _, room := range rooms {
		name := nameRegex.FindString(room)
		id, _ := strconv.Atoi(idRegex.FindString(room))
		checksum := checksumRegex.FindStringSubmatch(room)[1]

		charCountMap := make(map[rune]*CharCount)
		var charCounts []*CharCount
		var expectedChecksum string

		for _, char := range name {
			_, pr := charCountMap[char]
			if pr {
				charCountMap[char].count++
			} else {
				charCount := CharCount{char, 1}
				charCountMap[char] = &charCount
				charCounts = append(charCounts, &charCount)
			}
		}

		sort.Sort(ByCount(charCounts))

		for i := 0; i < 5; i++ {
			expectedChecksum += string(charCounts[i].char)
		}

		if expectedChecksum == checksum {
			roomSum += id
		}
	}

	println(roomSum)
}
