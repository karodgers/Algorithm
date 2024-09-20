package handlers

import (
	"testing"

	"math-skills/handlers"
)

func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
	}{
		{"TestSingleElement", []float64{5.5}, 5.5},                          // Test with a single element
		{"TestOddLength", []float64{1.0, 2.0, 3.0, 4.0, 5.0}, 3.0},          // Test with odd length slice
		{"TestEvenLength", []float64{1.0, 2.0, 3.0, 4.0}, 2.5},              // Test with even length slice
		{"TestSortedSlice", []float64{-3.0, -2.0, -1.0, 0.0, 1.0}, -1.0},    // Test with sorted slice
		{"TestReverseSortedSlice", []float64{5.0, 4.0, 3.0, 2.0, 1.0}, 3.0}, // Test with reverse sorted slice
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := handlers.Median(test.numbers)
			if result != test.expected {
				t.Errorf("Expected %f, got %f instead", test.expected, result)
			}
		})
	}
}
