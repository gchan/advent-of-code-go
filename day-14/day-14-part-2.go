package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Reindeer struct {
	speed, duration, rest, distance int
	points, restingTime, flyingTime int
}

func (r *Reindeer) tick() {
	if r.restingTime > 0 {
		r.restingTime--
	} else {
		r.flyingTime++
		r.distance += r.speed
	}

	if r.flyingTime == r.duration {
		r.flyingTime = 0
		r.restingTime = r.rest
	}
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
	reindeers := []*Reindeer{}

	regex := regexp.MustCompile("(\\d+) km/s.* (\\d+) seconds.* (\\d+)")

	seconds := 2503

	for _, description := range descriptions {
		matches := regex.FindStringSubmatch(description)[1:]

		speed, _ := strconv.Atoi(matches[0])
		duration, _ := strconv.Atoi(matches[1])
		rest, _ := strconv.Atoi(matches[2])

		reindeer := Reindeer{speed: speed, duration: duration, rest: rest}
		reindeers = append(reindeers, &reindeer)
	}

	for s := 0; s < seconds; s++ {
		maximumDistance := 0
		for _, reindeer := range reindeers {
			reindeer.tick()
			if reindeer.distance > maximumDistance {
				maximumDistance = reindeer.distance
			}
		}

		for _, reindeer := range reindeers {
			if reindeer.distance == maximumDistance {
				reindeer.points++
			}
		}
	}

	maximumPoints := -1
	for _, reindeer := range reindeers {
		if reindeer.points > maximumPoints {
			maximumPoints = reindeer.points
		}
	}

	println(maximumPoints)
}
