package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type ReportSafety bool

const (
	Safe   ReportSafety = true
	Unsafe              = false
)

func ClassifyReport(report string) (ReportSafety, int) {
	safety := Safe
	idx := -1
	data := strings.Split(report, " ")
	var leftLevel, midLevel, rightLevel int
	for i := 2; i < len(data); i++ {
		fmt.Sscanf(data[i-2], "%d", &leftLevel)
		fmt.Sscanf(data[i-1], "%d", &midLevel)
		fmt.Sscanf(data[i], "%d", &rightLevel)

		d1, d2 := midLevel-leftLevel, rightLevel-midLevel

		if d1 == 0 || d2 == 0 {
			safety = Unsafe
			idx = i - 2
			break
		}

		if (d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0) {
			safety = Unsafe
			idx = i - 2
			break
		}

		if int(math.Abs(float64(d1))) < 1 || int(math.Abs(float64(d1))) > 3 {
			safety = Unsafe
			idx = i - 2
			break
		}

		if int(math.Abs(float64(d2))) < 1 || int(math.Abs(float64(d2))) > 3 {
			safety = Unsafe
			idx = i - 2
			break
		}
	}
	return safety, idx
}

func ClassifyReportByProblemDampener(report string) ReportSafety {
	safety, idx := ClassifyReport(report)
	if !safety {
		data := strings.Split(report, " ")
		for i := idx; i < len(data); i++ {
			newList := strings.Join(slices.Concat(data[:i], data[i+1:]), " ")
			safety, _ = ClassifyReport(newList)
			if safety {
				break
			}
		}
	}
	return safety
}

func ReadFile(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file %v: %v\n", file, err)
		return "", err
	}
	txt := strings.TrimSpace(string(content))
	return txt, nil
}

func SolvePart1(file string) int {
	var solution int

	txt, _ := ReadFile(file)

	chunks := strings.Split(txt, "\n")

	for _, chunk := range chunks {
		safety, _ := ClassifyReport(chunk)
		if safety {
			solution++
		}
	}

	return solution
}

func SolvePart2(file string) int {
	var solution int

	txt, _ := ReadFile(file)

	chunks := strings.Split(txt, "\n")

	for _, chunk := range chunks {
		safety := ClassifyReportByProblemDampener(chunk)
		if safety {
			solution++
		}
	}

	return solution
}

func main() {
	var part int
	fmt.Println("Which part you want to solve? [1 or 2]")
	if _, err := fmt.Scanf("%d", &part); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	switch part {
	case 1:
		{
			file := "./part1-puzzle.txt"
			solution := SolvePart1(file)
			fmt.Printf("Solution: %v\n", solution)
		}
	case 2:
		{
			file := "./part2-puzzle.txt"
			solution := SolvePart2(file)
			fmt.Printf("Solution: %v\n", solution)
		}
	default:
		{
			fmt.Printf("%v: invalid option.\n", part)
		}
	}
}
