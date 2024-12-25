package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) ([][]rune) {
	grid := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}
	return grid
}

func part1(input string) int {
	grid := parseInput(input)

	n, m := len(grid), len(grid[0])
	currentX, currentY := -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '^' {
				currentX, currentY = i, j
				break
			}
		}
	}

	visited := make([][]bool, 0)
	for i := 0; i < n; i++ {
		visited = append(visited, make([]bool, m))
	}

	dirX, dirY := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	direction := 0
	for {
		visited[currentX][currentY] = true
		nextX, nextY := currentX + dirX[direction], currentY + dirY[direction]
		if nextX < 0 || nextY < 0 || nextX >= n || nextY >= m {
			break
		}
		for grid[nextX][nextY] == '#' {
			direction = (direction + 1) % 4
			nextX, nextY = currentX + dirX[direction], currentY + dirY[direction]
		}
		currentX, currentY = nextX, nextY
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] {
				count++
			}
		}
	}
	return count
}


func part2(input string) int {
	grid := parseInput(input)

	n, m := len(grid), len(grid[0])
	startX, startY := -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '^' {
				startX, startY = i, j
				break
			}
		}
	}

	count := 0
	dirX, dirY := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	for blockX := 0; blockX < n; blockX++ {
		for blockY := 0; blockY < m; blockY++ {
			if grid[blockX][blockY] != '.' {
				continue
			}
			visited := make([][][]bool, 0)
			for i := 0; i < n; i++ {
				visited = append(visited, make([][]bool, m))
				for j := 0; j < m; j++ {
					visited[i][j] = make([]bool, 4)
				}
			}
			grid[blockX][blockY] = '#'
			direction := 0
			loop := false
			currentX, currentY := startX, startY
			for {
				visited[currentX][currentY][direction] = true
				nextX, nextY := currentX + dirX[direction], currentY + dirY[direction]
				if nextX < 0 || nextY < 0 || nextX >= n || nextY >= m {
					break
				}
				for grid[nextX][nextY] == '#' {
					direction = (direction + 1) % 4
					nextX, nextY = currentX + dirX[direction], currentY + dirY[direction]
				}
				if visited[nextX][nextY][direction] {
					loop = true
					break
				}
				currentX, currentY = nextX, nextY
			}
			if loop {
				count += 1
			}
			grid[blockX][blockY] = '.'
		}
	}
	return count
}

func main() {
	b, err := os.ReadFile("day6.input")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}