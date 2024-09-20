package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"math-skills/handlers"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Arguments must be two")
		os.Exit(1)
	}
	// Read data from file
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split data into lines
	lines := strings.Split(string(data), "\n")

	// Convert lines to float64 numbers

	var numbers []float64

	for _, line := range lines {

		// Remove spaces between numbers
		line = strings.ReplaceAll(line, " ", "")

		// remove commas from number entries
		line = strings.ReplaceAll(line, ",", "")

		// Trim leading and trailing spaces
		line = strings.TrimSpace(line)

		if line == "" {
			continue // Skip empty lines
		}

		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error non-digit detected:", err)
			return
		}
		numbers = append(numbers, num)
	}

	// Check if there is only one entry in the data.txt file
	if len(numbers) == 1 {
		fmt.Println("Error: Enter atleast 2 data values to determine variance")
		return
	}
	// Check if the file is empty
	if len(numbers) == 0 {
		fmt.Println("Error: The data file is empty.")
		return
	}

	// call back the functions to calculate and print statistics
	// fmt.Println("numbers:", numbers)
	fmt.Println("Average:", handlers.RoundToInt(handlers.Average(numbers)))
	fmt.Println("Median:", handlers.RoundToInt(handlers.Median(numbers)))
	fmt.Println("Variance:", handlers.RoundToInt(handlers.Variance(numbers)))
	fmt.Println("Standard Deviation:", handlers.RoundToInt(handlers.StandardDeviation(numbers)))
}
