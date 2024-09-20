package pkg

import "math"

func StandardDeviation(nums []int, mean float64) float64 {
	if len(nums) < 2 {
		return 0
	}

	var sum float64
	for _, num := range nums {
		sum += math.Pow(float64(num)-mean, 2)
	}
	variance := sum / float64(len(nums)-1)
	return math.Sqrt(variance)
}
