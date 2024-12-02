package main

import "testing"

type TestCase struct {
	name     string
	file     string
	expected int
}

type ReportClassifiedTest struct {
	name     string
	report   string
	expected ReportSafety
}

func TestClassifyReport(t *testing.T) {
	tests := []ReportClassifiedTest{
		{
			"Report 1",
			"7 6 4 2 1",
			Safe,
		}, {
			"Report 2",
			"1 2 7 8 9",
			Unsafe,
		}, {
			"Report 3",
			"9 7 6 2 1",
			Unsafe,
		}, {
			"Report 4",
			"1 3 2 4 5",
			Unsafe,
		}, {
			"Report 5",
			"8 6 4 4 1",
			Unsafe,
		}, {
			"Report 6",
			"1 3 6 7 9",
			Safe,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res, _ := ClassifyReport(test.report); res != test.expected {
				t.Errorf("Test %v: expected %v, got %v",
					test.report,
					test.expected,
					res,
				)
			}
		})
	}
}

func TestClassifyReportByProblemDampener(t *testing.T) {
	tests := []ReportClassifiedTest{
		{
			"Report 1",
			"7 6 4 2 1",
			Safe,
		}, {
			"Report 2",
			"1 2 7 8 9",
			Unsafe,
		}, {
			"Report 3",
			"9 7 6 2 1",
			Unsafe,
		}, {
			"Report 4",
			"1 3 2 4 5",
			Safe,
		}, {
			"Report 5",
			"8 6 4 4 1",
			Safe,
		}, {
			"Report 6",
			"1 3 6 7 9",
			Safe,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := ClassifyReportByProblemDampener(test.report); res != test.expected {
				t.Errorf("Test %v: expected %v, got %v",
					test.report,
					test.expected,
					res,
				)
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	test := TestCase{
		"Test Part 1",
		"./part1-test.txt",
		2,
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

// func TestSolvePart2(t *testing.T) {
// 	test := TestCase{
// 		"Test Part 2",
// 		"./part2-test.txt",
// 		31,
// 	}
// 	t.Run(test.name, func(t *testing.T) {
// 	})
// }
