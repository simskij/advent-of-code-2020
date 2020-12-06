package types

import "time"

type Solution struct {
	Day       int
	Answers   Answers
	Variation string
	Duration  time.Duration
}

type Answers struct {
	A string
	B string
}
