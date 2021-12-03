package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/walrusk/aoc2021/pkg/day1"
	"github.com/walrusk/aoc2021/pkg/day2"
	"github.com/walrusk/aoc2021/pkg/day3"
	"github.com/walrusk/aoc2021/pkg/iterable"
	"github.com/walrusk/aoc2021/pkg/util"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(errors.New("no input provided"))
	}

	var answer int
	var dir = os.Args[1] // input dir assumed to contain `day1/input.txt`, etc.

	intro()

	// Day 1
	answer, _ = day1.Part1(day_input(1, dir))
	fmt.Printf("\nday 1\npt. 1 answer: %d\n", answer)
	answer, _ = day1.Part2(day_input(1, dir))
	fmt.Printf("pt. 2 answer: %d\n", answer)

	// Day 2
	answer, _ = day2.Part1(day_input(2, dir))
	fmt.Printf("\nday 2\npt. 1 answer: %d\n", answer)
	answer, _ = day2.Part2(day_input(2, dir))
	fmt.Printf("pt. 2 answer: %d\n", answer)

	// Day 3
	answer, _ = day3.Part1(day_input(3, dir))
	fmt.Printf("\nday 3\npt. 1 answer: %d\n", answer)
}

func intro() {
	fmt.Printf(`
** advent of code dec 2021 **
             in
   .oooooo.      .oooooo.           
  d8P'  ` + "`" + `Y8b    d8P'  ` + "`" + `Y8b          
 888           888      888         
 888           888      888         
 888     ooooo 888      888         
 ` + "`" + `88.    .88'  ` + "`" + `88b    d88'         
  ` + "`" + `Y8bood8P'    ` + "`" + `Y8bood8P'` + "\n\n")
}

func day_input(day int, input_dir string) *iterable.IterableFile {
	input_path := filepath.Join(util.CurrentDir(), input_dir, fmt.Sprintf("day%d/input.txt", day))
	return iterable.NewIterableFile(input_path)
}
