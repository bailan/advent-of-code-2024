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

func (p Point) toString() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func parseInput(input string) ([]Point)  {
	var points []Point
	for _, line := range strings.Split(input, "\n") {
		var point Point
		fmt.Sscanf(line, "%d,%d", &point.X, &point.Y)
		points = append(points, point)
	}
	return points
}

func part1(input string, n int, m int) (int, error) {
	walls := parseInput(input)
	grid := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			grid[i][j] = false
		}
	}

	for i, wall := range walls {
		if i < m {
			grid[wall.X][wall.Y] = true
		}
	}

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			visited[i][j] = false
		}
	}

	queue := []Point{Point{0, 0}}
	visited[0][0] = true
	step := 0
	for len(queue) > 0 {
		var next []Point
		for _, current := range queue {
			if current.X == n - 1 && current.Y == n - 1 {
				return step, nil
			}
			for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				newX := current.X + dir.X
				newY := current.Y + dir.Y
				if newX >= 0 && newX < n && newY >= 0 && newY < n && !visited[newX][newY] && !grid[newX][newY] {
					next = append(next, Point{newX, newY})
					visited[newX][newY] = true
				}
			}
		}
		queue = next
		step += 1
	}
	return -1, fmt.Errorf("No path found")
}

func find(parent [][]int, a Point) int {
	aParent := parent[a.X][a.Y]
	x := aParent / len(parent)
	y := aParent % len(parent)
	if x == a.X && y == a.Y {
		return aParent
	}
	parent[a.X][a.Y] = find(parent, Point{x, y})
	return parent[a.X][a.Y]
}

func union(parent [][]int, a Point, b Point) {
	aParent := find(parent, a)
	bParent := find(parent, b)
	if aParent != bParent {
		parent[aParent / len(parent)][aParent % len(parent)] = min(aParent, bParent)
		parent[bParent / len(parent)][bParent % len(parent)] = min(aParent, bParent)
	}
}

func part2(input string, nn int) (string, error) {
	n := nn + 2
	walls := parseInput(input)
	
	grid := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			grid[i][j] = false
		}
	}

	parent := make([][]int, n)
	for i := 0; i < n; i++ {
		parent[i] = make([]int, n)
		for j := 0; j < n; j++ {
			parent[i][j] = i * n + j
		}
	}

/**
Extend the grid by one in each direction
Set up two walls for top-right and left-bottom edges
Add walls until the two walls connected

Example of 7x7 grid
..#######
.S......#
#.......#
#.......#
#.......#
#.......#
#.......#
#......E.
#######..
*/
	grid[0][2] = true
	for j := 3; j < n; j++ {
		union(parent, Point{0, j}, Point{0, j - 1})
		grid[0][j] = true
	}
	for i := 1; i < n - 2; i++ {
		union(parent, Point{i, n - 1}, Point{i - 1,  n - 1})
		grid[i][n - 1] = true
	}

	grid[2][0] = true
	for i := 3; i < n; i++ {
		union(parent, Point{i, 0}, Point{i - 1, 0})
		grid[i][0] = true
	}
	for j := 1; j < n - 2; j++ {
		union(parent, Point{n - 1, j}, Point{n - 1, j - 1})
		grid[n - 1][j] = true
	}

	for _, wall := range walls {
		grid[wall.X + 1][wall.Y + 1] = true
		for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
			newX := wall.X + 1 + dir.X
			newY := wall.Y + 1 + dir.Y
			if grid[newX][newY] {
				union(parent, Point{wall.X + 1, wall.Y + 1}, Point{newX, newY})
			}
		}
		if find(parent, Point{0, 2}) == find(parent, Point{2, 0}) {
			return wall.toString(), nil
		}
	}
	return "", fmt.Errorf("No byte found")
}

func main() {
	b, err := os.ReadFile("day18.input")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	step, error := part1(input, 101, 1024)
	if error == nil {
		fmt.Println(step)
	}
	byte, error := part2(input, 101)
	if error == nil {
		fmt.Println(byte)
	}
}