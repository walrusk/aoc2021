package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Part1(input_dir string) {
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
	line_count := 0
	last_i := int(^uint(0) >> 1)
	increase_count := 0
	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		if i > last_i {
			increase_count += 1
		}

		last_i = i
		line_count += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Scanned %d lines\n", line_count)
	fmt.Printf("Counted increases: %d\n", increase_count)
}
