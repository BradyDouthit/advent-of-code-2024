package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNums(line string) []int {
	delimeter := " "
	separated := strings.Split(line, delimeter)
	nums := []int{}

	for _, numStr := range separated {
		num, err := strconv.Atoi(numStr)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		nums = append(nums, num)
	}

	return nums
}

func getDifference(x int, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}

// A report only counts as safe if both of the following are true:
// - The levels are either all increasing or all decreasing.
// - Any two adjacent levels differ by at least one and at most three.
func getIsSafe(nums []int) bool {
	const minDifference = 1
	const maxDifference = 3

	isSafe := true

	shouldIncrease := nums[0] < nums[1]

	for i, num := range nums {
		if i < len(nums)-1 {
			next := nums[i+1]
			diff := getDifference(num, next)

			if diff < minDifference || diff > maxDifference {
				isSafe = false
			} else if shouldIncrease && next < num {
				isSafe = false
			} else if !shouldIncrease && next > num {
				isSafe = false
			}
		}
	}

	return isSafe
}

func main() {
	data, err := os.ReadFile("./data.txt")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rawReports := string(data)
	lines := strings.Split(rawReports, "\r\n")
	safeReports := 0
	unsafeReports := 0

	for _, line := range lines {
		nums := getNums(line)
		isSafe := getIsSafe(nums)

		if isSafe {
			safeReports += 1
		} else {
			unsafeReports += 1
		}
	}

	fmt.Printf("Out of %v reports, %v were safe and %v were unsafe", safeReports+unsafeReports, safeReports, unsafeReports)
}
