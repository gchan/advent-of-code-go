package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Circuit struct {
	output_wire string
	input_wires []string
	values      []int
	gate        string
}

func NewCircuit(instruction string) *Circuit {
	gate_regexp := regexp.MustCompile("[A-Z]+")
	output_regexp := regexp.MustCompile("-> (.+)")
	input_regexp := regexp.MustCompile("([a-z0-9]+) ")
	wire_regexp := regexp.MustCompile("[a-z]+")

	gate := gate_regexp.FindString(instruction)
	output_wire := output_regexp.FindStringSubmatch(instruction)[1]
	inputs := input_regexp.FindAllStringSubmatch(instruction, -1)

	var input_wires []string
	var values []int

	for _, input := range inputs {
		input := input[1]
		if wire_regexp.MatchString(input) {
			input_wires = append(input_wires, input)
		} else {
			value, _ := strconv.Atoi(input)
			values = append(values, value)
		}
	}

	return &Circuit{output_wire, input_wires, values, gate}
}

func (circuit *Circuit) simplify(wire string, value int) {
	for i, input_wire := range circuit.input_wires {
		if input_wire == wire {
			circuit.input_wires = append(circuit.input_wires[:i], circuit.input_wires[i+1:]...)
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

func (circuit *Circuit) has_output() bool {
	return len(circuit.input_wires) == 0
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
	target_wire := "a"

	for _, instruction := range instructions {
		circuit := NewCircuit(instruction)

		for _, input := range circuit.input_wires {
			circuits[input] = append(circuits[input], circuit)
		}

		if circuit.has_output() {
			outputs[circuit.output_wire] = circuit.output()
		}
	}

	outputs["b"] = 16076

	for {
		evaluated := make(map[string]int)

		for wire, wire_value := range outputs {
			for _, circuit := range circuits[wire] {
				circuit.simplify(wire, wire_value)

				if circuit.has_output() {
					evaluated[circuit.output_wire] = circuit.output()
				}
			}

			delete(outputs, wire)
		}

		for wire, wire_value := range evaluated {
			outputs[wire] = wire_value
		}

		final_output, present := outputs[target_wire]
		if present {
			println(final_output)
			return
		}
	}
}
