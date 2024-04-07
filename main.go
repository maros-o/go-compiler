package main

import (
	"bufio"
	"github.com/antlr4-go/antlr/v4"
	"go-compiler/parser"
	"log"
	"os"
	"time"
)

func eval(input string) {
	is := antlr.NewInputStream(input)
	lexer := parser.NewStackLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewStackParser(stream)

	v := &Visitor{}
	v.Visit(p.Program())
	v.PrintVariables()
}

func testFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	lines := ""
	for scanner.Scan() {
		lines += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	eval(lines)
}

func main() {
	time.Sleep(100 * time.Millisecond)

	testFile("./inputs/input_3.txt")
}
