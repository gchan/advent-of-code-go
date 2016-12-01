package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Instruction struct {
	name, register string
	offset         int
}

func main() {
	input, err := ioutil.ReadFile("./day-23-input.txt")
	if err != nil {
		panic(err)
	}

	instructions := []Instruction{}

	for _, instruction := range strings.Split(string(input), "\n") {
		instruction = strings.Replace(instruction, ",", "", -1)
		components := strings.Split(instruction, " ")

		name := components[0]
		var register string
		var offset int

		if len(components) == 3 {
			register = components[1]
			offset, _ = strconv.Atoi(components[2])
		} else { // == 2
			if name == "jmp" {
				offset, _ = strconv.Atoi(components[1])
			} else {
				register = components[1]
			}
		}

		instr := Instruction{name: name, register: register, offset: offset}
		instructions = append(instructions, instr)
	}

	registers := make(map[string]int)
	registers["a"] = 1

	nextInstruction := 0
	instruction := instructions[nextInstruction]

	for nextInstruction < len(instructions) {
		instruction = instructions[nextInstruction]

		switch instruction.name {
		case "hlf":
			registers[instruction.register] /= 2
			nextInstruction++
		case "tpl":
			registers[instruction.register] *= 3
			nextInstruction++
		case "inc":
			registers[instruction.register]++
			nextInstruction++
		case "jmp":
			nextInstruction += instruction.offset
		case "jie":
			if registers[instruction.register]%2 == 0 {
				nextInstruction += instruction.offset
			} else {
				nextInstruction++
			}
		case "jio":
			if registers[instruction.register] == 1 {
				nextInstruction += instruction.offset
			} else {
				nextInstruction++
			}
		}
	}

	println(registers["b"])
}
