package part1

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	sample := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	reader := strings.NewReader(sample)

	scanner := bufio.NewScanner(reader)
	got := Part1(reader, scanner)
	expected := 4277556
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
