package handlers

func Median(numbers []float64) float64 {
	// using bubble sort to arrange the numbers in ascending order
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				// Swap
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	length := len(numbers)

	middle := length / 2
	if length%2 == 0 {
		// If the length of the slice is even, average the two middle elements
		return (numbers[middle-1] + numbers[middle]) / 2
	} else {
		// If the length of the slice is odd, return the middle element
		return numbers[middle]
	}
}
