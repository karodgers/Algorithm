package functions

import (
	"math"
	"testing"
)

func TestLinearRegression(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expectedM := 1.0
	expectedB := 1.0

	m, b := LinearRegression(nums)

	if m != expectedM {
		t.Errorf("Test failed: expected slope %v, got %v", expectedM, m)
	}
	if b != expectedB {
		t.Errorf("Test failed: expected intercept %v, got %v", expectedB, b)
	}
}

func TestMean(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := 3.0

	result := Mean(nums)

	if result != expected {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}
}

func TestStandardDeviation(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	mean := 3.0
	expected := math.Sqrt(2.5)

	result := StandardDeviation(nums, mean)

	if result != expected {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}
}
