package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("day13.example")
    input := string(b)

	if part1(input) != 480 {
		t.Errorf("Expected 480, got %d", part1(input))
	}
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("day13.example")
    input := string(b)

	if part2(input) != 459236326669 {
		t.Errorf("Expected 459236326669, got %d", part2(input))
	}
}