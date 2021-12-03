package day2

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/walrusk/aoc2021/pkg/util"
)

func Part2(input_dir string) {
	input_path := filepath.Join(util.CurrentDir(), input_dir, "day2/input.txt")
	line_count := 0
	aim := 0
	depth := 0
	pos_x := 0

	util.FileByLine(input_path, func(line string) {
		words := strings.Fields(line)

		i, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatal(err)
		}

		switch words[0] {
		case "forward":
			pos_x += i
			depth += aim * i
		case "down":
			aim += i
		case "up":
			aim -= i
		}

		if depth < 0 {
			depth = 0
		}

		line_count += 1
	})

	fmt.Printf("Scanned %d lines\n", line_count)
	fmt.Printf("Final depth: %d\n", depth)
	fmt.Printf("Final x position: %d\n", pos_x)
	fmt.Printf("Multiplied: %d\n", pos_x*depth)
}
