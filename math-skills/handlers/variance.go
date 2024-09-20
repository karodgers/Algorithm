package handlers

// the function is used to calculate the variance by finding the sum of squared differences from the mean and dividing by the total number of values

func Variance(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		// calculating the sum of digits
		sum += num
	}
	mean := sum / float64(len(numbers))

	sumOfSquaredDifferences := 0.0

	for _, num := range numbers {
		difference := num - mean
		// calculate the sum of squared differences
		sumOfSquaredDifferences += difference * difference
	}

	variance := sumOfSquaredDifferences / float64(len(numbers))
	return variance
}
