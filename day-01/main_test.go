package main

import "testing"

type TestCase struct {
	name     string
	file     string
	expected int
}

func TestSolvePart1(t *testing.T) {
	test := TestCase{
		"Test Part 1",
		"./part1-test.txt",
		11,
	}
	t.Run(test.name, func(t *testing.T) {
		solution := SolvePart1(test.file)
		if solution != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, solution)
		}
	})
}

func TestSolvePart2(t *testing.T) {
	test := TestCase{
		"Test Part 2",
		"./part2-test.txt",
		31,
	}
	t.Run(test.name, func(t *testing.T) {
		solution := SolvePart2(test.file)
		if solution != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, solution)
		}
	})
}
