package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var elves []int

func preprocess() {
	file, err := os.Open("day01.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	elf := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			elves = append(elves, elf)
			elf = 0
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}
		elf += num
	}
	elves = append(elves, elf)
	sort.Ints(elves)
	file.Close()
}

func part1() int {
	return elves[len(elves)-1]
}

func part2() int {
	result := 0
	for i := len(elves) - 1; i > len(elves)-4; i-- {
		result += elves[i]
	}
	return result
}

func main() {
	preprocess()
	fmt.Println(part1())
	fmt.Println(part2())
}
