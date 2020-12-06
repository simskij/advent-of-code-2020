package day5

import (
	"fmt"
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
		Variation: "bitwise",
		Answers: types.Answers{
			A: fmt.Sprintf("%d", resolveSeats(lines)),
			B: fmt.Sprintf("%d", getFreeSeat(lines)),
		},
	}
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

	bin := 0b0000000000
	for i, char := range chars {
		if char == "B" || char == "R" {
			bin = bin | (1 << (10 - (i + 1)))
		}
	}
	return bin
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
