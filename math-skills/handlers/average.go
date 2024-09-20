package handlers

// this function calculates the average/mean of athe numbers entered into the txt file by summing all the entries and dividing them by the number of etries/length of numbers

func Average(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	average := sum / float64(len(numbers))
	return average
}
