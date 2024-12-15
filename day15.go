package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Grid [][]rune

func (grid Grid) print() {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func (grid Grid) score() int {
	score := 0
	for x, row := range grid {
		for y, cell := range row {
			if cell == 'O' || cell == '[' {
				score += x*100 + y
			}
		}
	}
	return score
}

func (grid Grid) robot() (Point, error) {
	n, m := len(grid), len(grid[0])
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if grid[x][y] == '@' {
				return Point{x, y}, nil
			}
		}
	}
	return Point{}, fmt.Errorf("No robot found")
}

func parseInput(input string) (Grid, []Point) {
	grid_string, moves_string := strings.Split(input, "\n\n")[0], strings.Split(input, "\n\n")[1]

	lines := strings.Split(grid_string, "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	var moves []Point
	for _, char := range moves_string {
		switch char {
		case '^':
			moves = append(moves, Point{-1, 0})
		case 'v':
			moves = append(moves, Point{1, 0})
		case '<':
			moves = append(moves, Point{0, -1})
		case '>':
			moves = append(moves, Point{0, 1})
		}
	}
	return grid, moves
}

func part1(input string) int {
	grid, moves := parseInput(input)
	robot, _ := grid.robot()

	for _, direction := range moves {
		if pushable(grid, robot, direction) {
			push(grid, robot, direction)
			robot.X += direction.X
			robot.Y += direction.Y
		}
		// grid.print()
	}
	return grid.score()
}

func part2(input string) int {
	return part1(strings.Replace(strings.Replace(strings.Replace(strings.Replace(input, "#", "##", -1), "O", "[]", -1), ".", "..", -1), "@", "@.", -1))
}

func pushable(grid Grid, position Point, direction Point) bool {
	switch grid[position.X + direction.X][position.Y + direction.Y] {
	case '#':
		return false
	case '.':
		return true
	case 'O':
		return pushable(grid, Point{position.X + direction.X, position.Y + direction.Y}, direction)
	case '[':
		switch direction {
		case Point{0, 1}:
			return pushable(grid, Point{position.X + direction.X, position.Y + direction.Y + 1}, direction)
		case Point{0, -1}:
			return false
		case Point{1, 0}, Point{-1, 0}:
			return pushable(grid, Point{position.X + direction.X, position.Y + direction.Y}, direction) &&
				pushable(grid, Point{position.X + direction.X, position.Y + direction.Y + 1}, direction)
		}
	case ']':
		switch direction {
		case Point{0, -1}:
			return pushable(grid, Point{position.X + direction.X, position.Y + direction.Y - 1}, direction)
		case Point{0, 1}:
			return false
		case Point{1, 0}, Point{-1, 0}:
			return pushable(grid, Point{position.X + direction.X, position.Y + direction.Y}, direction) &&
				pushable(grid, Point{position.X + direction.X, position.Y + direction.Y - 1}, direction)
		}
	}
	return false
}

func push(grid Grid, position Point, direction Point) {
	switch grid[position.X + direction.X][position.Y + direction.Y] {
	case '#':
	case '.':
		grid[position.X + direction.X][position.Y + direction.Y] = grid[position.X][position.Y]
	case 'O':
		push(grid, Point{position.X + direction.X, position.Y + direction.Y}, direction)
		grid[position.X + direction.X][position.Y + direction.Y] = grid[position.X][position.Y]
	case '[':
		switch direction {
		case Point{1, 0}, Point{-1, 0}:
			push(grid, Point{position.X + direction.X, position.Y + direction.Y + 1}, direction)
		}
		push(grid, Point{position.X + direction.X, position.Y + direction.Y}, direction)
		grid[position.X + direction.X][position.Y + direction.Y] = grid[position.X][position.Y]
	case ']':
		switch direction {
		case Point{1, 0}, Point{-1, 0}:
			push(grid, Point{position.X + direction.X, position.Y + direction.Y - 1}, direction)
		}
		push(grid, Point{position.X + direction.X, position.Y + direction.Y}, direction)
		grid[position.X + direction.X][position.Y + direction.Y] = grid[position.X][position.Y]
	}
	grid[position.X][position.Y] = '.'
}

func main() {
	b, err := os.ReadFile("day15.input")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
