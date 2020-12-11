package day7

import (
	"fmt"
	"github.com/simskij/advent-of-code-2020/internal/data"
	"github.com/simskij/advent-of-code-2020/internal/types"
	"os"
	"strconv"
)

func GetSolution() types.Solution {
	lines := data.GetData("inputs/9", "\n")
	numbers := transformToNumbers(lines)
	a := findWeakNumber(numbers)
	b := findContiguousSet(numbers, a)
	return types.Solution{
		Day: 9,
		Answers: types.Answers{
			A: fmt.Sprintf("%d", a),
			B: fmt.Sprintf("%d", b),
		},
	}
}

func transformToNumbers(lines []string) []int {
	var numbers []int
	for i, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Invalid input at line %d. Unable to convert to integer.", i)
			os.Exit(1)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func findContiguousSet(numbers []int, weak int) int {
	i, j := findIndexOfTerms(numbers, weak)
	sub := numbers[i:i+j]
	min, max := weak, 0

	for _, num := range sub {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min + max
}

func findIndexOfTerms(numbers []int, weak int) (int, int) {
	for i, a := range numbers {
		sum := a
		for j, b := range numbers[i+1:] {
			sum += b
			if sum == weak {
				return i, j
			}
		}
	}
	return 0, 0
}

func findWeakNumber(numbers []int) int {
	for i, n := range numbers {
		if i < 25 {
			continue
		}

		prevNums := getPreviousRange(i, numbers)
		matched := false

		for ia, a := range prevNums {
			for ib, b := range prevNums {
				if ia == ib {
					continue
				}
				if a+b == n {
					matched = true
					break
				}
			}
			if matched {
				break
			}
		}
		if !matched {
			return n
		}
	}
	return 0
}

func getPreviousRange(i int, numbers []int) []int {
	var prevNums []int
	if i-25 < 0 {
		prevNums = numbers[0:i]
	} else {
		prevNums = numbers[i-25 : i]
	}
	return prevNums
}