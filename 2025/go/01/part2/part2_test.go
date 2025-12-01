package part2

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
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
	got := Part2(50, scanner)
	expected := 6
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
func TestPart2_overspin(t *testing.T) {
	sample := `R1000`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part2(50, scanner)
	expected := 10
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
func TestPart2_R50R50(t *testing.T) {
	sample := `R50
R50`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part2(50, scanner)
	expected := 1
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
func TestPart2_L50L50(t *testing.T) {
	sample := `L50
L50`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part2(50, scanner)
	expected := 1
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestPart2_R200(t *testing.T) {
	sample := `R200
L1`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part2(0, scanner)
	expected := 2
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
func TestPart2_Playground(t *testing.T) {
	sample := `L75
R50`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part2(50, scanner)
	expected := 2
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}