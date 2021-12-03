package day1

import (
	"log"
	"strconv"

	"github.com/walrusk/aoc2021/pkg/iterable"
)

func Part1(input iterable.StringIterator) (int, error) {
	last_i := int(^uint(0) >> 1)
	increase_count := 0

	for {
		line, done := input.Next()
		if done {
			break
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		if i > last_i {
			increase_count += 1
		}

		last_i = i
	}

	return increase_count, nil
}
