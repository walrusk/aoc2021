package day1

import (
	"testing"

	"github.com/walrusk/aoc2021/pkg/iterable"
)

func TestPart2_Answer(t *testing.T) {
	test_input := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}

	expected := 5
	answer, err := Part2(iterable.NewIterableSlice(test_input))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if answer != expected {
		t.Errorf("Answer was incorrect, got: %d, want: %d.", answer, expected)
	}
}
