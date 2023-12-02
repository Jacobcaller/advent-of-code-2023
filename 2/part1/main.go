package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cubes struct {
	red   int
	green int
	blue  int
}

var boc = &Cubes{
	red: 12, green: 13, blue: 14,
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func getGid(s string) int {
	var gidStr string
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if isDigit(char) {
			gidStr += char
		}
	}
	gid, err := strconv.Atoi(gidStr)
	if err != nil {
		return 0
	}
	return gid
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
		fspl := strings.Split(line, ":")
		game := fspl[0]
		gid := getGid(game)
		rots := fspl[1]
		gameFailed := false

		// SSPL are all sets of one game
		sspl := strings.Split(rots, ";")
		for _, s := range sspl {
			// TSPL are all the cubes of these sets
			tspl := strings.Split(s, ",")
			for _, s := range tspl {
				fospl := strings.Split(s, " ")
				if fospl[2] == "red" {
					noc, err := strconv.Atoi(fospl[1])
					if err != nil {
						log.Fatalf("can't convert %s to int, err: %s", fospl[1], err)
					}
					if noc > boc.red {
						log.Printf("RED %v is bigger than %v", noc, boc.red)
						gameFailed = true
					}
				} else if fospl[2] == "green" {
					noc, err := strconv.Atoi(fospl[1])
					if err != nil {
						log.Fatalf("can't convert %s to int, err: %s", fospl[1], err)
					}
					if noc > boc.green {
						log.Printf("GREEN %v is bigger than %v", noc, boc.green)
						gameFailed = true
					}
				} else if fospl[2] == "blue" {
					noc, err := strconv.Atoi(fospl[1])
					if err != nil {
						log.Fatalf("can't convert %s to int, err: %s", fospl[1], err)
					}
					if noc > boc.blue {
						gameFailed = true
					}
				}
			}
		}

		if gameFailed {
			log.Println(" Game ID failed: ", gid)
		} else {
			sum += gid
			log.Println(" Game ID passed: ", gid)
			gameFailed = false
		}
	}
	log.Println("Sum : ", sum)
}
