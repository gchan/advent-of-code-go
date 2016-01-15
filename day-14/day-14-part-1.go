package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Reindeer struct {
	speed, duration, rest, distance int
}

func (r *Reindeer) race(seconds int) {
	for seconds > 0 {
		if seconds < r.duration {
			r.distance += r.speed * seconds
		} else {
			r.distance += r.speed * r.duration
		}
		seconds -= r.duration
		seconds -= r.rest
	}
}

func main() {
	input, err := ioutil.ReadFile("./day-14-input.txt")
	if err != nil {
		panic(err)
	}

	descriptions := strings.Split(string(input), "\n")

	regex := regexp.MustCompile("(\\d+) km/s.* (\\d+) seconds.* (\\d+)")

	seconds := 2503
	maximumDistance := -1

	for _, description := range descriptions {
		matches := regex.FindStringSubmatch(description)[1:]

		speed, _ := strconv.Atoi(matches[0])
		duration, _ := strconv.Atoi(matches[1])
		rest, _ := strconv.Atoi(matches[2])

		reindeer := Reindeer{speed: speed, duration: duration, rest: rest}
		reindeer.race(seconds)

		if reindeer.distance > maximumDistance {
			maximumDistance = reindeer.distance
		}
	}

	println(maximumDistance)
}
