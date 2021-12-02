package main

import (
	"fmt"
	"os"

	"github.com/walrusk/aoc2021/day1"
)

func main() {
	input_dir := os.Args[1] // input dir assumed to contain `day1/input.txt`, etc.

	intro()

	fmt.Println("\n-- DAY 1, PART 1 --")
	day1.Part1(input_dir)

	fmt.Println("\n-- DAY 1, PART 2 --")
	day1.Part2(input_dir)
}

func intro() {
	fmt.Printf(`
ADVENT OF CODE DEC 2021 in

  .oooooo.      .oooooo.           
 d8P'  ` + "`" + `Y8b    d8P'  ` + "`" + `Y8b          
888           888      888         
888           888      888         
888     ooooo 888      888         
` + "`" + `88.    .88'  ` + "`" + `88b    d88'         
 ` + "`" + `Y8bood8P'    ` + "`" + `Y8bood8P'` + "\n")
}
