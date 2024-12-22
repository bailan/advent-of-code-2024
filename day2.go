package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) ([][]int) {
	var numbers [][]int
	for _, line := range strings.Split(input, "\n") {
		levels := make([]int, 0)
		for _, levelString := range strings.Split(line, " ") {
			level, _ := strconv.Atoi(levelString)
			levels = append(levels, level)
		}
		numbers = append(numbers, levels)
	}
	return numbers
}

func safe(numbers []int) bool {
	increase, decrease := true, true
	for i := 0; i < len(numbers) - 1; i++ {
		if numbers[i] == numbers[i + 1] {
			return false
		}
		if numbers[i] > numbers[i + 1] {
			increase = false
			if numbers[i] - numbers[i + 1] > 3 {
				return false
			}
		}
		if numbers[i] < numbers[i + 1] {
			decrease = false
			if numbers[i + 1] - numbers[i] > 3 {
				return false
			}
		}
	}
	return increase || decrease
}

func part1(input string) int {
	numbers := parseInput(input)
	total := 0
	for _, levels := range numbers {
		if safe(levels) {
			total += 1
		}
	}
	return total
}

func part2(input string) int {
	numbers := parseInput(input)
	total := 0
	for _, levels := range numbers {
		if safe(levels) {
			total += 1
		} else {
			for i := 0; i < len(levels); i++ {
				var newLevels []int
				for j := 0; j < len(levels); j++ {
					if i != j {
						newLevels = append(newLevels, levels[j])
					}
				}
				if safe(newLevels) {
					total += 1
					break
				}
			}
		}
	}
	return total
}

func main() {
	b, err := os.ReadFile("day2.input")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}