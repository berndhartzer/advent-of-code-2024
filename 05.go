package aoc

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func dayFive(input []string) (int, int) {
	orderRules := map[string][]string{}

	inputLooper := 0
	for ; inputLooper < len(input); inputLooper++ {
		if input[inputLooper] == "" {
			inputLooper++
			break
		}

		splitted := strings.Split(input[inputLooper], "|")

		gotOrderDependencies, ok := orderRules[splitted[0]]
		if !ok {
			gotOrderDependencies = []string{}
		}
		gotOrderDependencies = append(gotOrderDependencies, splitted[1])
		orderRules[splitted[0]] = gotOrderDependencies
	}

	// ordered, fixed
	totals := []int{0, 0}
	for ; inputLooper < len(input); inputLooper++ {
		splitted := strings.Split(input[inputLooper], ",")

		rightOrder := true
		seen := map[string]bool{}
		for _, page := range splitted {
			seen[page] = true

			orderRule, ok := orderRules[page]
			if !ok {
				continue
			}

			for _, dep := range orderRule {
				_, ok := seen[dep]
				if ok {
					rightOrder = false
					break
				}
			}

			if !rightOrder {
				break
			}
		}

		totalsIdx := 0

		if !rightOrder {
			totalsIdx = 1

			slices.SortStableFunc(splitted, func(a, b string) int {
				orderRule, ok := orderRules[a]
				if !ok {
					return 0
				}

				for _, dep := range orderRule {
					if dep == b {
						return -1
					}
				}

				return 0
			})
		}

		middle := splitted[len(splitted)/2]
		n, err := strconv.Atoi(middle)
		if err != nil {
			panic("could not convert str to number")
		}
		totals[totalsIdx] += n
	}

	return totals[0], totals[1]
}

func dayFivePartOne(input []string) int {
	total, _ := dayFive(input)
	return total
}

func dayFivePartTwo(input []string) int {
	_, total := dayFive(input)
	return total
}

func dayFiveTests() (map[string]stringSliceToIntTestConfig, map[string]stringSliceToIntTestConfig, error) {
	fileInput, err := getInput(5)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get input: %v", err)
	}
	input := fileInput.asStringSlice()

	partOne := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"47|53",
				"97|13",
				"97|61",
				"97|47",
				"75|29",
				"61|13",
				"75|53",
				"29|13",
				"97|29",
				"53|29",
				"61|53",
				"97|53",
				"61|29",
				"47|13",
				"75|47",
				"97|75",
				"47|61",
				"75|61",
				"47|29",
				"75|13",
				"53|13",
				"",
				"75,47,61,53,29",
				"97,61,53,29,13",
				"75,29,13",
				"75,97,47,61,53",
				"61,13,29",
				"97,13,75,29,47",
			},
			expected: 143,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	partTwo := map[string]stringSliceToIntTestConfig{
		"1": {
			input: []string{
				"47|53",
				"97|13",
				"97|61",
				"97|47",
				"75|29",
				"61|13",
				"75|53",
				"29|13",
				"97|29",
				"53|29",
				"61|53",
				"97|53",
				"61|29",
				"47|13",
				"75|47",
				"97|75",
				"47|61",
				"75|61",
				"47|29",
				"75|13",
				"53|13",
				"",
				"75,47,61,53,29",
				"97,61,53,29,13",
				"75,29,13",
				"75,97,47,61,53",
				"61,13,29",
				"97,13,75,29,47",
			},
			expected: 123,
		},
		"solution": {
			input:     input,
			logResult: true,
		},
	}

	return partOne, partTwo, nil
}
