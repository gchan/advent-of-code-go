package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Circuit struct {
	outputWire string
	inputWires []string
	values     []int
	gate       string
}

func NewCircuit(instruction string) *Circuit {
	gateRegexp := regexp.MustCompile("[A-Z]+")
	outputRegexp := regexp.MustCompile("-> (.+)")
	inputRegexp := regexp.MustCompile("([a-z0-9]+) ")
	wireRegexp := regexp.MustCompile("[a-z]+")

	gate := gateRegexp.FindString(instruction)
	outputWire := outputRegexp.FindStringSubmatch(instruction)[1]
	inputs := inputRegexp.FindAllStringSubmatch(instruction, -1)

	var inputWires []string
	var values []int

	for _, input := range inputs {
		input := input[1]
		if wireRegexp.MatchString(input) {
			inputWires = append(inputWires, input)
		} else {
			value, _ := strconv.Atoi(input)
			values = append(values, value)
		}
	}

	return &Circuit{outputWire, inputWires, values, gate}
}

func (circuit *Circuit) simplify(wire string, value int) {
	for i, inputWire := range circuit.inputWires {
		if inputWire == wire {
			circuit.inputWires = append(circuit.inputWires[:i], circuit.inputWires[i+1:]...)
			circuit.values = append([]int{value}, circuit.values...)
			return
		}
	}
}

func (circuit *Circuit) output() int {
	values := circuit.values

	switch circuit.gate {
	case "AND":
		return values[0] & values[1]
	case "OR":
		return values[0] | values[1]
	case "NOT":
		return ^values[0]
	case "LSHIFT":
		return values[0] << uint(values[1])
	case "RSHIFT":
		return values[0] >> uint(values[1])
	default:
		return values[0]
	}
}

func (circuit *Circuit) hasOutput() bool {
	return len(circuit.inputWires) == 0
}

func main() {
	input, err := ioutil.ReadFile("./day-07-input.txt")
	if err != nil {
		panic(err)
	}

	instructions := strings.Split(string(input), "\n")

	// Circuits organised by input wires
	circuits := make(map[string][]*Circuit)
	outputs := make(map[string]int)
	targetWire := "a"

	for _, instruction := range instructions {
		circuit := NewCircuit(instruction)

		for _, input := range circuit.inputWires {
			circuits[input] = append(circuits[input], circuit)
		}

		if circuit.hasOutput() {
			outputs[circuit.outputWire] = circuit.output()
		}
	}

	outputs["b"] = 16076

	for {
		evaluated := make(map[string]int)

		for wire, wireValue := range outputs {
			for _, circuit := range circuits[wire] {
				circuit.simplify(wire, wireValue)

				if circuit.hasOutput() {
					evaluated[circuit.outputWire] = circuit.output()
				}
			}

			delete(outputs, wire)
		}

		for wire, wireValue := range evaluated {
			outputs[wire] = wireValue
		}

		finalOutput, present := outputs[targetWire]
		if present {
			println(finalOutput)
			return
		}
	}
}
