package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) ([][]rune) {
	var puzzle [][]rune
	for _, line := range strings.Split(input, "\n") {
		puzzle = append(puzzle, []rune(line))
	}
	return puzzle
}

func part1(input string) int {
	puzzle := parseInput(input)

	dirx := []int{1, 1, 0, -1, -1, -1, 0, 1}
	diry := []int{0, 1, 1, 1, 0, -1, -1, -1}
	word := "XMAS"
	count := 0
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			for dir := 0; dir < 8; dir++ {
				for k := 0; k < len(word); k++ {
					if i + k * dirx[dir] >= 0 && i + k * dirx[dir] < len(puzzle) && j + k * diry[dir] >= 0 && j + k * diry[dir] < len(puzzle[i]) {
						if puzzle[i + k * dirx[dir]][j + k * diry[dir]] != rune(word[k]) {
							break
						}
						if k == len(word) - 1 {
							count += 1
						}
					}
				}
			}
		}
	}
	return count
}

func part2(input string) int {
	puzzle := parseInput(input)
	count := 0
	for i := 1; i < len(puzzle) - 1; i++ {
		for j := 1; j < len(puzzle[i]) - 1; j++ {
			if puzzle[i][j] != 'A' {
				continue
			}
			if puzzle[i - 1][j - 1] == 'M' && puzzle[i - 1][j + 1] == 'M' && puzzle[i + 1][j + 1] == 'S' && puzzle[i + 1][j - 1] == 'S' {
				count += 1
				continue
			}
			if puzzle[i - 1][j - 1] == 'S' && puzzle[i - 1][j + 1] == 'M' && puzzle[i + 1][j + 1] == 'M' && puzzle[i + 1][j - 1] == 'S' {
				count += 1
				continue
			}
			if puzzle[i - 1][j - 1] == 'S' && puzzle[i - 1][j + 1] == 'S' && puzzle[i + 1][j + 1] == 'M' && puzzle[i + 1][j - 1] == 'M' {
				count += 1
				continue
			}
			if puzzle[i - 1][j - 1] == 'M' && puzzle[i - 1][j + 1] == 'S' && puzzle[i + 1][j + 1] == 'S' && puzzle[i + 1][j - 1] == 'M' {
				count += 1
				continue
			}
		}
	}
	return count
}

func main() {
	b, err := os.ReadFile("day4.input")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}