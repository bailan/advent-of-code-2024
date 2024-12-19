package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) (map[string]bool, []string) {
	pattern, design := strings.Split(input, "\n\n")[0], strings.Split(input, "\n\n")[1]
	patterns := make(map[string]bool)
	for _, line := range strings.Split(pattern, ", ") {
		patterns[line] = true
	}
	designs := make([]string, 0)
	for _, line := range strings.Split(design, "\n") {
		designs = append(designs, line)
	}
	return patterns, designs
}

func match(patterns map[string]bool, design string) bool {
	if len(design) == 0 {
		return true
	}
	for i := 1; i <= len(design); i++ {
		if _, ok := patterns[design[:i]]; ok {
			if match(patterns, design[i:]) {
				return true
			}
		}
	}
	return false
}

func part1(input string) int {
	patterns, designs := parseInput(input)
	count := 0
	for _, design := range designs {
		if match(patterns, design) {
			count += 1
		}
	}
	return count
}

var cache = make(map[string]int)

func match2(patterns map[string]bool, design string) int {
	if v, ok := cache[design]; ok {
		return v
	}
	if len(design) == 0 {
		return 1
	}
	count := 0
	for i := 1; i <= len(design); i++ {
		if _, ok := patterns[design[:i]]; ok {
			count += match2(patterns, string(design[i:]))
		}
	}
	cache[design] = count
	return count
}

func part2(input string) int {
	patterns, designs := parseInput(input)
	count := 0
	for _, design := range designs {
		count += match2(patterns, design)
	}
	return count
}

func main() {
	b, err := os.ReadFile("day19.input")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}