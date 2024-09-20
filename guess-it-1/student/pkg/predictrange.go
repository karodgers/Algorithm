package pkg

import "math"

func PredictRange(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 100 // Default range
	}

	avg := Mean(nums)
	stdDev := StandardDeviation(nums, avg)

	// using SD for the range to cover ~95% of distribution
	lower := int(math.Floor(avg - 2*stdDev))
	upper := int(math.Ceil(avg + 2*stdDev))

	// Ensure lower bound is not negative
	if lower < 0 {
		lower = 0
	}

	// Ensure the range is at least 100
	if upper-lower < 100 {
		upper = lower + 100
	}

	return lower, upper
}
