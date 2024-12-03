package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func loadDataTask3(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}


func ComputeMul(data string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(data, -1)

	res := 0


	for _, match := range matches {
		if len(match) == 3 {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			res += num1 * num2
		}
	}

	return res
}

func ComputeConditionalMul(data string) int {
	splitData := strings.Split(data, "don't()")


	res := 0

	for val := range splitData {
		if val == 0 {
			res += ComputeMul(splitData[val])
		} else {
			doData := strings.Split(splitData[val],"do()")
			for i := 1; i < len(doData); i++ {
				res += ComputeMul(doData[i])
			}
		}
	}
	return res
}
