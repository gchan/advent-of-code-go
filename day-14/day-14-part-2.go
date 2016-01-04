package main

import (
  "io/ioutil"
  "strings"
  "regexp"
  "strconv"
)

type Reindeer struct {
  speed, duration, rest, distance int
  points, resting_time, flying_time int
}

func (r *Reindeer) tick() {
  if r.resting_time > 0 {
    r.resting_time--
  } else {
    r.flying_time++
    r.distance += r.speed
  }

  if r.flying_time == r.duration {
    r.flying_time = 0
    r.resting_time = r.rest
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
  if err != nil { panic(err) }

  descriptions := strings.Split(string(input), "\n")
  reindeers := []*Reindeer{}

  regex := regexp.MustCompile("(\\d+) km/s.* (\\d+) seconds.* (\\d+)")

  seconds := 2503

  for _, description := range descriptions {
    matches  := regex.FindStringSubmatch(description)[1:]

    speed, _    := strconv.Atoi(matches[0])
    duration, _ := strconv.Atoi(matches[1])
    rest, _     := strconv.Atoi(matches[2])

    reindeer := Reindeer{speed: speed, duration: duration, rest: rest}
    reindeers = append(reindeers, &reindeer)
  }

  for s := 0; s < seconds; s++ {
    maximum_distance := 0
    for _, reindeer := range reindeers {
      reindeer.tick()
      if reindeer.distance > maximum_distance {
        maximum_distance = reindeer.distance
      }
    }

    for _, reindeer := range reindeers {
      if reindeer.distance == maximum_distance {
        reindeer.points++
      }
    }
  }

  maximum_points := -1
  for _, reindeer := range reindeers {
    if reindeer.points > maximum_points {
      maximum_points = reindeer.points
    }
  }

  println(maximum_points)
}
