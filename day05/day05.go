package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var procedures [][3]int
var crates []string
var stacks_1 [][]byte
var stacks_2 [][]byte

func preprocess() {
	file, err := os.Open("day05.txt")
	if err != nil {
		panic(err)
	}
	is_crates := true

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			is_crates = false
		} else if is_crates {
			crates = append(crates, line)
		} else {
			left := strings.Split(line, " from ")[0]
			right := strings.Split(line, " from ")[1]
			move, _ := strconv.Atoi(strings.Split(left, "move ")[1])
			from, _ := strconv.Atoi(strings.Split(right, " to ")[0])
			to, _ := strconv.Atoi(strings.Split(right, " to ")[1])
			procedures = append(procedures, [3]int{move, from - 1, to - 1})
		}
	}

	last_line := crates[len(crates)-1]
	num_stacks, _ := strconv.Atoi(last_line[len(last_line)-2 : len(last_line)-1])
	for i := 0; i < num_stacks; i++ {
		var stack_1 []byte
		var stack_2 []byte
		pos := 1 + i*4
		for row := len(crates) - 2; row >= 0; row-- {
			letter := crates[row][pos]
			if letter != ' ' {
				stack_1 = append(stack_1, letter)
				stack_2 = append(stack_2, letter)
			}
		}
		stacks_1 = append(stacks_1, stack_1)
		stacks_2 = append(stacks_2, stack_2)
	}
}

func part1() string {
	for _, proc := range procedures {
		for i := 0; i < proc[0]; i++ {
			from := stacks_1[proc[1]]
			to := stacks_1[proc[2]]
			item := from[len(from)-1]
			stacks_1[proc[1]] = from[:len(from)-1]
			stacks_1[proc[2]] = append(to, item)
		}
	}
	var result []byte
	for i := 0; i < len(stacks_1); i++ {
		if len(stacks_1[i]) > 0 {
			result = append(result, stacks_1[i][len(stacks_1[i])-1])
		}
	}

	return string(result)
}

func part2() string {
	for _, proc := range procedures {
		num_crates := proc[0]
		from := stacks_2[proc[1]]
		to := stacks_2[proc[2]]
		items := from[len(from)-num_crates:]
		stacks_2[proc[1]] = from[:len(from)-num_crates]
		stacks_2[proc[2]] = append(to, items...)
	}
	var result []byte
	for i := 0; i < len(stacks_2); i++ {
		if len(stacks_2[i]) > 0 {
			result = append(result, stacks_2[i][len(stacks_2[i])-1])
		}
	}

	return string(result)
}

func main() {
	preprocess()
	fmt.Println(part1())
	fmt.Println(part2())
}
