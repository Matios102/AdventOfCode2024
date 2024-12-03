package main

import (
	"fmt"
	
)


func main() {

	t1_1, t1_2, err := Task("./inputs/input1.txt")
	if err != nil {
		fmt.Printf("An error ocurred: %s", err)
		return
	}

	fmt.Printf("Task1.1: %d\n", t1_1)
	fmt.Printf("Task1.2: %d\n", t1_2)
}