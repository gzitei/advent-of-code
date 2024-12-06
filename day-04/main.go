package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file %v: %v\n", file, err)
		return "", err
	}
	txt := strings.TrimSpace(string(content))
	return txt, nil
}

func CountFront(x, y int, matrix []string) int {
	if y+4 > len(matrix[x]) {
		return 0
	}
	s := matrix[x][y : y+4]
	if s == "XMAS" || s == "SAMX" {
		return 1
	}
	return 0
}

func CountBack(x, y int, matrix []string) int {
	if y-3 < 0 {
		return 0
	}
	s := matrix[x][y-3 : y+1]
	if s == "XMAS" || s == "SAMX" {
		return 1
	}
	return 0
}

func CountUp(x, y int, matrix []string) int {
	if x-3 < 0 {
		return 0
	}
	s := string([]byte{matrix[x][y], matrix[x-1][y], matrix[x-2][y], matrix[x-3][y]})
	if s == "XMAS" || s == "SAMX" {
		return 1
	}
	return 0
}

func CountDown(x, y int, matrix []string) int {
	if x+3 >= len(matrix) {
		return 0
	}
	s := string([]byte{matrix[x][y], matrix[x+1][y], matrix[x+2][y], matrix[x+3][y]})
	if s == "XMAS" || s == "SAMX" {
		return 1
	}
	return 0
}

func CountDiagonal45(x, y int, matrix []string) int {
	if x+3 >= len(matrix) || y-3 < 0 {
		return 0
	}
	s := string([]byte{matrix[x][y], matrix[x+1][y-1], matrix[x+2][y-2], matrix[x+3][y-3]})
	if s == "XMAS" || s == "SAMX" {
		return 1
	}
	return 0
}

func CountDiagonal135(x, y int, matrix []string) int {
	if x+3 >= len(matrix) || y+3 >= len(matrix[x]) {
		return 0
	}
	s := string([]byte{matrix[x][y], matrix[x+1][y+1], matrix[x+2][y+2], matrix[x+3][y+3]})
	if s == "XMAS" || s == "SAMX" {
		return 1
	}
	return 0
}

func CountDiagonal225(x, y int, matrix []string) int {
	if x-3 < 0 || y-3 < 0 {
		return 0
	}
	s := string([]byte{matrix[x][y], matrix[x-1][y-1], matrix[x-2][y-2], matrix[x-3][y-3]})
	if s == "XMAS" || s == "SAMX" {
		return 1
	}
	return 0
}

func CountDiagonal315(x, y int, matrix []string) int {
	if x-3 < 0 || y+3 >= len(matrix[x]) {
		return 0
	}
	s := string([]byte{matrix[x][y], matrix[x-1][y+1], matrix[x-2][y+2], matrix[x-3][y+3]})
	if s == "XMAS" || s == "SAMX" {
		return 1
	}
	return 0
}

func SolvePart1(file string) int {
	var solution int
	s, _ := ReadFile(file)

	rows := strings.Split(s, "\n")

	for x := 0; x < len(rows); x++ {
		for y := 0; y < len(rows[x]); y++ {
			if rows[x][y] == 'X' {
				solution += CountFront(x, y, rows)
				solution += CountBack(x, y, rows)
				solution += CountUp(x, y, rows)
				solution += CountDown(x, y, rows)
				solution += CountDiagonal45(x, y, rows)
				solution += CountDiagonal135(x, y, rows)
				solution += CountDiagonal225(x, y, rows)
				solution += CountDiagonal315(x, y, rows)
			}
		}
	}
	return solution
}

func CheckAroung(x, y int, rows []string) bool {
	d1 := string([]byte{rows[x-1][y-1], rows[x][y], rows[x+1][y+1]})
	d2 := string([]byte{rows[x+1][y-1], rows[x][y], rows[x-1][y+1]})
	return (d1 == "SAM" || d1 == "MAS") && (d2 == "SAM" || d2 == "MAS")
}

func SolvePart2(file string) int {
	var solution int
	s, _ := ReadFile(file)

	rows := strings.Split(s, "\n")

	for x := 1; x < len(rows)-1; x++ {
		for y := 1; y < len(rows[x])-1; y++ {
			if rows[x][y] == 'A' {
				if CheckAroung(x, y, rows) {
					solution++
				}
			}
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
