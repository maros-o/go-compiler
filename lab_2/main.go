package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type TokenType int

const (
	Identifier TokenType = iota
	Number
	Operator
	Delimiter
	Keyword
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	prefixes := [...]string{"ID", "NUM", "OP"}

	if t.Type == Keyword {
		return fmt.Sprintf("%v", t.Value)
	} else if t.Type == Delimiter {
		if t.Value == ";" {
			return "SEMICOLON\n"
		} else if t.Value == "(" {
			return "LPAREN\n"
		} else if t.Value == ")" {
			return "RPAREN\n"
		}
	}

	return fmt.Sprintf("%v: %v", prefixes[t.Type], t.Value)
}

func LexicalAnalyzer(input string) []Token {
	var tokens []Token
	var currentToken string
	isComment := false

	for _, char := range input {
		if char == '/' && !isComment {
			isComment = true
		}

		if isComment {
			if char == '\n' {
				isComment = false
			}
			continue
		}

		if unicode.IsSpace(char) {
			if currentToken != "" {
				tokens = append(tokens, createToken(currentToken))
				currentToken = ""
			}
			continue
		}

		if isSpecialSymbol(char) {
			if currentToken != "" {
				tokens = append(tokens, createToken(currentToken))
				currentToken = ""
			}
			tokens = append(tokens, createToken(string(char)))
		} else {
			currentToken += string(char)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, createToken(currentToken))
	}

	return tokens
}

func isSpecialSymbol(char rune) bool {
	specialSymbols := map[rune]bool{
		'(': true,
		')': true,
		';': true,
		'+': true,
		'-': true,
		'*': true,
		'/': true,
	}
	return specialSymbols[char]
}

func createToken(value string) Token {
	// Check for keywords
	keywords := map[string]TokenType{
		"div": Keyword,
		"mod": Keyword,
	}
	if tokenType, exists := keywords[value]; exists {
		return Token{Type: tokenType, Value: value}
	}

	// Check for numbers
	if isNumber(value) {
		return Token{Type: Number, Value: value}
	}

	// Check for identifiers
	if isIdentifier(value) {
		return Token{Type: Identifier, Value: value}
	}

	// Check for operators
	operators := map[string]TokenType{
		"+": TokenType(Operator),
		"-": TokenType(Operator),
		"*": TokenType(Operator),
		"/": TokenType(Operator),
	}
	if tokenType, exists := operators[value]; exists {
		return Token{Type: tokenType, Value: value}
	}

	// Assume it's an identifier if no match is found
	return Token{Type: Identifier, Value: value}
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isIdentifier(s string) bool {
	if len(s) == 0 {
		return false
	}
	if !unicode.IsLetter(rune(s[0])) {
		return false
	}
	for _, char := range s[1:] {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("./lab_2/input.txt")
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

	input := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input += line + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	tokens := LexicalAnalyzer(input)
	for _, token := range tokens {
		fmt.Printf("%v\n", token)
	}
}
