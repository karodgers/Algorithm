package main

import (
	"fmt"
	"os"

	"linear-stats/functions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your-program.go data.txt")
		return
	}

	filePath := os.Args[1]
	data, err := functions.ReadData(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	m, b, r := functions.LinearRegressionCalc(data)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
