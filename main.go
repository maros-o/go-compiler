package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

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

	testFile("./inputs/lab_8.txt")
}
