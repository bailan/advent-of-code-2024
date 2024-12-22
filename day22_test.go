package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("day22.example1")
    input := string(b)

	if part1(input) != 37327623 {
		t.Errorf("Expected 37327623, got %d", part1(input))
	}
}

func TestPart2Case1(t *testing.T) {
	b, _ := os.ReadFile("day22.example2")
    input := string(b)

	if part2(input, 10) != 6 {
		t.Errorf("Expected 6, got %d", part2(input, 10))
	}
}

func TestPart2Case2(t *testing.T) {
	b, _ := os.ReadFile("day22.example3")
    input := string(b)

	if part2(input, 2000) != 23 {
		t.Errorf("Expected 23, got %d", part2(input, 2000))
	}
}