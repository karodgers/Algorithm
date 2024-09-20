package functions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reads the data from the .txt file
func ReadData(filePath string) ([]float64, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var data []float64
	lines := strings.Split(string(content), "\n")
	count := 0
	for _, line := range lines {
		count++
		if line == "" {
			continue
		}
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing value in line: %v, %v", count, err)
		}
		data = append(data, value)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("data file is empty")
	}

	return data, nil
}
