package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func parseInput(input string) (map[string]map[string]bool) {
	graph := make(map[string]map[string]bool)
	for _, line := range strings.Split(input, "\n") {
		computer1, computer2 := strings.Split(line, "-")[0], strings.Split(line, "-")[1]
		if _, ok := graph[computer1]; !ok {
			graph[computer1] = make(map[string]bool)
		}
		graph[computer1][computer2] = true
		if _, ok := graph[computer2]; !ok {
			graph[computer2] = make(map[string]bool)
		}
		graph[computer2][computer1] = true
	}
	return graph
}

func part1(input string) int {
	graph := parseInput(input)

	var computers []string
	for computer, _ := range graph {
		computers = append(computers, computer)
	}

	count := 0
	for i := 0; i < len(computers); i++ {
		for j := i + 1; j < len(computers); j++ {
			if _, ok := graph[computers[i]][computers[j]]; ok {
				for k := j + 1; k < len(computers); k++ {
					if _, ok := graph[computers[j]][computers[k]]; ok {
						if _, ok := graph[computers[k]][computers[i]]; ok {
							if computers[i][0] == 't' || computers[j][0] == 't' || computers[k][0] == 't' {
								count += 1
							}
						}
					}
				}
			}
		}
	}
	return count
}

func part2(input string) string {
	graph := parseInput(input)

	var computers []string
	for computer, _ := range graph {
		computers = append(computers, computer)
	}

	largestClique := []int{}
	for i := 0; i < len(computers); i++ {
		for j := i + 1; j < len(computers); j++ {
			if _, ok := graph[computers[i]][computers[j]]; ok {
				for k := j + 1; k < len(computers); k++ {
					if _, ok := graph[computers[j]][computers[k]]; ok {
						if _, ok := graph[computers[k]][computers[i]]; ok {
							clique := []int{i, j, k}
							cliqueSet := make(map[int]bool)
							cliqueSet[i] = true
							cliqueSet[j] = true
							cliqueSet[k] = true
							for l := 0; l < len(computers); l++ {
								if !cliqueSet[l] {
									isConnected := true
									for _, c := range clique {
										if _, ok := graph[computers[l]][computers[c]]; !ok {
											isConnected = false
											break
										}
									}
									if isConnected {
										clique = append(clique, l)
										cliqueSet[l] = true
									}
								}
							}
							if len(clique) > len(largestClique) {
								largestClique = clique
							}
						}
					}
				}
			}
		}
	}
	computerNames := []string{}
	for _, i := range largestClique {
		computerNames = append(computerNames, computers[i])
	}
	sort.Strings(computerNames)
	return strings.Join(computerNames, ",")
}

func main() {
	b, err := os.ReadFile("day23.input")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}