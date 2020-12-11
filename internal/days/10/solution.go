package day10

import (
	"fmt"
	"github.com/simskij/advent-of-code-2020/internal/data"
	"github.com/simskij/advent-of-code-2020/internal/types"
	"os"
	"sort"
	"strconv"
)

func GetSolution() types.Solution {
	lines := data.GetData("inputs/10", "\n")
	numbers := append(transformToNumbers(lines), 0)
	sort.Ints(numbers)

	a := createSequence(numbers)
	b := findPermutations(numbers, map[string]int{})

	return types.Solution{
		Day: 10,
		Answers: types.Answers{
			A: fmt.Sprintf("%d", a),
			B: fmt.Sprintf("%d", b),
		},
	}
}

func createSequence(numbers []int) int {
	diffs := map[int]int {}
	for i, num := range numbers {
		if i == len(numbers) -1 || num + 3 == numbers[i+1] {
			diffs[3]++
			continue
		} else
		if num + 1 == numbers[i+1] {
			diffs[1]++
		}
	}
	return diffs[1] * (diffs[3])
}

func findPermutations(numbers []int, mem map[string]int) int {
	key := fmt.Sprint(numbers)

	if mem[key] == 0 {
		result := followSequence(numbers, mem)
		mem[key] = result
	}
	return mem[key]
}

func followSequence(numbers []int, mem map[string]int) int {
	result := 1
	for i := 1; i < len(numbers)-1; i++ {
		if numbers[i+1]-numbers[i-1] > 3 {
			continue
		}
		next := append([]int{numbers[i-1]}, numbers[i+1:]...)
		result += findPermutations(next, mem)
	}
	return result
}

func transformToNumbers(lines []string) []int {
		numbers := []int{}

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