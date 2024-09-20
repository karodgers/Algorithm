package functions

import "math"

func LinearRegressionCalc(data []float64) (m, b, r float64) {
	n := float64(len(data))
	sumX, sumY, sumXY, sumX2, sumY2 := 0.0, 0.0, 0.0, 0.0, 0.0

	// Calculate sums
	for i, y := range data {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y
	}

	// slope (m) and y-intercept (b)
	m = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b = (sumY - m*sumX) / n

	// correlation coefficient (r)
	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))
	r = numerator / denominator

	return m, b, r
}
