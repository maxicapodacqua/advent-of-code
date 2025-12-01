package part1

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	sample := `3   4
4   3
2   5
1   3
3   9
3   3`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part1(scanner)
	expected := 11
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}