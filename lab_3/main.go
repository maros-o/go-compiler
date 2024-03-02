package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SymbolType int

const (
	terminal SymbolType = iota
	nonTerminal
	epsilon
)

type Symbol struct {
	Value rune
	Type  SymbolType
}

func parseRules(text string) map[rune][][]Symbol {
	symbolToRules := make(map[rune][][]Symbol)

	trimmed := strings.TrimRight(text, ";")
	symbolRules := strings.Split(trimmed, ";")
	for _, rule := range symbolRules {
		sides := strings.Split(rule, ":")
		left, right := sides[0], sides[1]
		input := rune(left[0])
		outputs := strings.Split(right, "|")
		for _, output := range outputs {
			symbols := stringToSymbols(output)
			symbolToRules[input] = append(symbolToRules[input], symbols)
		}
	}

	return symbolToRules
}

func stringToSymbols(text string) []Symbol {
	symbols := make([]Symbol, 0)

	for _, r := range text {
		if r == '{' || r == '}' {
			continue
		}
		if r == 'e' {
			symbols = append(symbols, Symbol{Value: r, Type: epsilon})
		} else if r >= 'a' && r <= 'z' {
			symbols = append(symbols, Symbol{Value: r, Type: terminal})
		} else if r >= 'A' && r <= 'Z' {
			symbols = append(symbols, Symbol{Value: r, Type: nonTerminal})
		} else if r == 'e' {
			symbols = append(symbols, Symbol{Value: r, Type: epsilon})
		} else {
			panic("Invalid symbol: " + string(r))
		}
	}
	return symbols
}

func createTable(rules map[rune][][]Symbol) ([][]bool, []rune, map[rune]int) {
	uniqueTerminals := make(map[rune]bool)

	for _, rules := range rules {
		for _, rule := range rules {
			for _, symbol := range rule {
				if symbol.Type == terminal {
					uniqueTerminals[symbol.Value] = true
				}
			}
		}
	}
	runes := make([]rune, 0)

	index := 0
	for k := range rules {
		runes = append(runes, k)
		index++
	}
	for k := range uniqueTerminals {
		runes = append(runes, k)
		index++
	}

	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}

	runeToIndex := make(map[rune]int)
	for i, r := range runes {
		runeToIndex[r] = i
	}

	nonTerminalsCount := len(rules)
	symbolsCount := len(uniqueTerminals) + nonTerminalsCount

	rows := nonTerminalsCount
	cols := symbolsCount

	table := make([][]bool, rows)
	for r := range table {
		table[r] = make([]bool, cols)
		for c := range table[r] {
			table[r][c] = false
		}
	}

	return table, runes, runeToIndex
}

func printTable(table [][]bool, runes []rune) {
	print("  ")
	for n := range runes {
		print(string(runes[n]), " ")
	}
	println()

	for r := range table {
		print(string(runes[r]), " ")
		for c := range table[r] {
			if table[r][c] {
				print("1 ")
			} else {
				print("0 ")
			}
		}
		println()
	}
}

func canOutputBeEmpty(output []Symbol, rules map[rune][][]Symbol) bool {
	for _, symbol := range output {
		if symbol.Type == epsilon {
			continue
		} else if symbol.Type == terminal {
			return false
		} else if symbol.Type == nonTerminal {
			if !canAnyOutputBeEmpty(symbol.Value, rules) {
				return false
			}
		}
	}
	return true
}

func canAnyOutputBeEmpty(input rune, rules map[rune][][]Symbol) bool {
	for _, output := range rules[input] {
		if canOutputBeEmpty(output, rules) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("./lab_3/input.txt")
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

	text := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text += scanner.Text()
	}
	text = strings.ReplaceAll(text, " ", "")

	rules := parseRules(text)
	table, runes, runeToIndex := createTable(rules)
	indexToRune := make(map[int]rune)
	for k, v := range runeToIndex {
		indexToRune[v] = k
	}

	for input, outputs := range rules {
		for _, output := range outputs {
			for _, symbol := range output {
				if symbol.Type == epsilon {
					continue
				} else if symbol.Type == terminal {
					table[runeToIndex[input]][runeToIndex[symbol.Value]] = true
					break
				} else {
					table[runeToIndex[input]][runeToIndex[symbol.Value]] = true
					if !canAnyOutputBeEmpty(symbol.Value, rules) {
						break
					}
				}
			}
		}
	}

	rows := len(table)
	cols := rows
	for {
		tableChanged := false
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if r != c && table[r][c] {
					for i := 0; i < cols; i++ {
						if table[c][i] {
							if !table[r][i] {
								table[r][i] = true
								tableChanged = true
							}
						}
					}
				}
			}
		}
		if !tableChanged {
			break
		}
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r != c && table[r][c] {
				for i := cols; i < len(table[0]); i++ {
					if table[c][i] {
						table[r][i] = true
					}
				}
			}
		}
	}

	printTable(table, runes)

	for input, outputs := range rules {
		for _, output := range outputs {
			print("first[", string(input), ":")
			for _, symbol := range output {
				if symbol.Type == epsilon {
					print("{e}")
				} else {
					print(string(symbol.Value))
				}
			}
			print("] = ")

			first := make(map[rune]bool)
			allCanBeEmpty := true
			for _, symbol := range output {
				if symbol.Type == terminal {
					first[symbol.Value] = true
					allCanBeEmpty = false
					break
				} else if symbol.Type == nonTerminal {
					r := runeToIndex[symbol.Value]
					for i := len(table); i < len(table[0]); i++ {
						if table[r][i] {
							first[indexToRune[i]] = true
						}
					}
					canBeEmpty := canAnyOutputBeEmpty(symbol.Value, rules)
					if !canBeEmpty {
						allCanBeEmpty = false
						break
					}
				}
			}
			if allCanBeEmpty {
				first['ε'] = true
			}
			for k := range first {
				if k == 'ε' {
					print("{e} ")
				} else {
					print(string(k), " ")
				}
			}
			println()
		}
	}
}
