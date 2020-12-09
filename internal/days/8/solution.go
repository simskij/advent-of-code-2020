package day7

import (
	"fmt"
	"github.com/simskij/advent-of-code-2020/internal/data"
	"github.com/simskij/advent-of-code-2020/internal/types"
	"strconv"
	"strings"
)

func GetSolution() types.Solution {
	lines := data.GetData("inputs/8", "\n")
	a, _ := runInstructions(lines, false)
	b := findCorrectInstructions(lines)
	return types.Solution{
		Day: 8,
		Answers: types.Answers{
			A: fmt.Sprintf("%d", a),
			B: fmt.Sprintf("%d", b),
		},
	}
}

func shouldContinue(executed map[int]bool, current int, lines []string, fix bool) bool {
	if fix == true {
		return current >= 0 && current < len(lines) && !executed[current]
	}
	return !executed[current]
}

func findCorrectInstructions(lines []string) int {
	old := ""
	for i, line := range lines {
		old = line
		if strings.Contains(line, "jmp") {
			lines[i] = strings.Replace(lines[i], "jmp", "nop", 1)
		} else if strings.Contains(line, "jmp") {
			lines[i] = strings.Replace(lines[i], "jmp", "nop", 1)
		}
		if sum, inst := runInstructions(lines, true); inst == len(lines) {
			return sum
		}
		lines[i] = old
	}

	return 0
}


func runInstructions(lines []string, fix bool) (int, int) {
	executed := map[int]bool{}
	current := 0
	sum := 0

	for shouldContinue(executed, current, lines, fix) {
		executed[current] = true
		instruction := lines[current]
		parts := strings.Split(instruction, " ")
		left := parts[0]
		right := parts[1]

		size, err := strconv.Atoi(right)

		if err != nil {
			fmt.Printf("Could not parse argument to %s instruction on line %d\n", current, left)
		}

		if left == "acc" {
			sum += size
		}

		if left == "jmp" {
			current += size
			continue
		}

		current++
	}
	return sum, current
}