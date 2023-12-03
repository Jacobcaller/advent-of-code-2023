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

				if char == "*" {
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
	for _, symbol := range symbols {
		var upNum int
		var downNum int
		var leftNum int
		var rightNum int
		for _, number := range numbers {
			if number.line == symbol.line {
				for i := number.firstCharPos; i <= number.lastCharPos; i++ {
					if i == symbol.position-1 {
						n, err := strconv.Atoi(number.number)
						if err != nil {
							log.Fatal(err)
						}
						leftNum = n
						break
					}
					if i == symbol.position+1 {
						n, err := strconv.Atoi(number.number)
						if err != nil {
							log.Fatal(err)
						}
						rightNum = n
						break
					}
				}
			}
			if number.line == symbol.line-1 {
				for i := number.firstCharPos - 1; i <= number.lastCharPos+1; i++ {
					if i == symbol.position {
						n, err := strconv.Atoi(number.number)
						if err != nil {
							log.Fatal(err)
						}
						upNum = n
						break
					}
				}
			}
			if number.line == symbol.line+1 {
				for i := number.firstCharPos - 1; i <= number.lastCharPos+1; i++ {
					if i == symbol.position {
						n, err := strconv.Atoi(number.number)
						if err != nil {
							log.Fatal(err)
						}
						downNum = n
						break
					}
				}
			}
		}

		if upNum != 0 && downNum != 0 {
			sum += upNum * downNum
		}
		if leftNum != 0 && rightNum != 0 {
			log.Printf("%v * %v = %v", leftNum, rightNum, leftNum+rightNum)
			sum += leftNum * rightNum
		}
	}

	log.Println("Sum of all part numbers is: ", sum)
}
