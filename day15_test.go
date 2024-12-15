package main

import (
	"os"
	"testing"
)

func TestPart1Case1(t *testing.T) {
	b, _ := os.ReadFile("day15.example1")
    input := string(b)

	if part1(input) != 2028 {
		t.Errorf("Expected 2028, got %d", part1(input))
	}
}

func TestPart1Case2(t *testing.T) {
	b, _ := os.ReadFile("day15.example2")
    input := string(b)

	if part1(input) != 10092 {
		t.Errorf("Expected 10092, got %d", part1(input))
	}
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("day15.example2")
    input := string(b)

	if part2(input) != 9021 {
		t.Errorf("Expected 9021, got %d", part2(input))
	}
}