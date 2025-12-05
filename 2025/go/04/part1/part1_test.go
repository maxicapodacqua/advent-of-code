package part1

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	sample := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part1(scanner)
	expected := 13
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
