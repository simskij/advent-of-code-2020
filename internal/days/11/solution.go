package day11

import (
	"fmt"
	"github.com/simskij/advent-of-code-2020/internal/data"
	"github.com/simskij/advent-of-code-2020/internal/types"
	"reflect"
	"strings"
)

func GetSolution() types.Solution {
	lines := data.GetData("inputs/11", "\n")
	a := getAnswerForA(lines)
	b := getAnswerForB(lines)
	return types.Solution{
		Day: 11,
		Answers: types.Answers{
			A: fmt.Sprintf("%d", a),
			B: fmt.Sprintf("%d", b),
		},
	}
}

func getAnswerForB(lines []string) int {
	matrix := createMatrix(lines)
	matrix = runFarUntilStale(matrix)
	return countOccupied(matrix)
}

func getAnswerForA(lines []string) int {
	matrix := createMatrix(lines)
	matrix = runUntilStale(matrix)
	return countOccupied(matrix)
}

func createMatrix(lines []string) [][]string {
	var matrix [][]string
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}
	return matrix
}

func runUntilStale(matrix [][]string) [][]string {
	c := make([][]string, len(matrix))
	for i := range matrix {
		c[i] = append([]string{}, matrix[i]...)
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == "." {
				continue
			}
			adjacent := 0
			if i > 0 && matrix[i-1][j] == "#" {
				adjacent++
			}
			if i < len(matrix) - 1 && matrix[i+1][j] == "#" {
				adjacent++
			}
			if j > 0 && matrix[i][j-1] == "#" {
				adjacent++
			}
			if j < len(matrix[i]) - 1 && matrix[i][j+1] == "#" {
				adjacent++
			}
			if i < len(matrix) - 1 && j < len(matrix[i]) - 1 && matrix[i+1][j+1] == "#" {
				adjacent++
			}
			if i < len(matrix) - 1 && j > 0 && matrix[i+1][j-1] == "#" {
				adjacent++
			}

			if i > 0 && j < len(matrix[i]) - 1 && matrix[i-1][j+1] == "#" {
				adjacent++
			}
			if i > 0 && j > 0 && matrix[i-1][j-1] == "#" {
				adjacent++
			}

			if adjacent >= 4 {
				c[i][j] = "L"
			} else if adjacent == 0 {
				c[i][j] = "#"
			}
		}
	}

	if reflect.DeepEqual(c, matrix) {
		return c
	}

	return runUntilStale(c)
}


func runFarUntilStale(matrix [][]string) [][]string {
	c := make([][]string, len(matrix))
	for i := range matrix {
		c[i] = append([]string{}, matrix[i]...)
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == "." {
				continue
			}
			adjacent := 0


			// Up
			for d := 1; i-d >= 0; d++ {
				if matrix[i-d][j] == "L" {
					break
				}
				if matrix[i-d][j] == "#" {
					adjacent++
					break
				}
			}
			// Down
			for d := 1; i + d < len(matrix); d++ {
				if matrix[i+d][j] == "L" {
					break
				}
				if matrix[i+d][j] == "#" {
					adjacent++
					break
				}
			}

			// Left
			for d := 1; j - d >= 0; d++ {
				if matrix[i][j-d] == "L" {
					break
				}
				if matrix[i][j-d] == "#" {
					adjacent++
					break
				}
			}

			// Right
			for d := 1; j + d < len(matrix[i]); d++ {
				if matrix[i][j+d] == "L" {
					break
				}
				if matrix[i][j+d] == "#" {
					adjacent++
					break
				}
			}

			// Down/Right
			for r := 1; i+r < len(matrix) && j+r < len(matrix[i]); r++ {
				if matrix[i+r][j+r] == "L" {
					break
				}
				if matrix[i+r][j+r] == "#" {
					adjacent++
					break
				}
			}

			// Up/Right
			for r := 1; i-r >= 0 && j+r < len(matrix[i]); r++ {
				if matrix[i-r][j+r] == "L" {
					break
				}
				if matrix[i-r][j+r] == "#" {
					adjacent++
					break
				}
			}
			// Up/Left
			for r := 1; i+r < len(matrix) && j-r >= 0; r++ {
				if matrix[i+r][j-r] == "L" {
					break
				}
				if matrix[i+r][j-r] == "#" {
					adjacent++
					break
				}
			}

			// Down/Left
			for r := 1; i-r >= 0 && j-r >= 0; r++ {
				if matrix[i-r][j-r] == "L" {
					break
				}
				if matrix[i-r][j-r] == "#" {
					adjacent++
					break
				}
			}

			if adjacent >= 5 {
				c[i][j] = "L"
			} else if adjacent == 0 {
				c[i][j] = "#"
			}
		}
	}
	if reflect.DeepEqual(c, matrix) {
		return c
	}

	return runFarUntilStale(c)
}

func countOccupied(matrix [][]string) int {
	count := 0
	for _, row := range matrix {
		for _, col := range row {
			if col == "#" {
				count++
			}
		}
	}
	return count
}