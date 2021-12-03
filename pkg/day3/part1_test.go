package day3

import (
	"testing"

	"github.com/walrusk/aoc2021/pkg/iterable"
)

func TestPart1_Answer(t *testing.T) {
	test_input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	expected := 198
	answer, err := Part1(iterable.NewIterableSlice(test_input))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if answer != expected {
		t.Errorf("Answer was incorrect, got: %d, want: %d.", answer, expected)
	}
}
