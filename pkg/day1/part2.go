package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Part2(input_dir string) {
	block_size := 3
	blocks := make([][]int, 0)
	line_count := 0
	increase_count := 0

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(filepath.Join(pwd, input_dir, "day1/input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
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

		line_count += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Scanned %d lines\n", line_count)
	fmt.Printf("Counted increases: %d\n", increase_count)
}

func sum_slice(s []int) int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}
