package handlers

import (
	"math"
	"testing"

	"math-skills/handlers"
)

func TestVariance(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
	}{
		{
			name:     "Test with positive numbers",
			numbers:  []float64{2, 4, 6, 8, 10},
			expected: 8,
		},
		{
			name:     "Test with negative numbers",
			numbers:  []float64{-2, -4, -6, -8, -10},
			expected: 8,
		},
		{
			name:     "Test with mixed positive and negative numbers",
			numbers:  []float64{-2, 4, -6, 8, -10},
			expected: 42.56,
		},
		{
			name:     "Test with floating point numbers",
			numbers:  []float64{1.5, 2.5, 3.5, 4.5, 5.5},
			expected: 2.0,
		},
		{
			name:     "Test with single number",
			numbers:  []float64{10},
			expected: 0,
		},
		{
			name:     "Test with empty slice",
			numbers:  []float64{},
			expected: math.NaN(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := handlers.Variance(test.numbers)
			if math.Abs(actual-test.expected) > 0.00001 {
				t.Errorf("Test '%s' failed: expected variance %f, got %f", test.name, test.expected, actual)
			}
		})
	}
}
