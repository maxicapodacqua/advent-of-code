package part2

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
	sample := `3   4
4   3
2   5
1   3
3   9
3   3`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part2(scanner)
	expected := 31
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}