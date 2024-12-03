package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Task2(filePath string) (int, int, error) {
	data, err := loadDataTask2(filePath)
	if err != nil {
		return 0, 0, err
	}

	t1 := ComputeSafe(data)
	t2 := ComputeSafeWithDamping(data)
	return t1, t2, nil
}

func ComputeSafe(data [][]int) int {
	count := 0
	for _, x := range data {
		if isSafe(x) {
			count++
		}
	}
	return count
}

func ComputeSafeWithDamping(data [][]int) int {
	count := 0
	for _, x := range data {
		if isSafe(x) {
			count++
			continue
		}
		for idx := range x {
			if isSafeWithRemoval(x, idx) {
				count++
				break
			}
		}
	}
	return count
}

func isSafe(x []int) bool {
	if len(x) < 2 {
		return true
	}

	isIncreasing := x[1] > x[0]

	for i := 1; i < len(x); i++ {
		diff := x[i] - x[i-1]
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
		if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			return false
		}
	}
	return true
}

func isSafeWithRemoval(x []int, removeIdx int) bool {
    if len(x) - 1 < 2 {
        return true
    }

    var isIncreasing bool

	if removeIdx == 0 {
        isIncreasing = x[2] > x[1]
    } else if removeIdx == 1 {
        isIncreasing = x[2] > x[0]
    } else {
        isIncreasing = x[1] > x[0]
    }

	for i := 1; i < len(x); i++ {
		var diff int

		if((i == removeIdx && i == len(x) -1) || ((i-1) == removeIdx && (i-1) == 0)) {
			continue
		}

		if i == removeIdx && i != len(x) -1 {
			diff = x[i+1] - x[i-1]
		} else if (i-1) == removeIdx && (i-1) != 0 {
			diff = x[i] - x[i-2]
		} else {
			diff = x[i] - x[i-1]
		}
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
		if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			return false
		}

		
	}
	return true
}

func loadDataTask2(filepath string) ([][]int, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	var res [][]int

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		numbersStr := strings.Split(line, " ")
		var numbers []int

		for _, numStr := range numbersStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("failed to convert number: %w", err)
			}
			numbers = append(numbers, num)
		}

		res = append(res, numbers)
	}

	return res, nil
}