package aoc

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func dayOnePartOne(input []string) int {
	left, right := []int{}, []int{}

	for _, in := range input {
		splitted := strings.Split(in, "   ")

		leftInt, err := strconv.Atoi(splitted[0])
		if err != nil {
			panic("could not convert str to number")
		}
		rightInt, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic("could not convert str to number")
		}

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDist := 0
	for i, l := range left {
		dist := l - right[i]
		totalDist += abs(dist)
	}

	return totalDist
}

func dayOnePartTwo(input []string) int {
	left := []string{}
	right := map[string]int{}

	for _, in := range input {
		splitted := strings.Split(in, "   ")
		left = append(left, splitted[0])
		right[splitted[1]]++
	}

	total := 0
	for _, v := range left {
		rightCount, ok := right[v]
		if !ok {
			continue
		}

		n, err := strconv.Atoi(v)
		if err != nil {
			panic("could not convert str to number")
		}

		total += n * rightCount
	}

	return total
}

func dayOneTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(1)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"3   4",
				"4   3",
				"2   5",
				"1   3",
				"3   9",
				"3   3",
			},
			expected: 11,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"3   4",
				"4   3",
				"2   5",
				"1   3",
				"3   9",
				"3   3",
			},
			expected: 31,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
