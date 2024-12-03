package main

import "fmt"

func main() {
	data, err := loadDataTask3(".inputs/input.txt")
	if err != nil {
		fmt.Printf("failed to compute mul: %s\n", err)
		return
	}

	t1 := ComputeMul(data)
	t2 := ComputeConditionalMul(data)

	fmt.Printf("Task3.1: %d\n", t1)
	fmt.Printf("Task3.2: %d\n", t2)
}