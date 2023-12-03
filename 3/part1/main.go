package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Symbol struct {
	line     int
	position int
}

type Number struct {
	line         int
	number       string
	firstCharPos int
	lastCharPos  int
}

type Symbols []*Symbol
type Numbers []*Number

func isDigit(s string) (int, bool) {
	d, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}

	return d, true
}

func retrieveNumsAndSymbols(lines []string) (Symbols, Numbers) {
	var strDigits string
	var numbers Numbers
	var symbols Symbols
	var fchp int

	for ln, line := range lines {
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if _, ok := isDigit(char); ok {
				strDigits += char
				if fchp == 0 {
					fchp = i
				}
			} else {
				if strDigits != "" {
					numbers = append(numbers, &Number{line: ln, number: strDigits, firstCharPos: fchp, lastCharPos: i - 1})
					strDigits = ""
					fchp = 0
				}

				if char != "." {
					symbols = append(symbols, &Symbol{line: ln, position: i})
				}
			}
		}

		if strDigits != "" {
			numbers = append(numbers, &Number{line: ln, number: strDigits, firstCharPos: fchp, lastCharPos: len(line) - 1})
			strDigits = ""
			fchp = 0
		}
	}

	return symbols, numbers
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		currentLine := scanner.Text()

		lines = append(lines, currentLine)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading from file:", err)
	}
	var sum int
	symbols, numbers := retrieveNumsAndSymbols(lines)
	for _, number := range numbers {
		var nWithSym bool
		for _, symbol := range symbols {
			if symbol.line == number.line || symbol.line == number.line+1 || symbol.line == number.line-1 {
				for i := number.firstCharPos - 1; i <= number.lastCharPos+1; i++ {
					if i == symbol.position {
						nWithSym = true
						n, err := strconv.Atoi(number.number)
						if err != nil {
							log.Fatal(err)
						}
						sum += n
						break
					}
				}
				if nWithSym {
					break
				}
			}
		}
	}

	log.Println("Sum is: ", sum)
}
