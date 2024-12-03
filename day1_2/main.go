package main

import (
	"fmt"
	
)


func main() {

	t1_1, t1_2, err := Task1("./inputs/input1.txt")
	if err != nil {
		fmt.Printf("An error ocurred: %s", err)
		return
	}

	t2_1, t2_2, err := Task2("./inputs/input2.txt")
	if err != nil {
		fmt.Printf("An error ocurred: %s", err)
		return
	}

	fmt.Printf("Task1.1: %d\n", t1_1)
	fmt.Printf("Task1.2: %d\n", t1_2)
	fmt.Printf("Task2.1: %d\n", t2_1)
	fmt.Printf("Task2.2: %d\n", t2_2)
}