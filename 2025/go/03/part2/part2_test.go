package part2

import (
	"bufio"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
	sample := `987654321111111
811111111111119
234234234234278
818181911112111`
	scanner := bufio.NewScanner(strings.NewReader(sample))
	got := Part2(scanner)
	expected := 3121910778619
	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func Test_GetLargestJoltage(t *testing.T) {
	type test struct {
		name  string
		input []int
		want  int
	}

	tests := []test{
		{
			name:  "987654321111111",
			input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			want:  987654321111,
		},
		{
			name:  "811111111111119",
			input: []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			want:  811111111119,
		},
		{
			name:  "234234234234278",
			input: []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			want:  434234234278,
		},
		{
			name:  "818181911112111",
			input: []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			want:  888911112111,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetLargestJoltage(tc.input)
			if got != tc.want {
				t.Errorf("Expected %d, got %d", tc.want, got)
			}
		})
	}
}
