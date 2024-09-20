package handlers

import "math"

// this function is used to calculate the square root of the variance or SD

func StandardDeviation(numbers []float64) float64 {
	variance := Variance(numbers)
	return math.Sqrt(variance)
}
