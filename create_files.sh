#!/bin/bash

for y in `seq -w 2015 2016`; do
  mkdir $y

  for i in `seq -w 1 25`; do
    mkdir $y/day-$i
    touch $y/day-$i/README.md
    touch $y/day-$i/day-$i-part-1.go
    touch $y/day-$i/day-$i-part-2.go
    touch $y/day-$i/day-$i-input.txt
  done
done