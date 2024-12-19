package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("day19.example")
    input := string(b)

	if part1(input) != 6 {
		t.Errorf("Expected 6, got %d", part1(input))
	}
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("day19.example")
    input := string(b)

	if part2(input) != 16 {
		t.Errorf("Expected 16, got %d", part2(input))
	}
}