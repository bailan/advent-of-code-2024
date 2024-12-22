package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("day2.example")
    input := string(b)

	if part1(input) != 2 {
		t.Errorf("Expected 2, got %d", part1(input))
	}
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("day2.example")
    input := string(b)

	if part2(input) != 4 {
		t.Errorf("Expected 4, got %d", part2(input))
	}
}