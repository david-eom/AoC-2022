package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pairs [][][]int

func strToInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return number
}

func contain(a1, a2, b1, b2 int) bool {
	if a1 < b1 {
		return false
	}
	if a2 > b2 {
		return false
	}
	return true
}

func overlap(a1, a2, b1, b2 int) bool {
	return a1 <= b2 && b1 <= a2
}

func preprocess() {
	file, err := os.Open("day04.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), ",")
		p1 := strings.Split(line[0], "-")
		p2 := strings.Split(line[1], "-")
		intervals := [][]int{
			{strToInt((p1[0])), strToInt((p1[1]))},
			{strToInt((p2[0])), strToInt((p2[1]))},
		}
		pairs = append(pairs, intervals)
	}
	file.Close()
}

func part1() int {
	result := 0
	for _, p := range pairs {
		if contain(p[0][0], p[0][1], p[1][0], p[1][1]) ||
			contain(p[1][0], p[1][1], p[0][0], p[0][1]) {
			result += 1
		}
	}
	return result
}

func part2() int {
	result := 0
	for _, p := range pairs {
		if overlap(p[0][0], p[0][1], p[1][0], p[1][1]) ||
			overlap(p[1][0], p[1][1], p[0][0], p[0][1]) {
			result += 1
		}
	}
	return result
}

func main() {
	preprocess()
	fmt.Println(part1())
	fmt.Println(part2())
}
