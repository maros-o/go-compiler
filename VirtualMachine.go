package main

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type VM struct {
	data      []interface{}
	labels    map[int]int
	variables map[string]Variable
}

func (vm *VM) Push(item interface{}) {
	vm.data = append(vm.data, item)
}

func (vm *VM) Pop() interface{} {
	if len(vm.data) == 0 {
		return nil
	}
	item := vm.data[len(vm.data)-1]
	vm.data = vm.data[:len(vm.data)-1]
	return item
}

func (vm *VM) PopTwo() (interface{}, interface{}) {
	if len(vm.data) < 2 {
		return nil, nil
	}
	b := vm.data[len(vm.data)-1]
	a := vm.data[len(vm.data)-2]
	vm.data = vm.data[:len(vm.data)-2]
	return a, b
}

func (vm *VM) PrintVM() {
	fmt.Printf("[")
	for i := 0; i < len(vm.data); i++ {
		fmt.Printf("%v", vm.data[i])
		if i < len(vm.data)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")
}

func (vm *VM) IsEmpty() bool {
	return len(vm.data) == 0
}

func splitInstruction(instruction string) []string {
	var parts []string
	var buffer strings.Builder
	inQuote := false

	for _, char := range instruction {
		if char == '"' {
			inQuote = !inQuote
		}
		if char == ' ' && !inQuote {
			if buffer.Len() > 0 {
				parts = append(parts, buffer.String())
				buffer.Reset()
			}
		} else {
			buffer.WriteRune(char)
		}
	}
	if buffer.Len() > 0 {
		parts = append(parts, buffer.String())
	}
	return parts
}

func (vm *VM) Evaluate(instructions []string) {
	vm.data = make([]interface{}, 0)
	vm.labels = make(map[int]int)
	vm.variables = make(map[string]Variable)

	for i := 0; i < len(instructions); i++ {
		parts := splitInstruction(instructions[i])
		opcode := parts[0]
		if opcode == "label" {
			n, _ := strconv.Atoi(parts[1])
			vm.labels[n] = i
		}
	}
	for i := 0; i < len(instructions); i++ {
		parts := splitInstruction(instructions[i])
		opcode := parts[0]

		switch opcode {
		case "push":
			value := parts[2]
			switch parts[1] {
			case "I":
				intValue, _ := strconv.Atoi(value)
				vm.Push(intValue)
			case "F":
				floatValue, _ := strconv.ParseFloat(value, 64)
				vm.Push(floatValue)
			case "S":
				vm.Push(value)
			case "B":
				boolValue, _ := strconv.ParseBool(value)
				vm.Push(boolValue)
			}
		case "print":
			n, _ := strconv.Atoi(parts[1])
			prints := make([]string, n)
			for i := 0; i < n; i++ {
				str := fmt.Sprintf("%v", vm.Pop())
				if str[0] == '"' {
					str = str[1 : len(str)-1]
				}
				prints[i] = str
			}
			for i := n - 1; i >= 0; i-- {
				fmt.Print(color.GreenString(prints[i]))
			}
			println()
		case "pop":
			vm.Pop()
		case "save":
			id := parts[1]
			vm.variables[id] = Variable{id: id, value: vm.Pop()}
		case "load":
			id := parts[1]
			vm.Push(vm.variables[id].value)
		case "label":
		case "jmp":
			n, _ := strconv.Atoi(parts[1])
			labelIdx, ok := vm.labels[n]
			if !ok {
				panic(fmt.Sprintf("Label %d not found", n))
			}
			i = labelIdx
		case "fjmp":
			n, _ := strconv.Atoi(parts[1])
			if !vm.Pop().(bool) {
				labelIdx, ok := vm.labels[n]
				if !ok {
					panic(fmt.Sprintf("Label %d not found", n))
				}
				i = labelIdx
			}
		case "read":
			var input string
			fmt.Printf("Enter value: ")
			_, err := fmt.Scan(&input)
			if err != nil {
				panic(fmt.Sprintf("Error reading input: %vm", err))
			}
			switch parts[1] {
			case "I":
				intValue, ok := strconv.Atoi(input)
				if ok != nil {
					panic(fmt.Sprintf("Error parsing int value: %vm", ok))
				}
				vm.Push(intValue)
			case "F":
				floatValue, ok := strconv.ParseFloat(input, 64)
				if ok != nil {
					panic(fmt.Sprintf("Error parsing float value: %vm", ok))
				}
				vm.Push(floatValue)
			case "S":
				vm.Push(input)
			case "B":
				boolValue, ok := strconv.ParseBool(input)
				if ok != nil {
					panic(fmt.Sprintf("Error parsing boolean value: %vm", ok))
				}
				vm.Push(boolValue)
			}
		case "itof":
			vm.Push(float64(vm.Pop().(int)))
		case "add":
			a, b := vm.PopTwo()
			switch a.(type) {
			case int:
				vm.Push(a.(int) + b.(int))
			case float64:
				vm.Push(a.(float64) + b.(float64))
			}
		case "sub":
			a, b := vm.PopTwo()
			switch a.(type) {
			case int:
				vm.Push(a.(int) - b.(int))
			case float64:
				vm.Push(a.(float64) - b.(float64))
			}
		case "mul":
			a, b := vm.PopTwo()
			switch a.(type) {
			case int:
				vm.Push(a.(int) * b.(int))
			case float64:
				vm.Push(a.(float64) * b.(float64))
			}
		case "div":
			a, b := vm.PopTwo()
			switch a.(type) {
			case int:
				vm.Push(a.(int) / b.(int))
			case float64:
				vm.Push(a.(float64) / b.(float64))
			}
		case "mod":
			a, b := vm.PopTwo()
			vm.Push(a.(int) % b.(int))
		case "uminus":
			a := vm.Pop()
			switch a.(type) {
			case int:
				vm.Push(-a.(int))
			case float64:
				vm.Push(-a.(float64))
			}
		case "not":
			a := vm.Pop()
			vm.Push(!a.(bool))
		case "concat":
			b := vm.Pop().(string)
			a := vm.Pop().(string)
			a = a[1 : len(a)-1]
			b = b[1 : len(b)-1]
			vm.Push(a + b)
		case "and":
			b := vm.Pop().(bool)
			a := vm.Pop().(bool)
			vm.Push(a && b)
		case "or":
			b := vm.Pop().(bool)
			a := vm.Pop().(bool)
			vm.Push(a || b)
		case "gt":
			b := vm.Pop()
			a := vm.Pop()
			switch a.(type) {
			case int:
				vm.Push(a.(int) > b.(int))
			case float64:
				vm.Push(a.(float64) > b.(float64))
			}
		case "lt":
			b := vm.Pop()
			a := vm.Pop()
			switch a.(type) {
			case int:
				vm.Push(a.(int) < b.(int))
			case float64:
				vm.Push(a.(float64) < b.(float64))
			}
		case "eq":
			b := vm.Pop()
			a := vm.Pop()
			vm.Push(a == b)
		}
	}
	if !vm.IsEmpty() {
		panic("VM is not empty")
	}
}
