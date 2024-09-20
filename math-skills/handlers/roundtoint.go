package handlers

import "math"

// this function returns the nearest integer by rounding off the results of each calculations

func RoundToInt(num float64) int {
	return int(math.Round(num))
}
