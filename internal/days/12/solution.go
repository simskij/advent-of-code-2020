package day12

import (
	"fmt"
	"github.com/simskij/advent-of-code-2020/internal/data"
	"github.com/simskij/advent-of-code-2020/internal/types"
	"math"
	"strconv"
)

func GetSolution() types.Solution {
	lines := data.GetData("inputs/12", "\n")
	a := getAnswerForA(lines)
	b := getAnswerForB(lines)
	return types.Solution{
		Day: 12,
		Answers: types.Answers{
			A: fmt.Sprintf("%d", a),
			B: fmt.Sprintf("%d", b),
		},
	}
}

func abs (val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func getAnswerForA(lines []string) int {
	wx, wy := 1, 0
	x, y := 0, 0
	for _, line := range lines {
		move := line[0:1]
		count, _ := strconv.Atoi(line[1:])
		switch move {
		case "E":
			x += count
			break
		case "W":
			x -= count
			break
		case "N":
			y += count
			break
		case "S":
			y -= count
			break
		case "R":
			wx, wy = Rotate(wx, wy, 360 - count)
			break
		case "L":
			wx, wy = Rotate(wx, wy, count)
			break
		case "F":
			x += wx * count
			y += wy * count
		}
	}
	return abs(x) + abs(y)
}

func getAnswerForB(lines []string) int {
	wx, wy := 10, 1
	x, y := 0, 0

	for _, line := range lines {
		move := line[0:1]
		count, _ := strconv.Atoi(line[1:])
		switch move {

		// Move WP
		case "E":
			wx += count
			break
		case "W":
			wx -= count
			break
		case "N":
			wy += count
			break
		case "S":
			wy -= count
			break

		// Rotate
		case "L":
			wx, wy = Rotate(wx, wy, count)
			break
		case "R":
			wx, wy = Rotate(wx, wy, -count)
			break
		case "F":
			x += wx * count
			y += wy * count
			break
		}
	}
	return abs(x) + abs(y)
}

func Rotate(x, y, theta int) (int, int) {
	t := float64(theta)
	cos := int(math.Round(math.Cos(t * math.Pi / 180)))
	sin := int(math.Round(math.Sin(t * math.Pi / 180)))

	ox := cos * x - sin * y
	oy := sin * x + cos * y

	return ox, oy
}