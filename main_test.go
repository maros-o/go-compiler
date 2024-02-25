package main

import (
	"reflect"
	"testing"
)

func TestRuneToTokenType(t *testing.T) {
	tests := []struct {
		input rune
		want  TokenType
	}{
		{'0', Operand},
		{'+', Operator},
		{'(', OpenBracket},
		{')', CloseBracket},
		{'%', 0},
	}
	for _, test := range tests {
		got, err := runeToTokenType(test.input)
		if err != nil && test.want != 0 {
			t.Errorf("runeToTokenType(%v) returned an error, expected %v", test.input, test.want)
		}
		if got != test.want {
			t.Errorf("runeToTokenType(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestInfixToPostfix(t *testing.T) {
	tests := []struct {
		input string
		want  []Token
	}{
		{"3 + 4 * 2 / ( 1 - 5 )", []Token{
			{Value: 3, Type: Operand},
			{Value: 4, Type: Operand},
			{Value: 2, Type: Operand},
			{Value: '*', Type: Operator},
			{Value: 1, Type: Operand},
			{Value: 5, Type: Operand},
			{Value: '-', Type: Operator},
			{Value: '/', Type: Operator},
			{Value: '+', Type: Operator},
		}},
		{"lol", nil},
	}
	for _, test := range tests {
		got, err := infixToPostfix(test.input)
		if err != nil && test.want != nil {
			t.Errorf("infixToPostfix(%v) returned an error, expected %v", test.input, test.want)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("infixToPostfix(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestRunesToInt32(t *testing.T) {
	tests := []struct {
		input []rune
		want  int32
	}{
		{[]rune{'1', '2', '3'}, 123},
		{[]rune{'-', '4', '5'}, -45},
		{[]rune{'a', 'b', 'c'}, 0},
	}
	for _, test := range tests {
		got, err := runesToInt32(test.input)
		if err != nil && test.want != 0 {
			t.Errorf("runesToInt32(%v) returned an error, expected %v", test.input, test.want)
		}
		if got != test.want {
			t.Errorf("runesToInt32(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestEvalPostfix(t *testing.T) {
	tests := []struct {
		input []Token
		want  int
	}{
		{[]Token{
			{Value: 3, Type: Operand},
			{Value: 4, Type: Operand},
			{Value: '+', Type: Operator},
		}, 7},
		{nil, 0},
		{[]Token{
			{Value: 3, Type: Operand},
			{Value: 3, Type: Operand},
			{Value: '*', Type: Operator},
		}, 9},
		{[]Token{
			{Value: 3, Type: Operand},
			{Value: 4, Type: Operand},
			{Value: 2, Type: Operand},
			{Value: '*', Type: Operator},
			{Value: 1, Type: Operand},
			{Value: 5, Type: Operand},
			{Value: '-', Type: Operator},
			{Value: '/', Type: Operator},
			{Value: '+', Type: Operator},
		}, 1},
	}
	for _, test := range tests {
		got, err := evalPostfix(test.input)
		if err != nil && test.want != 0 {
			t.Errorf("evalPostfix(%v) returned an error, expected %v", test.input, test.want)
		}
		if got != test.want {
			t.Errorf("evalPostfix(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
