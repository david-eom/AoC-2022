package main

import (
	"bufio"
	"fmt"
	"os"
)

func enqueue(queue []rune, element rune) []rune {
	return append(queue, element)
}

func dequeue(queue []rune) (rune, []rune) {
	front := queue[0]
	if len(queue) == 1 {
		return front, []rune{}
	}
	return front, queue[1:]
}

var str string

func preprocess() {
	file, err := os.Open("day06.txt")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		str = fileScanner.Text()
	}
	file.Close()
}

func helper(window int) int {
	queue := make([]rune, 0)
	hashset := make(map[rune]bool)
	for idx, char := range str {
		if _, exists := hashset[char]; exists {
			for len(queue) > 0 {
				front, rest := dequeue(queue)
				queue = rest
				if front == char {
					break
				} else {
					delete(hashset, front)
				}
			}
		}
		hashset[char] = true
		queue = enqueue(queue, char)
		if len(queue) == window {
			return idx + 1
		}
	}
	return -1
}

func part1() int {
	return helper(4)
}

func part2() int {
	return helper(14)
}

func main() {
	preprocess()
	fmt.Println(part1())
	fmt.Println(part2())
}
