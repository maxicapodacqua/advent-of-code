package part2

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
	sample := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part2(scanner)
	expected := 14
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPart2_edgeCase(t *testing.T) {
	sample := `3-5
10-14
16-20
12-18
5-6

1
5
8
11
17
32`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part2(scanner)
	expected := 15
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
