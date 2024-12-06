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
		18,
	}
	t.Run(test.name, func(t *testing.T) {
		t.Run(test.name, func(t *testing.T) {
			if sol := SolvePart1(test.file); sol != test.expected {
				t.Errorf("Test %v: expected %v, got %v",
					test.name,
					test.expected,
					sol,
				)
			}
		})
	})
}

func TestSolvePart2(t *testing.T) {
	test := TestCase{
		"Test Part 2",
		"./part2-test.txt",
		9,
	}
	t.Run(test.name, func(t *testing.T) {
		if sol := SolvePart2(test.file); sol != test.expected {
			t.Errorf("Test %v: expected %v, got %v",
				test.name,
				test.expected,
				sol,
			)
		}
	})
}
