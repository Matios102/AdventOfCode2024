package main

import "fmt"

func main() {
	res_1, res_2, err := Task4("./input.txt")
	if err != nil {
		fmt.Printf("An error ocurred: %s", err)
		return
	}

	fmt.Printf("Task 4.1: %d\n", res_1)
	fmt.Printf("Task 4.2: %d\n", res_2)
}