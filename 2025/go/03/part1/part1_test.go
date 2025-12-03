package part1

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	sample := `987654321111111
811111111111119
234234234234278
818181911112111`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part1(scanner)
	expected := 357
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
