package day3

import (
	"fmt"

	"github.com/simskij/advent-of-code-2020/internal/data"
	"github.com/simskij/advent-of-code-2020/internal/types"
)

func GetSolution() types.Solution {
	lines := data.GetData("inputs/3", "\n")

	one := solvePartOne(lines)
	two := solvePartTwo(lines)
	// ...
	return types.Solution{
		Day: 3,
		Answers: types.Answers{
			A: fmt.Sprintf("%d", one),
			B: fmt.Sprintf("%d", two),
		},
	}
}

func solvePartOne(lines []string) int {
	return getTreesForSlope(lines, 3, 1)
}

func solvePartTwo(lines []string) int {

	return getTreesForSlope(lines, 1, 1) *
		getTreesForSlope(lines, 3, 1) *
		getTreesForSlope(lines, 5, 1) *
		getTreesForSlope(lines, 7, 1) *
		getTreesForSlope(lines, 1, 2)
}

func getTreesForSlope(lines []string, xIncrement int, yIncrement int) int {
	var trees int
	x := 0
	for y := yIncrement; y < len(lines); y += yIncrement {
		line := lines[y]
		x = (x + xIncrement) % len(line)
		content := line[x : x+1]

		if content == "." {
		} else {
			trees++
		}
	}
	return trees
}
