package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Part1(input_dir string) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(filepath.Join(pwd, input_dir, "day2/input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	line_count := 0
	depth := 0
	pos_x := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		i, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatal(err)
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

		line_count += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Scanned %d lines\n", line_count)
	fmt.Printf("Final depth: %d\n", depth)
	fmt.Printf("Final x position: %d\n", pos_x)
	fmt.Printf("Multiplied: %d\n", pos_x*depth)
}
