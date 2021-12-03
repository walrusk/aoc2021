package day2

import (
	"strconv"
	"strings"

	"github.com/walrusk/aoc2021/pkg/iterable"
)

func Part1(input iterable.StringIterator) (int, error) {
	depth := 0
	pos_x := 0

	for {
		line, done := input.Next()
		if done {
			break
		}

		words := strings.Fields(line)

		i, err := strconv.Atoi(words[1])
		if err != nil {
			return 0, err
		}

		switch words[0] {
		case "forward":
			pos_x += i
		case "down":
			depth += i
		case "up":
			depth -= i
		}

		if depth < 0 {
			depth = 0
		}
	}

	return pos_x * depth, nil
}
