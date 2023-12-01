package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isDigit(s string) (int, bool) {
	d, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return d, true
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
		var ds []int
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			d, ok := isDigit(char)
			if ok {
				ds = append(ds, d)
			}
		}
		fd := fmt.Sprint(ds[0])
		ld := fmt.Sprint(ds[len(ds)-1])
		log.Printf("First digit: %s", fd)
		log.Printf("Last digit: %s", ld)

		digSum := fd + ld
		n, err := strconv.Atoi(digSum)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Connected digs: %v", n)
		sum += n
	}

	log.Println("Sum: ", sum)

}
