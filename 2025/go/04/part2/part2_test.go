package part2

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
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
	got := Part2(scanner)
	expected := 43
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
