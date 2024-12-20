package aoc

import (
	"fmt"
)

func dayFourPartOne(input []string) int {
	xmas := []byte("XMAS")

	directionMatch := func(x, y, xDir, yDir int) bool {
		for _, c := range xmas {
			if y < 0 {
				return false
			}
			if y > len(input)-1 {
				return false
			}
			if x < 0 {
				return false
			}
			if x > len(input[0])-1 {
				return false
			}

			if input[y][x] != c {
				return false
			}

			x += xDir
			y += yDir
		}

		return true
	}

	findAllXmas := func(x, y int) int {
		found := 0

		// north
		if directionMatch(x, y, 0, -1) {
			found++
		}

		// north east
		if directionMatch(x, y, 1, -1) {
			found++
		}

		// east
		if directionMatch(x, y, 1, 0) {
			found++
		}

		// south east
		if directionMatch(x, y, 1, 1) {
			found++
		}

		// south
		if directionMatch(x, y, 0, 1) {
			found++
		}

		// south west
		if directionMatch(x, y, -1, 1) {
			found++
		}

		// west
		if directionMatch(x, y, -1, 0) {
			found++
		}

		// north west
		if directionMatch(x, y, -1, -1) {
			found++
		}

		return found
	}

	total := 0
	for y, line := range input {
		for x, char := range line {
			if byte(char) == xmas[0] {
				total += findAllXmas(x, y)
			}
		}
	}

	return total
}

func dayFourPartTwo(input []string) int {
	mas := []byte("MAS")

	inXMas := func(x, y int) bool {
		if y == 0 {
			return false
		}
		if y == len(input)-1 {
			return false
		}
		if x == 0 {
			return false
		}
		if x == len(input[0])-1 {
			return false
		}

		// north west, south east
		diagonalOne := []byte{
			input[y-1][x-1],
			input[y][x],
			input[y+1][x+1],
		}

		diagonalOneMatch := false
		if diagonalOne[0] == mas[0] && diagonalOne[2] == mas[2] {
			diagonalOneMatch = true
		}
		if diagonalOne[0] == mas[2] && diagonalOne[2] == mas[0] {
			diagonalOneMatch = true
		}
		if !diagonalOneMatch {
			return false
		}

		// north east, south west
		diagonalTwo := []byte{
			input[y-1][x+1],
			input[y][x],
			input[y+1][x-1],
		}

		diagonalTwoMatch := false
		if diagonalTwo[0] == mas[0] && diagonalTwo[2] == mas[2] {
			diagonalTwoMatch = true
		}
		if diagonalTwo[0] == mas[2] && diagonalTwo[2] == mas[0] {
			diagonalTwoMatch = true
		}
		if !diagonalTwoMatch {
			return false
		}

		return true
	}

	total := 0
	for y, line := range input {
		for x, char := range line {
			if byte(char) == mas[1] {
				if inXMas(x, y) {
					total++
				}
			}
		}
	}

	return total
}

func dayFourTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(4)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			expected: 18,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			expected: 9,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
