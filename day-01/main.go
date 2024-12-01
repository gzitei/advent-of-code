package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Search(n int, list []int) int {
	idx, left, right := -1, 0, len(list)-1
	for left <= right {
		mid := left + (right-left)/2
		if list[mid] == n {
			idx = mid
			break
		} else if list[mid] < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

func CheckLeft(n int, list []int) int {
	occur := 0
	for i := range list {
		if list[len(list)-(1+i)] == n {
			occur++
		} else {
			break
		}
	}
	return occur
}

func CheckRight(n int, list []int) int {
	occur := 0
	for i := range list {
		if list[i] == n {
			occur++
		} else {
			break
		}
	}
	return occur
}

func GetOcurrences(n int, list []int, idx int) int {
	var occur int

	if idx == -1 {
		occur = 0
	} else {
		occur = 1
		if idx-1 >= 0 {
			occur += CheckLeft(n, list[0:idx])
		}
		if idx+1 <= len(list) {
			occur += CheckRight(n, list[idx+1:])
		}
	}

	return occur
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

func ParseContent(content string) ([]int, []int, error) {
	rows := strings.Split(content, "\n")
	arr1, arr2 := make([]int, 0, len(rows)), make([]int, 0, len(rows))

	for _, row := range rows {
		values := strings.Split(row, "  ")

		v1, err := strconv.Atoi(strings.TrimSpace(values[0]))
		if err != nil {
			fmt.Printf("Error parsing value %v: %v\n", values[0], err)
			return nil, nil, err
		}
		arr1 = append(arr1, v1)

		v2, err := strconv.Atoi(strings.TrimSpace(values[1]))
		if err != nil {
			fmt.Printf("Error parsing value %v: %v\n", values[1], err)
			return nil, nil, err
		}
		arr2 = append(arr2, v2)
	}

	return arr1, arr2, nil
}

func DistanceBetween(v1, v2 int) int {
	return int(math.Abs(float64(v1 - v2)))
}

func GetDistances(arr1, arr2 []int) int {
	var dist int
	for i, v1 := range arr1 {
		v2 := arr2[i]
		dist += DistanceBetween(v1, v2)
	}
	return dist
}

func GetSimilarity(arr1, arr2 []int) int {
	similarity := 0
	for _, v := range arr1 {
		idx := Search(v, arr2)
		occurrences := GetOcurrences(v, arr2, idx)
		similarity += (v * occurrences)
	}
	return similarity
}

func SolvePart1(file string) int {
	var solution int

	content, err := ReadFile(file)
	if err != nil {
		os.Exit(1)
	}

	arr1, arr2, err := ParseContent(content)
	if err != nil {
		os.Exit(1)
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	solution = GetDistances(arr1, arr2)

	return solution
}

func SolvePart2(file string) int {
	var solution int

	content, err := ReadFile(file)
	if err != nil {
		os.Exit(1)
	}

	arr1, arr2, err := ParseContent(content)
	if err != nil {
		os.Exit(1)
	}

	slices.Sort(arr2)

	solution = GetSimilarity(arr1, arr2)

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
