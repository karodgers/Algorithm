package functions

// Determine slope (m) and intercept (b) of LR line.
func LinearRegression(nums []int) (m, b float64) {
	n := float64(len(nums))
	if n == 0 {
		return 0, 0
	}

	var sumX, sumY, sumXY, sumX2, sumY2 float64
	for i, y := range nums {
		x := float64(i)
		sumX += x
		sumY += float64(y)
		sumXY += x * float64(y)
		sumX2 += x * x
		sumY2 += float64(y) * float64(y)
	}

	m = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b = (sumY - m*sumX) / n

	return m, b
}
