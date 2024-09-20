package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"guess-it-1/student/pkg"
)

const windowSize = 4

func main() {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		nums = append(nums, num)
		if len(nums) > windowSize {
			nums = nums[1:] // Keep the last windowSize numbers only
		}

		if len(nums) > 1 {
			lower, upper := pkg.PredictRange(nums)
			fmt.Printf("%d %d\n", lower, upper)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
