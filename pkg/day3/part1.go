package day3

import (
	"strconv"

	"github.com/walrusk/aoc2021/pkg/iterable"
)

func Part1(input iterable.StringIterator) (int, error) {
	gamma_counts, line_count := make(map[int]int), 0

	for {
		line, done := input.Next()
		if done {
			break
		}

		line_count += 1
		for i, c := range line {
			if c == '1' {
				gamma_counts[i] += 1
			}
		}
	}

	g_rate, e_rate := "", ""
	for i := 0; i < len(gamma_counts); i++ {
		if gamma_counts[i] > line_count/2 {
			g_rate += "1"
			e_rate += "0"
		} else {
			g_rate += "0"
			e_rate += "1"
		}
	}

	g, err := strconv.ParseInt(g_rate, 2, 64)
	if err != nil {
		return 0, err
	}

	e, err := strconv.ParseInt(e_rate, 2, 64)
	if err != nil {
		return 0, err
	}

	return int(g * e), nil
}
