package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Digit struct {
	strForm string
	intForm int
}

var allDigits = []*Digit{
	{strForm: "one", intForm: 1},
	{strForm: "two", intForm: 2},
	{strForm: "three", intForm: 3},
	{strForm: "four", intForm: 4},
	{strForm: "five", intForm: 5},
	{strForm: "six", intForm: 6},
	{strForm: "seven", intForm: 7},
	{strForm: "eight", intForm: 8},
	{strForm: "nine", intForm: 9},
}

func isDigit(s string) (int, bool) {
	d, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}

	return d, true
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func findStrDigit(s string, isFd bool) int {
	var dsf string
	for i := 0; i < len(s); i++ {
		ss := s[i:]
		for _, digit := range allDigits {
			if len(ss) >= len(digit.strForm) {
				dsf = digit.strForm
				if !isFd {
					dsf = reverseString(digit.strForm)
				}

				if ss[:len(digit.strForm)] == dsf {
					return digit.intForm
				}
			}
		}
	}

	return 0
}

func getFd(line string) int {
	var fd int
	var strDigit string
	for i := 0; i < len(line); i++ {
		char := string(line[i])
		d, ok := isDigit(char)
		if ok {
			res := findStrDigit(strDigit, true)
			if res != 0 {
				fd = res
				strDigit = ""
				break
			}

			fd = d
			strDigit = ""
			break
		} else {
			strDigit += char
		}
	}

	return fd
}

func getLd(line string) int {
	var ld int
	var strDigit string

	for i := len(line) - 1; i >= 0; i-- {
		char := string(line[i])
		d, ok := isDigit(char)
		if ok {
			res := findStrDigit(strDigit, false)
			if res != 0 {
				ld = res
				strDigit = ""
				break
			}
			ld = d
			strDigit = ""
			break
		} else {
			strDigit += char
		}
	}

	return ld
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	var sum int

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		log.Printf("Line: %s", line)
		fd := getFd(line)
		ld := getLd(line)

		if fd == 0 || ld == 0 {
			log.Println("line contains no digits, skipping...")
		}

		log.Printf("First digit: %v", fd)
		log.Printf("Last digit: %v", ld)
		strSum := fmt.Sprint(fd) + fmt.Sprint(ld)
		intSum, err := strconv.Atoi(strSum)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Connected digits: %v", intSum)
		sum += intSum
	}

	log.Println("Sum of calibration values is: ", sum)
}
