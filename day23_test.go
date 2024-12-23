package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("day23.example")
    input := string(b)

	if part1(input) != 7 {
		t.Errorf("Expected 7, got %d", part1(input))
	}
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("day23.example")
    input := string(b)

	if part2(input) != "co,de,ka,ta" {
		t.Errorf("Expected co,de,ka,ta, got %s", part2(input))
	}
}