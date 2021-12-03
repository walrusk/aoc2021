package day1

import (
	"strconv"

	"github.com/walrusk/aoc2021/pkg/iterable"
)

func Part2(input iterable.StringIterator) (int, error) {
	block_size := 3
	blocks := make([][]int, 0)
	increase_count := 0

	for {
		line, done := input.Next()
		if done {
			break
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}

		blocks = append(blocks, make([]int, 0))
		for index, b := range blocks {
			if len(blocks[index]) < block_size {
				blocks[index] = append(b, i)
			}
		}

		if len(blocks) > 1 && len(blocks[1]) == block_size {
			if sum_slice(blocks[0]) < sum_slice(blocks[1]) {
				increase_count += 1
			}
			blocks = blocks[1:]
		}
	}

	return increase_count, nil
}

func sum_slice(s []int) int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}
