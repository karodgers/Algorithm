package functions

import "math"

// Range prediction using LR and SD
func PredictRange(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 50 // Default range
	}

	mean := Mean(nums)
	stdDev := StandardDeviation(nums, mean)
	m, b := LinearRegression(nums)

	// Using LR for prediction
	predicted := m*float64(len(nums)) + b

	lower := int(math.Floor(predicted - 2*stdDev))
	upper := int(math.Ceil(predicted + 2*stdDev))

	if lower < 0 {
		lower = 0
	}

	if upper-lower < 50 {
		upper = lower + 50
	}

	return lower, upper
}
