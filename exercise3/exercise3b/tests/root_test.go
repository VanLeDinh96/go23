package tests

import (
	"testing"

	"github.com/diegovanne/go23/exercise3/exercise3b/cmd"
)

func TestNumDifferentIntegers(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"A1b01c001", 1},
		{"abc", 0},
		{"123abc456", 2},
		{"000000", 1},
	}

	for _, test := range tests {
		result := cmd.NumDifferentIntegers(test.input)
		if result != test.expected {
			t.Errorf("NumDifferentIntegers(%s) returned %d, expected %d", test.input, result, test.expected)
		}
	}
}
