package main

import (
	"fmt"
	"os"
	"strings"
)

type Parser struct {
	start       int
	end         int
	parenthesis int
	index       int
	arr         [2]string
}

func newParser() Parser {
	return Parser{
		start: -1, end: -1, parenthesis: 0, index: 0, arr: [2]string{},
	}
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
	s, _ := ReadFile(file)
	solution := 0
	p := newParser()

	for i := 0; i < len(s); i++ {
		if s[i] == 'm' && p.start == -1 {
			p.start = i
			p.end = i
			continue
		}

		if p.start == p.end && s[i] == 'u' {
			p.end++
			continue
		}

		if p.end == p.start+1 && s[i] == 'l' {
			p.end++
			continue
		}

		if p.end == p.start+2 && s[i] == '(' {
			p.end++
			p.parenthesis = 1
			continue
		}

		if p.parenthesis > 0 {
			if s[i] == ')' {
				p.parenthesis--
				if len(p.arr[0]) > 0 && len(p.arr[1]) > 0 {
					var n1, n2 int
					fmt.Sscanf(p.arr[0], "%d", &n1)
					fmt.Sscanf(p.arr[1], "%d", &n2)
					solution += n1 * n2
				}
			} else if s[i] == ',' {
				p.index++
				if len(p.arr[0]) > 0 {
					continue
				}
			} else if s[i]-'0' >= 0 && s[i]-'0' <= 9 {
				if len(p.arr[p.index]) < 3 {
					p.arr[p.index] += string(s[i])
					continue
				}
			}
		}
		p = newParser()
	}
	return solution
}

func SolvePart2(file string) int {
	s, _ := ReadFile(file)
	do := true
	solution := 0
	p := newParser()

	for i := 0; i < len(s); i++ {
		if do {
			if (s[i] == 'm' || s[i] == 'd') && p.start == -1 {
				p.start = i
				p.end = i
				continue
			}

			if p.start < 0 {
				continue
			}

			if s[p.start] == 'm' {
				if p.start > 0 && p.start == p.end && s[i] == 'u' {
					p.end++
					continue
				}

				if p.end == p.start+1 && s[i] == 'l' {
					p.end++
					continue
				}

				if p.end == p.start+2 && s[i] == '(' {
					p.end++
					p.parenthesis = 1
					continue
				}

				if p.parenthesis > 0 {
					if s[i] == ')' {
						p.parenthesis--
						if len(p.arr[0]) > 0 && len(p.arr[1]) > 0 {
							var n1, n2 int
							fmt.Sscanf(p.arr[0], "%d", &n1)
							fmt.Sscanf(p.arr[1], "%d", &n2)
							solution += n1 * n2
						}
					} else if s[i] == ',' {
						p.index++
						if len(p.arr[0]) > 0 {
							continue
						}
					} else if s[i]-'0' >= 0 && s[i]-'0' <= 9 {
						if len(p.arr[p.index]) < 3 {
							p.arr[p.index] += string(s[i])
							continue
						}
					}
				}
			} else {
				if p.start > 0 && p.end == p.start && s[i] == 'o' {
					p.end++
					continue
				}

				if p.end == p.start+1 && s[i] == 'n' {
					p.end++
					continue
				}

				if p.end == p.start+2 && string(s[i]) == "'" {
					p.end++
					continue
				}

				if p.end == p.start+3 && s[i] == 't' {
					p.end++
					continue
				}

				if p.end == p.start+4 && s[i] == '(' {
					p.end++
					continue
				}

				if p.end == p.start+5 && s[i] == ')' {
					do = !do
				}
			}

		} else {
			if s[i] == 'd' && p.start == -1 {
				p.start = i
				p.end = p.start
				continue
			}

			if p.start > 0 && p.end == p.start && s[i] == 'o' {
				p.end++
				continue
			}

			if p.end == p.start+1 && s[i] == '(' {
				p.end++
				continue
			}

			if p.end == p.start+2 && s[i] == ')' {
				do = !do
			}

		}
		p = newParser()
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
			file := "./puzzle.txt"
			solution := SolvePart1(file)
			fmt.Printf("Solution: %v\n", solution)
		}
	case 2:
		{
			file := "./puzzle.txt"
			solution := SolvePart2(file)
			fmt.Printf("Solution: %v\n", solution)
		}
	default:
		{
			fmt.Printf("%v: invalid option.\n", part)
		}
	}
}
