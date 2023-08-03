package tests

import (
	"testing"

	"github.com/diegovanne/go23/exercise3/exercise3a/cmd"
)

func TestCountRectangles(t *testing.T) {
	arr := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	expectedCount := 6 // The expected count of rectangles in the given array

	count := cmd.CountRectangles(arr)

	if count != expectedCount {
		t.Errorf("CountRectangles returned %d, expected %d", count, expectedCount)
	}
}
