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
	next_instruction := 0
	instruction := instructions[next_instruction]

	for next_instruction < len(instructions) {
		instruction = instructions[next_instruction]

		switch instruction.name {
		case "hlf":
			registers[instruction.register] /= 2
			next_instruction += 1
		case "tpl":
			registers[instruction.register] *= 3
			next_instruction += 1
		case "inc":
			registers[instruction.register] += 1
			next_instruction += 1
		case "jmp":
			next_instruction += instruction.offset
		case "jie":
			if registers[instruction.register]%2 == 0 {
				next_instruction += instruction.offset
			} else {
				next_instruction += 1
			}
		case "jio":
			if registers[instruction.register] == 1 {
				next_instruction += instruction.offset
			} else {
				next_instruction += 1
			}
		}
	}

	println(registers["b"])
}
