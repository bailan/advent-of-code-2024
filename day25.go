package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) ([][5]int, [][5]int) {
	schematics := strings.Split(input, "\n\n")
	var keys [][5]int
	var pins [][5]int
	for _, scheme := range schematics {
		rows := strings.Split(scheme, "\n")
		if rows[0] == "#####" {
			key := [5]int{}
			for i := 0; i < 5; i++ {
				for j := 0; j < 6; j++ {
					if rows[j + 1][i] == '#' {
						key[i] = j + 1
					} else {
						break
					}
				}
			}
			keys = append(keys, key)
		} else {
			pin := [5]int{}
			for i := 0; i < 5; i++ {
				for j := 0; j < 6; j++ {
					if rows[5 - j][i] == '#' {
						pin[i] = j + 1
					} else {
						break
					}
				}
			}
			pins = append(pins, pin)
		}
	}
	return keys, pins
}

func part1(input string) int {
	keys, pins := parseInput(input)
	count := 0
	for _, key := range keys {
		for _, pin := range pins {
			match := true
			for i := 0; i < 5; i++ {
				if key[i] + pin[i] > 5 {
					match = false
					break
				}
			}
			if match {
				count++
			}
		}
	}
	return count
}

func main() {
	b, err := os.ReadFile("day25.input")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(part1(input))
}