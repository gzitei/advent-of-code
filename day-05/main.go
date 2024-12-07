package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Rule struct {
	page, before int
}

var listRules = []Rule{}

func ReadFile(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file %v: %v\n", file, err)
		return "", err
	}
	txt := strings.TrimSpace(string(content))
	return txt, nil
}

func ParseInput(input string) ([]string, []string) {
	lst := strings.Split(input, "\n")
	sep := slices.Index(lst, "")
	return lst[:sep], lst[sep+1:]
}

func IsImpressionOrderValid(order map[int]int) bool {
	result := true

	for _, rule := range listRules {
		if idx, ok := order[rule.page]; ok {
			beforeIdx, beforeOk := order[rule.before]
			if beforeOk && beforeIdx > idx {
				result = false
				break
			}
		}
	}

	return result
}

func SolvePart1(file string) int {
	var solution int

	content, _ := ReadFile(file)
	rules, impressions := ParseInput(content)

	for _, rule := range rules {
		newRule := Rule{}
		fmt.Sscanf(rule, "%d|%d", &newRule.before, &newRule.page)
		listRules = append(listRules, newRule)
	}

	for _, imp := range impressions {
		pageOrder := map[int]int{}
		var n int
		pages := strings.Split(imp, ",")
		for i, page := range pages {
			var p int
			fmt.Sscanf(page, "%d", &p)
			pageOrder[p] = i
			if i == int(len(pages)/2) {
				n = p
			}
		}
		if IsImpressionOrderValid(pageOrder) {
			solution += n
		}
	}

	return solution
}

func MakeImpressionOrderValid(order map[int]int) []int {
	arr := make([]int, len(order))
	result := IsImpressionOrderValid(order)
	for !result {
		for _, rule := range listRules {
			if idx, ok := order[rule.page]; ok {
				beforeIdx, beforeOk := order[rule.before]
				if beforeOk && beforeIdx > idx {
					order[rule.page], order[rule.before] = beforeIdx, idx
					result = IsImpressionOrderValid(order)
					break
				}
			}
		}
		for k, v := range order {
			arr[v] = k
		}
	}
	return arr
}

func SolvePart2(file string) int {
	var solution int

	content, _ := ReadFile(file)
	rules, impressions := ParseInput(content)

	for _, rule := range rules {
		newRule := Rule{}
		fmt.Sscanf(rule, "%d|%d", &newRule.before, &newRule.page)
		listRules = append(listRules, newRule)
	}

	for _, imp := range impressions {
		pageOrder := map[int]int{}
		pages := strings.Split(imp, ",")
		for i, page := range pages {
			var p int
			fmt.Sscanf(page, "%d", &p)
			pageOrder[p] = i
		}
		arr := MakeImpressionOrderValid(pageOrder)

		if len(arr) > 0 {
			fmt.Println(arr)
			pos := int(math.Floor(float64(len(arr) / 2)))
			fmt.Println(arr, pos, arr[pos])
			solution += arr[pos]
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
