package main

import "fmt"

func main() {
	t2_1, t2_2, err := Task("./inputs/input2.txt")
	if err != nil {
		fmt.Printf("An error ocurred: %s", err)
		return
	}

	fmt.Printf("Task2.1: %d\n", t2_1)
	fmt.Printf("Task2.2: %d\n", t2_2)
}