package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Precedence = map[rune]int{
	'+': 0,
	'-': 0,
	'*': 1,
	'/': 1,
}

var LeftAssociative = map[rune]bool{
	'+': true,
	'-': true,
	'*': true,
	'/': true,
}

type tokenType int

const (
	Operand tokenType = iota
	Operator
	OpenBracket
	CloseBracket
)

type token struct {
	Value int32
	Type  tokenType
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	size, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Error reading first line:", err)
		return
	}
	for i := 0; i < size; i++ {
		scanner.Scan()
		evalLine(scanner.Text())
	}
}

func evalLine(line string) {
	postfix, err := infixToPostfix(line)
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	result, err := evalPostfix(postfix)
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	fmt.Println(result)
}

func infixToPostfix(infix string) ([]token, error) {
	var stack []rune
	var output []token
	var number []rune
	for _, v := range infix {
		if strings.ContainsRune(" \t\n", v) {
			continue
		}
		vtokenType, err := runeTotokenType(v)
		if err != nil {
			return nil, err
		}
		if vtokenType == Operand {
			number = append(number, v)
			continue
		} else if len(number) > 0 {
			numInt32, err := runesToInt32(number)
			if err != nil {
				return nil, err
			}
			output = append(output, token{Value: numInt32, Type: Operand})
			number = []rune{}
		}
		if vtokenType == Operator {
			for len(stack) > 0 {
				y := stack[len(stack)-1]
				ytokenType, _ := runeTotokenType(y)
				if ytokenType != Operator {
					break
				}
				if (LeftAssociative[v] && Precedence[v] <= Precedence[y]) || (!LeftAssociative[v] && Precedence[v] < Precedence[y]) {
					stack = stack[:len(stack)-1]
					output = append(output, token{Value: y, Type: Operator})
				} else {
					break
				}
			}
			stack = append(stack, v)
		} else if vtokenType == OpenBracket {
			stack = append(stack, v)
		} else if vtokenType == CloseBracket {
			if len(stack) == 0 {
				return nil, errors.New("unmatched closing bracket")
			}
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top == '(' {
					break
				}
				toptokenType, _ := runeTotokenType(top)
				output = append(output, token{Value: top, Type: toptokenType})
			}
		}
	}
	if len(number) > 0 {
		numInt32, err := runesToInt32(number)
		if err != nil {
			return nil, err
		}
		output = append(output, token{Value: numInt32, Type: Operand})
		number = []rune{}
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		toptokenType, _ := runeTotokenType(top)

		if toptokenType == Operand {
			number = append(number, top)
		} else if len(number) > 0 {
			numInt32, err := runesToInt32(number)
			if err != nil {
				return nil, err
			}
			output = append(output, token{Value: numInt32, Type: Operand})
			number = []rune{}
		}
		if toptokenType == OpenBracket {
			return nil, errors.New("unmatched opening bracket")
		} else if toptokenType == CloseBracket {
			return nil, errors.New("unmatched closing bracket")
		}
		output = append(output, token{Value: top, Type: toptokenType})
	}
	return output, nil
}

func runeTotokenType(r rune) (tokenType, error) {
	if r >= '0' && r <= '9' {
		return Operand, nil
	}
	if strings.ContainsRune("+-*/", r) {
		return Operator, nil
	}
	if r == '(' {
		return OpenBracket, nil
	}
	if r == ')' {
		return CloseBracket, nil
	}
	return 0, errors.New("invalid token")
}

func runesToInt32(runes []rune) (int32, error) {
	numStr := string(runes)
	numInt, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, errors.New("invalid number")
	}
	return int32(numInt), nil
}

func evalPostfix(postfix []token) (int, error) {
	var stack []token
	for _, v := range postfix {
		if v.Type == Operand {
			stack = append(stack, v)
		} else if v.Type == Operator {
			if len(stack) < 2 {
				return 0, errors.New("stack has less than 2 elements")
			}
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			if a.Type != Operand || b.Type != Operand {
				return 0, errors.New("invalid operands in stack")
			}
			stack = stack[:len(stack)-2]
			eval, err := evalOp(int(b.Value), int(a.Value), v.Value)
			if err != nil {
				return 0, err
			}
			stack = append(stack, token{Value: int32(eval), Type: Operand})
		} else {
			return 0, errors.New("invalid token")
		}
	}
	if len(stack) != 1 {
		return 0, errors.New("stack is not empty")
	}
	return int(stack[0].Value), nil
}

func evalOp(a, b int, op rune) (int, error) {
	switch op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	}
	return 0, errors.New("invalid operator")
}
