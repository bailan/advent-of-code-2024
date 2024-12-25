package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("day25.example")
    input := string(b)

	if part1(input) != 3 {
		t.Errorf("Expected 3, got %d", part1(input))
	}
}