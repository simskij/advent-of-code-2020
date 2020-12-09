package day7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/simskij/advent-of-code-2020/internal/data"
	"github.com/simskij/advent-of-code-2020/internal/types"
)

func GetSolution() types.Solution {
	lines := data.GetData("inputs/7", "\n")
	lookup := buildBagMap(lines)
	a := getGoldenPossibilities(lookup)
	b := getContentOfGolden(lookup)
	return types.Solution{
		Day: 7,
		Answers: types.Answers{
			A: fmt.Sprintf("%d", a),
			B: fmt.Sprintf("%d", b),
		},
	}
}

type Bag struct {
	Type  string
	Count int
}

func getContentOfGolden(bagMap map[string][]Bag) int {
	return getContentOfBag("shiny gold", bagMap)
}

func getContentOfBag(current string, bagMap map[string][]Bag) int {
	content := bagMap[current]
	total := 0

	for _, bag := range content {
		if bagMap[bag.Type] != nil {
			total += bag.Count * getContentOfBag(bag.Type, bagMap)
		}
		total += bag.Count
	}
	return total
}

func buildBagMap(lines []string) map[string][]Bag {
	lookup := map[string][]Bag{}

	for _, line := range lines {
		parts := strings.Split(line, " bags contain ")
		bag := parts[0]
		contains := strings.TrimRight(parts[1], ".")
		if contains == "no other bags" {
			continue
		}
		for _, each := range strings.Split(contains, ", ") {
			r, _ := regexp.Compile("([0-9]*) (.*) bag")
			matches := r.FindStringSubmatch(each)
			count, _ := strconv.Atoi(matches[1])
			lookup[bag] = append(lookup[bag], Bag{Type: matches[2], Count: count})
		}
	}
	return lookup
}

func getGoldenPossibilities(bagMap map[string][]Bag) int {

	contains := 0
	total := 0
	for _, content := range bagMap {
		ok, _ := containsGolden(bagMap, content)
		total++
		if ok {
			contains++
		}

	}
	return contains
}

func containsGolden(lookup map[string][]Bag, content []Bag) (bool, string) {
	for _, entry := range content {
		if entry.Type == "shiny gold" {
			return true, entry.Type
		}
		ok, next := containsGolden(lookup, lookup[entry.Type])
		if ok {
			return ok, entry.Type + ", " + next
		}
	}
	return false, ""
}
