package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func outcome(opponent string, own string) int {
	if opponent == own {
		return 3
	}
	if opponent == "A" && own == "B" {
		return 6
	}
	if opponent == "B" && own == "C" {
		return 6
	}
	if opponent == "C" && own == "A" {
		return 6
	}
	return 0
}

func part1() int {
	result := 0

	file, err := os.Open("day02.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), " ")
		opponent, own := s[0], s[1]

		if own == "X" {
			result += 1
			own = "A"
		} else if own == "Y" {
			result += 2
			own = "B"
		} else if own == "Z" {
			result += 3
			own = "C"
		}

		result += outcome(opponent, own)
	}
	file.Close()

	return result
}

func strategy(opponent string, game string) int {
	strat := 0
	if opponent == "A" {
		strat = 0
	} else if opponent == "B" {
		strat = 1
	} else if opponent == "C" {
		strat = 2
	}

	if game == "X" {
		strat = (strat + 2) % 3
	} else if game == "Z" {
		strat = (strat + 1) % 3
	}

	return strat + 1
}

func part2() int {
	result := 0

	file, err := os.Open("day02.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), " ")
		opponent, game := s[0], s[1]

		if game == "X" {
			result += 0
		} else if game == "Y" {
			result += 3
		} else if game == "Z" {
			result += 6
		}
		result += strategy(opponent, game)
	}
	file.Close()

	return result
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
