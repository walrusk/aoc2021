package day2

import (
	"testing"

	"github.com/walrusk/aoc2021/pkg/iterable"
)

func TestPart2_Answer(t *testing.T) {
	test_input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	expected := 900
	answer, err := Part2(iterable.NewIterableSlice(test_input))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if answer != expected {
		t.Errorf("Answer was incorrect, got: %d, want: %d.", answer, expected)
	}
}

func TestPart2_CantFly(t *testing.T) {
	test_input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"up 10",
		"up 99",
		"forward 2",
	}

	expected := 0
	answer, err := Part2(iterable.NewIterableSlice(test_input))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if answer != expected {
		t.Errorf("Answer was incorrect, got: %d, want: %d.", answer, expected)
	}
}

func TestPart2_UnexpectedInput(t *testing.T) {
	test_input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"explode 8",
	}

	_, err := Part2(iterable.NewIterableSlice(test_input))
	if err == nil {
		t.Error("Should have errored on explode command.")
	}
}
