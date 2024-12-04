package main

import "testing"

func TestSolvePart1(t *testing.T) {
	test := struct {
		name     string
		file     string
		expected int
	}{
		name:     "Part 1",
		file:     "test1.txt",
		expected: 161,
	}
	t.Run(test.name, func(t *testing.T) {
		if sol := SolvePart1(test.file); test.expected != sol {
			t.Errorf("Test %v: expected %v, got %v",
				test.name,
				test.expected,
				sol,
			)
		}
	})
}

func TestSolvePart2(t *testing.T) {
	test := struct {
		name     string
		file     string
		expected int
	}{
		name:     "Part 2",
		file:     "test2.txt",
		expected: 48,
	}
	t.Run(test.name, func(t *testing.T) {
		if sol := SolvePart2(test.file); test.expected != sol {
			t.Errorf("Test %v: expected %v, got %v",
				test.name,
				test.expected,
				sol,
			)
		}
	})
}
