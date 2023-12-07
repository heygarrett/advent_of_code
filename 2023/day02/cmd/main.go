package main

import (
	solutions "2023/day02"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	rawInput, err := fs.ReadFile(os.DirFS("day02"), "input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input := string(rawInput)

	part1 := solutions.Part1(input)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := solutions.Part2(input)
	fmt.Printf("Part 1: %d\n", part2)
}
