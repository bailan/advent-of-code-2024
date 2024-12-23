package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("day5.example")
    input := string(b)

	if part1(input) != 143 {
		t.Errorf("Expected 143, got %d", part1(input))
	}
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("day5.example")
    input := string(b)

	if part2(input) != 123 {
		t.Errorf("Expected 123, got %d", part2(input))
	}
}