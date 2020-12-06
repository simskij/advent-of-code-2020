package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/simskij/advent-of-code-2020/internal/data"
	"github.com/simskij/advent-of-code-2020/internal/types"
)

func GetSolution() types.Solution {
	lines := data.GetData("inputs/4", "\n\n")
	return types.Solution{
		Day: 4,
		Answers: types.Answers{
			A: fmt.Sprintf("%d", validatePassports(lines, containsRequiredFields)),
			B: fmt.Sprintf("%d", validatePassports(lines, containsValidData)),
		},
	}
}

func validatePassports(lines []string, validationFn func(string) bool) int {
	valid := 0
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\n", " ")
		if validationFn(line) {
			valid++
		}
	}
	return valid
}

func containsValidData(line string) bool {
	fields := strings.Split(line, " ")
	kv := map[string]string{}
	for _, pair := range fields {
		kv[pair[0:3]] = pair[4:]
	}

	valid := 0
	for key, val := range kv {
		if (key == "byr" && isValidNumberBetween(val, 1920, 2002)) ||
			(key == "iyr" && isValidNumberBetween(val, 2010, 2020)) ||
			(key == "eyr" && isValidNumberBetween(val, 2020, 2030)) ||
			(key == "hgt" && isValidHeight(val)) ||
			(key == "hcl" && matchesPattern(val, "#[a-z0-9]{6}")) ||
			(key == "ecl" && isValidColor(val)) ||
			(key == "pid" && len(val) == 9 && matchesPattern(val, "[0-9]{9}")) {
			valid++
		}
	}
	return valid == 7
}

func isValidColor(val string) bool {
	for _, opt := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if val == opt {
			return true
		}
	}
	return false
}

func matchesPattern(val, pattern string) bool {
	if ok, err := regexp.MatchString(pattern, val); ok == true && err == nil {
		return true
	}
	return false
}

func isValidHeight(val string) bool {
	if strings.Contains(val, "cm") {
		hgt, _ := strconv.Atoi(strings.TrimRight(val, "cm"))
		if hgt >= 150 && hgt <= 193 {
			return true
		}
	}
	if strings.Contains(val, "in") {
		hgt, _ := strconv.Atoi(strings.TrimRight(val, "in"))
		if hgt >= 59 && hgt <= 76 {
			return true
		}
	}
	return false
}

func isValidNumberBetween(val string, from, to int) bool {
	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if num >= from && num <= to {
		return true
	}
	return false
}

func containsRequiredFields(line string) bool {
	return strings.Contains(line, "byr:") &&
		strings.Contains(line, "iyr:") &&
		strings.Contains(line, "eyr:") &&
		strings.Contains(line, "hgt:") &&
		strings.Contains(line, "hcl:") &&
		strings.Contains(line, "ecl:") &&
		strings.Contains(line, "pid:")
}
