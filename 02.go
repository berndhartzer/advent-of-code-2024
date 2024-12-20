package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func getUnsafeIdx(level []string, skipIdx int) int {
	currIdx, nextIdx := 0, 1
	lower, upper := 0, 0

	for nextIdx < len(level) {
		if currIdx == skipIdx {
			currIdx++
			continue
		}
		if nextIdx == skipIdx {
			nextIdx++
			continue
		}
		if currIdx == nextIdx {
			nextIdx++
		}

		curr, err := strconv.Atoi(level[currIdx])
		if err != nil {
			panic("could not convert str to number")
		}
		next, err := strconv.Atoi(level[nextIdx])
		if err != nil {
			panic("could not convert str to number")
		}

		diff := curr - next
		if diff == 0 {
			return currIdx
		}

		if lower == 0 {
			if diff < 0 {
				lower, upper = -3, -1
			} else if diff > 0 {
				lower, upper = 1, 3
			}
		}

		if diff < lower || diff > upper {
			return currIdx
		}

		currIdx++
		nextIdx++
	}

	return -1
}

func dayTwoPartOne(input []string) int {
	safe := 0

	for _, level := range input {
		splitted := strings.Split(level, " ")

		unsafe := getUnsafeIdx(splitted, -1)
		if unsafe == -1 {
			safe++
			continue
		}
	}

	return safe
}

func dayTwoPartTwo(input []string) int {
	safe := 0

	for _, level := range input {
		splitted := strings.Split(level, " ")

		unsafe := getUnsafeIdx(splitted, -1)
		if unsafe == -1 {
			safe++
			continue
		}

		unsafeWithRemoval := getUnsafeIdx(splitted, unsafe)
		if unsafeWithRemoval == -1 {
			safe++
			continue
		}

		// Probably a better way to handle this edge case but if the
		// bad index is at position 1 it might be the 0'th item
		// causing it
		if unsafe == 1 {
			unsafeWithRemoval = getUnsafeIdx(splitted, unsafe-1)
			if unsafeWithRemoval == -1 {
				safe++
				continue
			}
		}

		unsafeWithRemoval = getUnsafeIdx(splitted, unsafe+1)
		if unsafeWithRemoval == -1 {
			safe++
			continue
		}
	}

	return safe
}

func dayTwoTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(2)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			expected: 2,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			expected: 4,
		},
		"2": {
			input: []string{
				"39 39 42 43 46",
				"93 92 91 88 90 88",
				"1 2 3 4 5",
				"10 2 3 4 5",
				"1 20 3 4 5",
				"1 2 30 4 5",
				"1 2 3 40 5",
				"1 2 3 4 50",
			},
			expected: 8,
		},
		"3": {
			input: []string{
				"-39 -39 -42 -43 -46",
				"-93 -92 -91 -88 -90 -88",
				"-1 -2 -3 -4 -5",
				"-10 -2 -3 -4 -5",
				"-1 -20 -3 -4 -5",
				"-1 -2 -30 -4 -5",
				"-1 -2 -3 -40 -5",
				"-1 -2 -3 -4 -50",
			},
			expected: 8,
		},
		"4": {
			input: []string{
				"73 72 71 70 68 66",
				"16 15 14 11 10 9 6",
				"78 75 73 70 68 66 64 62",
				"47 46 45 45 43 41 39",
				"65 68 70 73 76 78 75 79",

				"65 68 70 73 76 78 75 79 79",
				"12 10 13 15 16 17 19 26",
				"4 2 6 8 9 9",
				"6 6 6 8 6",
				"21 21 22 22 28",
				"2 6 8 15 22",
				"5 4 3 2 1 4 4",
				"58 56 53 50 49 51 50 46",
				"88 85 84 86 80",
				"45 42 39 37 36 36 37",
				"69 68 67 66 65 62 62 62",
				"83 80 80 78 75 71",
			},
			expected: 5,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
