package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	result := 0
	for _, match := range matches {
		number1, _ := strconv.Atoi(match[1])
		number2, _ := strconv.Atoi(match[2])
		result += number1 * number2
	}
	return result
}

func part2(input string) int {
	splits := strings.Split(input, "don't()")
	result := part1(splits[0])
	for _, split := range splits[1:] {
		for _, doSplit := range strings.Split(split, "do()")[1:] {
			result += part1(doSplit)
		}
	}
	return result
}


func main() {
	b, err := os.ReadFile("day3.input")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}