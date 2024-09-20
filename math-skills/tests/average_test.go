package handlers

import (
	"testing"

	"math-skills/handlers"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
	}{
		{"TestSingleElement", []float64{5.5}, 5.5},                 // Test with a single element
		{"TestPositiveNumbers", []float64{1.0, 2.0, 3.0}, 2.0},     // Test with positive numbers
		{"TestNegativeNumbers", []float64{-1.0, -2.0, -3.0}, -2.0}, // Test with negative numbers
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := handlers.Average(test.numbers)
			if result != test.expected {
				t.Errorf("Expected %f, got %f instead", test.expected, result)
			}
		})
	}
}
