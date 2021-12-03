package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/walrusk/aoc2021/pkg/day1"
	"github.com/walrusk/aoc2021/pkg/day2"
	"github.com/walrusk/aoc2021/pkg/iterable"
	"github.com/walrusk/aoc2021/pkg/util"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(errors.New("no input provided"))
	}
	input_dir := os.Args[1] // input dir assumed to contain `day1/input.txt`, etc.

	intro()

	fmt.Println("\nday 1, part 1")
	day1.Part1(input_dir)

	fmt.Println("\nday 1, part 2")
	day1.Part2(input_dir)

	fmt.Println("\nday 2, part 1")
	day2.Part1(input_dir)

	fmt.Println("\nday 2, part 2")
	day2_answer, _ := day2.Part2(day_input(2, input_dir))
	fmt.Printf("answer: %d\n", day2_answer)
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
