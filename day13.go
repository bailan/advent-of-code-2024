package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func parseNumbers(input string) ([6]int, error) {
	re := regexp.MustCompile(`[+-]?\d+`)
	matches := re.FindAllString(input, -1)
    var numbers [6]int
	if len(matches) != 6 {
		return numbers, fmt.Errorf("expected 6 numbers, found %d", len(matches))
	}
	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return numbers, err
		}
		numbers[i] = num
	}
	return numbers, nil
}

func parseInput(input string) ([][6]int, error) {
    var machines [][6]int
    lines := strings.Split(input, "\n")
	for i, _ := range lines {
		if i % 4 == 3 {
            numbers, err := parseNumbers(strings.Join(lines[i - 3:i], " "))
            if err != nil {
                return machines, err
            }
            machines = append(machines, numbers)
        }
	}
    return machines, nil
}

func solveEquations(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int) (int, int) {
	det := x1*y2 - x2*y1
	if det == 0 {
		return 0, 0
	}
    if (y2*x3 - x2*y3) % det != 0 || (y2*x3 - x2*y3) / det < 0 {
        return 0, 0
    }
    if (-y1*x3 + x1*y3) % det != 0 || (-y1*x3 + x1*y3) / det < 0 {
        return 0, 0
    }
	return (y2*x3 - x2*y3) / det, (-y1*x3 + x1*y3) / det
}

func part1(input string) (int) {
    machines, err := parseInput(input)
    if err != nil {
        fmt.Print(err)
    }
    var score = 0
    for _, numbers := range machines {
        a, b := solveEquations(numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5])
        score += a * 3 + b * 1
    }
    return score
}

func part2(input string) (int) {
    machines, err := parseInput(input)
    if err != nil {
        fmt.Print(err)
    }
    var score = 0
    for _, numbers := range machines {
        a, b := solveEquations(numbers[0], numbers[1], numbers[2], numbers[3], 10000000000000 + numbers[4], 10000000000000 + numbers[5])
        score += a * 3 + b * 1
    }
    return score
}

func main() {
    b, err := os.ReadFile("day13.input")
    if err != nil {
        fmt.Print(err)
    }
    input := string(b)
    fmt.Println(part1(input))
    fmt.Println(part2(input))
}