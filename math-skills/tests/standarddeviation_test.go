package handlers

import (
	"math"
	"testing"

	"math-skills/handlers"
)

// TestStandardDeviation tests the StandardDeviation function
func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		numbers  []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 1.4142135623730951},      // Standard deviation for {1, 2, 3, 4, 5} is approximately 1.5811388300841898
		{[]float64{10, 20, 30, 40, 50}, 14.142135623730951}, // Standard deviation for {10, 20, 30, 40, 50} is approximately 15.811388300841896
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := handlers.StandardDeviation(test.numbers)
		if math.Abs(result-test.expected) > 0.00000001 { // Tolerating small differences due to floating-point arithmetic
			t.Errorf("StandardDeviation(%v) = %v, expected %v", test.numbers, result, test.expected)
		}
	}
}
