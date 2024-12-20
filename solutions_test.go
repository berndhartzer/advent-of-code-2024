package aoc

import (
	"fmt"
	"testing"
	"time"
)

type stringSliceToIntTestConfig struct {
	input     []string
	expected  int
	logResult bool
}

type stringToIntTestConfig struct {
	input     string
	expected  int
	logResult bool
}

func stringSliceToIntRunner(t *testing.T, tests map[string]stringSliceToIntTestConfig, fn func([]string) int) {
	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			start := time.Now()
			output := fn(cfg.input)
			finish := time.Since(start)
			if cfg.logResult {
				t.Log(fmt.Sprintf("\nsolution:\t%v\nelapsed time:\t%s", output, finish))
				return
			}

			if output != cfg.expected {
				t.Fatalf("Incorrect output - got: %v, want: %v", output, cfg.expected)
			}
		})
	}
}

func stringToIntRunner(t *testing.T, tests map[string]stringToIntTestConfig, fn func(string) int) {
	for name, cfg := range tests {
		cfg := cfg
		t.Run(name, func(t *testing.T) {
			start := time.Now()
			output := fn(cfg.input)
			finish := time.Since(start)
			if cfg.logResult {
				t.Log(fmt.Sprintf("\nsolution:\t%v\nelapsed time:\t%s", output, finish))
				return
			}

			if output != cfg.expected {
				t.Fatalf("Incorrect output - got: %v, want: %v", output, cfg.expected)
			}
		})
	}
}

func TestSolutions(t *testing.T) {
	t.Run("day 1", func(t *testing.T) {
		partOne, partTwo, err := dayOneTests()
		if err != nil {
			t.Errorf("failed to get tests: %v", err)
		}

		t.Run("part 1", func(t *testing.T) {
			stringSliceToIntRunner(t, partOne, dayOnePartOne)
		})

		t.Run("part 2", func(t *testing.T) {
			stringSliceToIntRunner(t, partTwo, dayOnePartTwo)
		})
	})

	t.Run("day 2", func(t *testing.T) {
		partOne, partTwo, err := dayTwoTests()
		if err != nil {
			t.Errorf("failed to get tests: %v", err)
		}

		t.Run("part 1", func(t *testing.T) {
			stringSliceToIntRunner(t, partOne, dayTwoPartOne)
		})

		t.Run("part 2", func(t *testing.T) {
			stringSliceToIntRunner(t, partTwo, dayTwoPartTwo)
		})
	})

	t.Run("day 3", func(t *testing.T) {
		partOne, partTwo, err := dayThreeTests()
		if err != nil {
			t.Errorf("failed to get tests: %v", err)
		}

		t.Run("part 1", func(t *testing.T) {
			stringSliceToIntRunner(t, partOne, dayThreePartOne)
		})

		t.Run("part 2", func(t *testing.T) {
			stringSliceToIntRunner(t, partTwo, dayThreePartTwo)
		})
	})

	t.Run("day 4", func(t *testing.T) {
		partOne, partTwo, err := dayFourTests()
		if err != nil {
			t.Errorf("failed to get tests: %v", err)
		}

		t.Run("part 1", func(t *testing.T) {
			stringSliceToIntRunner(t, partOne, dayFourPartOne)
		})

		t.Run("part 2", func(t *testing.T) {
			stringSliceToIntRunner(t, partTwo, dayFourPartTwo)
		})
	})

	t.Run("day 5", func(t *testing.T) {
		partOne, partTwo, err := dayFiveTests()
		if err != nil {
			t.Errorf("failed to get tests: %v", err)
		}

		t.Run("part 1", func(t *testing.T) {
			stringSliceToIntRunner(t, partOne, dayFivePartOne)
		})

		t.Run("part 2", func(t *testing.T) {
			stringSliceToIntRunner(t, partTwo, dayFivePartTwo)
		})
	})

	t.Run("day 6", func(t *testing.T) {
		partOne, partTwo, err := daySixTests()
		if err != nil {
			t.Errorf("failed to get tests: %v", err)
		}

		t.Run("part 1", func(t *testing.T) {
			stringSliceToIntRunner(t, partOne, daySixPartOne)
		})

		t.Run("part 2", func(t *testing.T) {
			stringSliceToIntRunner(t, partTwo, daySixPartTwo)
		})
	})
}
