package day6

import (
	"fmt"
	"strings"

	"github.com/simskij/advent-of-code-2020/internal/data"
	"github.com/simskij/advent-of-code-2020/internal/types"
)

func GetSolution() types.Solution {
	lines := data.GetData("inputs/6", "\n\n")
	consensus, total := getCounts(lines)

	return types.Solution{
		Day: 6,
		Answers: types.Answers{
			A: fmt.Sprintf("%d", total),
			B: fmt.Sprintf("%d", consensus),
		},
	}
}

func getCounts(lines []string) (int, int) {
	consensus, total := 0, 0
	for _, group := range lines {
		persons := strings.Split(group, "\n")
		counts := countForGroup(persons)
		for _, count := range counts {
			if count == len(persons) {
				consensus++
			}
		}
		total += len(counts)

	}
	return consensus, total
}

func countForGroup(group []string) map[string]int {
	counts := map[string]int{}
	for _, person := range group {
		for _, question := range strings.Split(person, "") {
			counts[question]++
		}
	}
	return counts
}
