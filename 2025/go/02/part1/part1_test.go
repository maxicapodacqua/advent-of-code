package part1

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	sample := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part1(scanner)
	expected := 1227775554
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func Test_IsInvalidProjectID(t *testing.T) {
	type test struct {
		input string
		want  bool
	}
	tests := []test{
		{"11", true},
		{"99", true},
		{"1010", true},
		{"1188511885", true},
		{"222222", true},
		{"1698522", false},
		{"1698523", false},
		{"1698528", false},
		{"565653", false},
		{"565654", false},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := IsInvalidProjectID(tc.input)
			if got != tc.want {
				t.Errorf("Expected %v got %v for %v", tc.want, got, tc.input)
			}
		})
	}
}
