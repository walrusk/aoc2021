package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/walrusk/aoc2021/pkg/iterable"
)

func Part2(input iterable.StringIterator) (int, error) {
	aim := 0
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
			log.Fatal(err)
		}
		command := words[0]

		switch command {
		case "forward":
			pos_x += i
			depth += aim * i
		case "down":
			aim += i
		case "up":
			aim -= i
		default:
			return 0, fmt.Errorf("Unknown command %s", command)
		}

		if depth < 0 { // Submarine can't shoot into the air after surfacing.
			depth = 0
		}
	}

	return pos_x * depth, nil
}
