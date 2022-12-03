package main

import (
	"bufio"
	"fmt"
	"os"
)

var rucksacks []string

func priority(char rune) rune {
	if char >= 97 {
		return char - 96
	} else {
		return char - 38
	}
}

func preprocess() {
	file, err := os.Open("day03.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		rucksacks = append(rucksacks, fileScanner.Text())
	}
	file.Close()
}

func part1() int {
	result := 0
	for _, rucksack := range rucksacks {
		m := make(map[rune]bool)
		front := rucksack[:len(rucksack)/2]
		back := rucksack[len(rucksack)/2:]
		for _, char := range front {
			m[char] = true
		}
		for _, char := range back {
			if _, exists := m[char]; exists {
				result += int(priority(char))
				break
			}
		}
	}
	return result
}

func part2() int {
	result := 0
	for i := 0; i < len(rucksacks); i = i + 3 {
		a := rucksacks[i]
		b := rucksacks[i+1]
		c := rucksacks[i+2]
		ma := make(map[rune]bool)
		mb := make(map[rune]bool)
		for _, char := range a {
			ma[char] = true
		}
		for _, char := range b {
			mb[char] = true
		}
		for _, char := range c {
			if _, exists := ma[char]; !exists {
				continue
			}
			if _, exists := mb[char]; !exists {
				continue
			}
			result += int(priority(char))
			break
		}
	}

	return result
}

func main() {
	preprocess()
	fmt.Println(part1())
	fmt.Println(part2())
}
