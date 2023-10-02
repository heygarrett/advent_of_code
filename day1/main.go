package main

import (
	reader "2022"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := reader.ReadFileFromFS(os.DirFS("day1"), "input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	maxCalories, nil := GetMaxCalories(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	totalCaloriesTopThree, nil := GetTotalCaloriesTopThree(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Max calories carried by one elf: %d\n", maxCalories)
	fmt.Printf("Total calories carried by top three: %d\n", totalCaloriesTopThree)
}

func GetMaxCalories(input string) (int, error) {
	calorieCounts, err := getCalorieCounts(input)
	if err != nil {
		return 0, nil
	}

	return calorieCounts[0], nil
}

func GetTotalCaloriesTopThree(input string) (int, error) {
	calorieCounts, err := getCalorieCounts(input)
	if err != nil {
		return 0, err
	}

	return sumAll(calorieCounts[:3]), nil
}

func getCalorieCounts(input string) ([]int, error) {
	parsedInput := parseInput(input)
	calorieCounts := make([]int, 0)
	for _, elf := range parsedInput {
		intList, err := convertStrListToIntList(elf)
		if err != nil {
			return nil, err
		}
		intTotal := sumAll(intList)
		calorieCounts = append(calorieCounts, intTotal)
	}

	slices.Sort(calorieCounts)
	slices.Reverse(calorieCounts)

	return calorieCounts, nil
}

func parseInput(input string) [][]string {
	splitOnBlankLines := strings.Split(input, "\n\n")
	splitOnNewLines := make([][]string, 0)
	for _, contiguousLines := range splitOnBlankLines {
		splitOnNewLines = append(
			splitOnNewLines,
			strings.Split(contiguousLines, "\n"),
		)
	}

	return splitOnNewLines
}

func sumAll(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func convertStrListToIntList(strs []string) ([]int, error) {
	result := make([]int, 0)
	for _, s := range strs {
		converted, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return nil, err
		}
		result = append(result, int(converted))
	}

	return result, nil
}
