package day5

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/simskij/advent-of-code-2020/internal/data"
	"github.com/simskij/advent-of-code-2020/internal/types"
)

const SEATS = 128 * 8

func GetSolution() types.Solution {
	lines := data.GetData("inputs/5", "\n")
	return types.Solution{
		Day:       5,
		Variation: "partitioning",
		Answers: types.Answers{
			A: fmt.Sprintf("%d", resolveSeats(lines)),
			B: fmt.Sprintf("%d", getFreeSeat(lines)),
		},
	}
}

func traverse(min, max, needle int, direction string) (nextMin int, nextMax int, nextNeedle int) {
	if direction == "B" || direction == "R" {
		return needle, max, max - ((max - needle) / 2)
	} else if direction == "F" || direction == "L" {
		return min, needle, needle - ((needle - min) / 2)
	}

	fmt.Printf("%s is not a valid direction", direction)
	os.Exit(1)
	return 0, 0, 0
}

func resolveSeats(lines []string) int {
	highest := 0
	for _, line := range lines {
		seatId := resolveSeat(line)
		if seatId > highest {
			highest = seatId
		}
	}
	return highest
}

func resolveSeat(line string) (seatId int) {
	chars := strings.Split(line, "")
	min, max, needle := 0, SEATS, SEATS/2
	for _, char := range chars {
		min, max, needle = traverse(min, max, needle, char)
	}
	return min
}

func getFreeSeat(lines []string) (freeSeatId int) {
	occupied := map[int]bool{}
	for _, line := range lines {
		seatId := resolveSeat(line)
		occupied[seatId] = true
	}
	for i := 0; i < SEATS; i++ {
		if occupied[i] == true || !isWithinBounds(i) {
			continue
		}
		if isAdjacentOccupied(occupied, i) {
			return i
		}
	}
	fmt.Printf("Could not find a free seat. Check your input.")
	os.Exit(1)
	return 0
}

func isWithinBounds(position int) bool {
	return position-1 > 0 &&
		position+1 < SEATS
}

func isAdjacentOccupied(occupied map[int]bool, position int) bool {
	return position-1 > 0 &&
		position+1 < SEATS &&
		occupied[position-1] == true &&
		occupied[position+1] == true
}

func getData() []string {
	b, err := ioutil.ReadFile("day5/data")
	if err != nil {
		fmt.Println("Could not read input")
		os.Exit(1)
	}
	return strings.Split(string(b), "\n")
}
