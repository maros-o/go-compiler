package main

import (
	"fmt"
)

func main() {
	input := "3 + 5 * 2"

	result := calc(input)

	fmt.Printf("%s = %d\n", input, result)
}
