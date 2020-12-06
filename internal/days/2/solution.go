package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/simskij/advent-of-code-2020/internal/data"
	"github.com/simskij/advent-of-code-2020/internal/types"
)

func GetSolution() types.Solution {
	lines := data.GetData("inputs/2", "\n")
	return types.Solution{
		Day: 2,
		Answers: types.Answers{
			A: solvePartOne(lines),
			B: solvePartTwo(lines),
		},
	}
}

func solvePartOne(lines []string) string {
	var validCount int
	for _, line := range lines {
		min, max, control, password := extractParts(line)
		count := strings.Count(password, control)
		if count >= min && count <= max {
			validCount++
		}
	}

	return fmt.Sprintf("%d", validCount)
}

func solvePartTwo(lines []string) string {
	var validCount int
	for _, line := range lines {
		min, max, control, password := extractParts(line)
		firstMatches := password[min-1:min] == control
		secondMatches := password[max-1:max] == control
		if (firstMatches || secondMatches) && !(firstMatches && secondMatches) {
			validCount++
		}
	}
	return fmt.Sprintf("%d", validCount)
}

func extractParts(line string) (int, int, string, string) {
	parts := strings.Split(line, ":")

	policy := strings.Trim(parts[0], " ")
	password := strings.Trim(parts[1], " ")

	policyParts := strings.Split(policy, " ")
	control := policyParts[1]
	minMaxParts := strings.Split(policyParts[0], "-")

	var min int
	var max int
	var err error

	if min, err = strconv.Atoi(minMaxParts[0]); err != nil {
		fmt.Println("Could not extract min occurance.")
		os.Exit(1)
	}

	if max, err = strconv.Atoi(minMaxParts[1]); err != nil {
		fmt.Println("Could not extract max occurance.")
		os.Exit(1)
	}

	return min, max, control, password
}
