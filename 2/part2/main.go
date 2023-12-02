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

var boc = Cubes{
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
	var powerSum int
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
		gameImpossible := false
		laoc := Cubes{}

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

					if noc > laoc.red {
						laoc.red = noc
					}

					if noc > boc.red {
						gameImpossible = true
					}

				} else if fospl[2] == "green" {
					noc, err := strconv.Atoi(fospl[1])
					if err != nil {
						log.Fatalf("can't convert %s to int, err: %s", fospl[1], err)
					}

					if noc > laoc.green {
						laoc.green = noc
					}

					if noc > boc.green {
						gameImpossible = true
					}

				} else if fospl[2] == "blue" {
					noc, err := strconv.Atoi(fospl[1])
					if err != nil {
						log.Fatalf("can't convert %s to int, err: %s", fospl[1], err)
					}

					if noc > laoc.blue {
						laoc.blue = noc
					}

					if noc > boc.blue {
						gameImpossible = true
					}
				}
			}
		}

		if gameImpossible {
			log.Printf(" Game ID  %v failed! ", gid)
		} else {
			sum += gid
			log.Printf(" Game ID %v passed! ", gid)
			gameImpossible = false
		}
		powerSum += laoc.red * laoc.green * laoc.blue
	}
	log.Println("Sum of ids of all possible games : ", sum)
	log.Println("Sum of the power of the least possible cubes : ", powerSum)

}
