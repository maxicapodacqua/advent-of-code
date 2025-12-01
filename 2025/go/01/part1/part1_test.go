package part1

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	sample := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part1(scanner)
	expected := 3
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}