package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) (map[string]map[string]bool, [][]string) {
	ordersString, updatesString := strings.Split(input, "\n\n")[0], strings.Split(input, "\n\n")[1]

	graph := make(map[string]map[string]bool)
	for _, line := range strings.Split(ordersString, "\n") {
		numbers := strings.Split(line, "|")
		if _, ok := graph[numbers[0]]; !ok {
			graph[numbers[0]] = make(map[string]bool)
		}
		graph[numbers[0]][numbers[1]] = true
	}

	var updates [][]string
	for _, line := range strings.Split(updatesString, "\n") {
		updates = append(updates, strings.Split(line, ","))
	}
	return graph, updates
}

func part1(input string) int {
	graph, updates := parseInput(input)

	result := 0
	for _, update := range updates {
		validUpdate := true
		for i := 0; i < len(update) - 1; i++ {
			if graph[update[i + 1]][update[i]] {
				validUpdate = false
				break
			}
		}
		if validUpdate {
			middleNumber, _ := strconv.Atoi(update[len(update)/2])
			result += middleNumber
		}
	}
	return result
}

func part2(input string) int {
	graph, updates := parseInput(input)

	result := 0
	for _, update := range updates {
		updated := false
		for i := 0; i < len(update) - 1; i++ {
			j := i
			for j >= 0 && graph[update[j + 1]][update[j]] {
				update[j + 1], update[j] = update[j], update[j + 1]
				j -= 1
				updated = true
			}
		}
		if updated {
			middleNumber, _ := strconv.Atoi(update[len(update)/2])
			result += middleNumber
		}
	}
	return result
}


func main() {
	b, err := os.ReadFile("day5.input")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}