package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("day4.example")
    input := string(b)

	if part1(input) != 18 {
		t.Errorf("Expected 18, got %d", part1(input))
	}
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("day4.example")
    input := string(b)

	if part2(input) != 9 {
		t.Errorf("Expected 9, got %d", part2(input))
	}
}