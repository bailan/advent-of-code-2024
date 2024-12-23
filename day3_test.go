package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("day3.example1")
    input := string(b)

	if part1(input) != 161 {
		t.Errorf("Expected 161, got %d", part1(input))
	}
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("day3.example2")
    input := string(b)

	if part2(input) != 48 {
		t.Errorf("Expected 48, got %d", part2(input))
	}
}