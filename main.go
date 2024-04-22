package main

import (
	"bufio"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/fatih/color"
	"go-compiler/parser"
	"log"
	"os"
	"time"
)

type ColorErrorListener struct {
	*antlr.DefaultErrorListener
}

func (c *ColorErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	color.Red("SyntaxError::line %d:%d %s", line, column, msg)
}

func compile(input string) []string {
	is := antlr.NewInputStream(input)

	lexer := parser.NewStackLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&ColorErrorListener{})

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewStackParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&ColorErrorListener{})

	v := &Visitor{}
	v.Visit(p.Program())

	return v.instructions
}

func testFile(inputFilePath, expectedFilePath string) {
	file, err := os.Open("./files/" + inputFilePath + ".txt")
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

	fmt.Printf("\nTesting " + inputFilePath + ".txt\n")
	instructions := compile(lines)

	file, err = os.Open("./files/" + expectedFilePath + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(file)

	idx := 0
	fmt.Printf("\nInstructions:\n")
	for scanner.Scan() {
		if scanner.Text() != instructions[idx] {
			color.Red("%-5d Expected: %s, got: %s", idx+1, scanner.Text(), instructions[idx])
		} else {
			color.Blue("%-5d %s", idx+1, instructions[idx])
		}
		idx++
	}

	time.Sleep(500 * time.Millisecond)

	stack := VM{}
	fmt.Printf("\nProgram:\n")
	stack.Evaluate(instructions)
}

func main() {
	time.Sleep(100 * time.Millisecond)

	testFile("input_1", "test_1")
	//testFile("input_2", "test_2")
	//testFile("input_3", "test_3")
}
