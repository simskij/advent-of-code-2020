package main

import (
	"fmt"
	day11 "github.com/simskij/advent-of-code-2020/internal/days/11"
	day12 "github.com/simskij/advent-of-code-2020/internal/days/12"
	"os"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
	day1 "github.com/simskij/advent-of-code-2020/internal/days/1"
	day2 "github.com/simskij/advent-of-code-2020/internal/days/2"
	day3 "github.com/simskij/advent-of-code-2020/internal/days/3"
	day4 "github.com/simskij/advent-of-code-2020/internal/days/4"
	day5 "github.com/simskij/advent-of-code-2020/internal/days/5"
	day5b "github.com/simskij/advent-of-code-2020/internal/days/5b"
	day6 "github.com/simskij/advent-of-code-2020/internal/days/6"
	day7 "github.com/simskij/advent-of-code-2020/internal/days/7"
	day8 "github.com/simskij/advent-of-code-2020/internal/days/8"
	day9 "github.com/simskij/advent-of-code-2020/internal/days/9"
	day10 "github.com/simskij/advent-of-code-2020/internal/days/10"
	"github.com/simskij/advent-of-code-2020/internal/types"
)

func main() {
	start := time.Now()

	solutions := []types.Solution{
		WithDuration(day1.GetSolution),
		WithDuration(day2.GetSolution),
		WithDuration(day3.GetSolution),
		WithDuration(day4.GetSolution),
		WithDuration(day5.GetSolution),
		WithDuration(day5b.GetSolution),
		WithDuration(day6.GetSolution),
		WithDuration(day7.GetSolution),
		WithDuration(day8.GetSolution),
		WithDuration(day9.GetSolution),
		WithDuration(day10.GetSolution),
		WithDuration(day11.GetSolution),
		WithDuration(day12.GetSolution),
	}

	fmt.Println("\n✨ Advent of Code 2020 ✨\n ")

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)
	for _, solution := range solutions {
		PrintStatsForDay(w, solution)
	}
	w.Flush()
	fmt.Printf("\nTotal duration %dµs\n", time.Since(start).Microseconds())

}

func PrintStatsForDay(w *tabwriter.Writer, solution types.Solution) {
	yellow := color.New(color.FgYellow, color.Bold)
	white := color.New(color.FgWhite, color.Bold)
	gray := color.New(color.FgHiWhite, color.Faint)

	var variation string
	if solution.Variation != "" {
		variation = fmt.Sprintf("- %-16s", solution.Variation)
	}
	fmt.Fprintf(w, "Day #%d %-20s\t\t", solution.Day, gray.Sprintf(variation))
	fmt.Fprintf(w, "%s %s\t", white.Sprint("★"), solution.Answers.A)
	fmt.Fprintf(w, "%s %s\t", yellow.Sprint("★"), solution.Answers.B)
	fmt.Fprintf(w, "%24v\n", gray.Sprintf("%dµs", solution.Duration.Microseconds()))
}

func WithDuration(fn func() types.Solution) types.Solution {
	start := time.Now()
	res := fn()
	res.Duration = time.Since(start)
	return res
}
