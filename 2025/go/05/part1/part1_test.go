package part1

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
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
	got := Part1(scanner)
	expected := 3
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
