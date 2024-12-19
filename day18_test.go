package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	b, _ := os.ReadFile("day18.example")
    input := string(b)

	step, error := part1(input, 7, 12)
	if error != nil {
		t.Errorf("Expected false, got %t", error)
	}
	if step != 22 {
		t.Errorf("Expected 22, got %d", step)
	}
}

func TestPart2(t *testing.T) {
	b, _ := os.ReadFile("day18.example")
    input := string(b)

	byte, error := part2(input, 7)
	if error != nil {
		t.Errorf("Expected false, got %t", error)
	}
	if byte != "6,1" {
		t.Errorf("Expected 6,1, got %s", byte)
	}
}