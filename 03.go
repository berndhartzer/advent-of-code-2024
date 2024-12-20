package aoc

import (
	"fmt"
	"regexp"
)

func dayThreePartOne(input []string) int {
	mulRe := regexp.MustCompile(`mul\(\d+,\d+\)`)

	total := 0

	for _, line := range input {
		found := mulRe.FindAllString(line, -1)

		for _, mulBlock := range found {
			nums := mulBlock[4 : len(mulBlock)-1]

			multiplier := 1
			multiply := []int{0, 0}
			numIdx := 0
			for k := len(nums) - 1; k >= 0; k-- {
				if nums[k] == ',' {
					numIdx++
					multiplier = 1
					continue
				}

				multiply[numIdx] += int(nums[k]-'0') * multiplier
				multiplier *= 10
			}

			total += multiply[1] * multiply[0]
		}
	}

	return total
}

func dayThreePartTwo(input []string) int {
	mulRe := regexp.MustCompile(`mul\(\d+,\d+\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	total := 0

	doMultiply := true

	for _, line := range input {
		mulFound := mulRe.FindAllStringIndex(line, -1)
		doFound := doRe.FindAllStringIndex(line, -1)
		dontFound := dontRe.FindAllStringIndex(line, -1)

		for {
			if len(mulFound) == 0 && len(doFound) == 0 && len(dontFound) == 0 {
				break
			}

			nextIdx := len(line)
			nextMulIdx, nextDoIdx, nextDontIdx := len(line), len(line), len(line)

			if len(mulFound) > 0 {
				nextMulIdx = mulFound[0][0]
			}
			if len(doFound) > 0 {
				nextDoIdx = doFound[0][0]
			}
			if len(dontFound) > 0 {
				nextDontIdx = dontFound[0][0]
			}

			nextIdx = nextMulIdx
			if nextDoIdx < nextIdx {
				nextIdx = nextDoIdx
			}
			if nextDontIdx < nextIdx {
				nextIdx = nextDontIdx
			}

			switch nextIdx {
			case nextMulIdx:
				if !doMultiply {
					mulFound = mulFound[1:]
					continue
				}

				nums := line[nextMulIdx+4 : mulFound[0][1]-1]

				multiplier := 1
				multiply := []int{0, 0}
				numIdx := 0

				for k := len(nums) - 1; k >= 0; k-- {
					if nums[k] == ',' {
						numIdx++
						multiplier = 1
						continue
					}

					multiply[numIdx] += int(nums[k]-'0') * multiplier
					multiplier *= 10
				}

				total += multiply[1] * multiply[0]

				mulFound = mulFound[1:]

			case nextDoIdx:
				doMultiply = true
				doFound = doFound[1:]

			case nextDontIdx:
				doMultiply = false
				dontFound = dontFound[1:]
			}
		}
	}

	return total
}

func dayThreeTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(3)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			expected: 161,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			expected: 48,
		},
		"2": {
			input: []string{
				"don't()xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))mul(2,4)",
			},
			expected: 48,
		},
		"3": {
			input: []string{
				"do()xmul(2,4)mul(2,4)&mul[3,7]!^_mul(5,5)+mul(32,64](mul(11,8)undon't()?mul(8,5))",
				"xmul(2,4)mul(2,4)&mul[3,7]!^_mul(5,5)+mul(32,64](mul(11,8)undon't()?mul(8,5))",
			},
			expected: 129,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
