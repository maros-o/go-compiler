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

func (s *VM) Push(item interface{}) {
	s.data = append(s.data, item)
}

func (s *VM) Pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s *VM) PopTwo() (interface{}, interface{}) {
	if len(s.data) < 2 {
		return nil, nil
	}
	b := s.data[len(s.data)-1]
	a := s.data[len(s.data)-2]
	s.data = s.data[:len(s.data)-2]
	return a, b
}

func (s *VM) PrintVM() {
	fmt.Printf("[")
	for i := 0; i < len(s.data); i++ {
		fmt.Printf("%v", s.data[i])
		if i < len(s.data)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")
}

func (s *VM) IsEmpty() bool {
	return len(s.data) == 0
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

func (s *VM) Evaluate(instructions []string) {
	s.data = make([]interface{}, 0)
	s.labels = make(map[int]int)
	s.variables = make(map[string]Variable)

	for i := 0; i < len(instructions); i++ {
		parts := splitInstruction(instructions[i])
		opcode := parts[0]
		if opcode == "label" {
			n, _ := strconv.Atoi(parts[1])
			s.labels[n] = i
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
				s.Push(intValue)
			case "F":
				floatValue, _ := strconv.ParseFloat(value, 64)
				s.Push(floatValue)
			case "S":
				s.Push(value)
			case "B":
				boolValue, _ := strconv.ParseBool(value)
				s.Push(boolValue)
			}
		case "print":
			n, _ := strconv.Atoi(parts[1])
			prints := make([]string, n)
			for i := 0; i < n; i++ {
				str := fmt.Sprintf("%v", s.Pop())
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
			s.Pop()
		case "save":
			id := parts[1]
			s.variables[id] = Variable{id: id, value: s.Pop()}
		case "load":
			id := parts[1]
			s.Push(s.variables[id].value)
		case "label":
		case "jmp":
			n, _ := strconv.Atoi(parts[1])
			labelIdx, ok := s.labels[n]
			if !ok {
				panic(fmt.Sprintf("Label %d not found", n))
			}
			i = labelIdx
		case "fjmp":
			n, _ := strconv.Atoi(parts[1])
			if !s.Pop().(bool) {
				labelIdx, ok := s.labels[n]
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
				panic(fmt.Sprintf("Error reading input: %s", err))
			}
			switch parts[1] {
			case "I":
				intValue, ok := strconv.Atoi(input)
				if ok != nil {
					panic(fmt.Sprintf("Error parsing int value: %s", ok))
				}
				s.Push(intValue)
			case "F":
				floatValue, ok := strconv.ParseFloat(input, 64)
				if ok != nil {
					panic(fmt.Sprintf("Error parsing float value: %s", ok))
				}
				s.Push(floatValue)
			case "S":
				s.Push(input)
			case "B":
				boolValue, ok := strconv.ParseBool(input)
				if ok != nil {
					panic(fmt.Sprintf("Error parsing boolean value: %s", ok))
				}
				s.Push(boolValue)
			}
		case "itof":
			s.Push(float64(s.Pop().(int)))
		case "add":
			a, b := s.PopTwo()
			switch a.(type) {
			case int:
				s.Push(a.(int) + b.(int))
			case float64:
				s.Push(a.(float64) + b.(float64))
			}
		case "sub":
			a, b := s.PopTwo()
			switch a.(type) {
			case int:
				s.Push(a.(int) - b.(int))
			case float64:
				s.Push(a.(float64) - b.(float64))
			}
		case "mul":
			a, b := s.PopTwo()
			switch a.(type) {
			case int:
				s.Push(a.(int) * b.(int))
			case float64:
				s.Push(a.(float64) * b.(float64))
			}
		case "div":
			a, b := s.PopTwo()
			switch a.(type) {
			case int:
				s.Push(a.(int) / b.(int))
			case float64:
				s.Push(a.(float64) / b.(float64))
			}
		case "mod":
			a, b := s.PopTwo()
			s.Push(a.(int) % b.(int))
		case "uminus":
			a := s.Pop()
			switch a.(type) {
			case int:
				s.Push(-a.(int))
			case float64:
				s.Push(-a.(float64))
			}
		case "not":
			a := s.Pop()
			s.Push(!a.(bool))
		case "concat":
			b := s.Pop().(string)
			a := s.Pop().(string)
			a = a[1 : len(a)-1]
			b = b[1 : len(b)-1]
			s.Push(a + b)
		case "and":
			b := s.Pop().(bool)
			a := s.Pop().(bool)
			s.Push(a && b)
		case "or":
			b := s.Pop().(bool)
			a := s.Pop().(bool)
			s.Push(a || b)
		case "gt":
			b := s.Pop()
			a := s.Pop()
			switch a.(type) {
			case int:
				s.Push(a.(int) > b.(int))
			case float64:
				s.Push(a.(float64) > b.(float64))
			}
		case "lt":
			b := s.Pop()
			a := s.Pop()
			switch a.(type) {
			case int:
				s.Push(a.(int) < b.(int))
			case float64:
				s.Push(a.(float64) < b.(float64))
			}
		case "eq":
			b := s.Pop()
			a := s.Pop()
			s.Push(a == b)
		}
	}
	if !s.IsEmpty() {
		panic("VM is not empty")
	}
}
