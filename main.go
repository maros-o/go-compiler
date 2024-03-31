package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	for i := 1; i <= 3; i++ {
		file, err := os.Open("./inputs/input_" + strconv.Itoa(i) + ".txt")
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
}
