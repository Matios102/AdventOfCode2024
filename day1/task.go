package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)


func Task(filePath string) (int, int, error) {
	x,y,err := loadDataTask1(filePath)
	if err != nil {
		fmt.Printf("An error ocurred: %s", err)
		return 0,0,err
	}

	t1 := ComputeDistance(x,y)
	t2 := ComputeSimilarity(x,y)
	

	return t1, t2, nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func loadDataTask1(filepath string) ([]int, []int, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, nil, err
	}

	var x, y []int

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		strNums := strings.Split(line, "   ")

		if(len(strNums) != 2){
			return x, y, nil
		}
		num1, err := strconv.Atoi(strNums[0])
		if err != nil {
			return nil, nil, err
		}
		num2, err := strconv.Atoi(strNums[1])
		if err != nil {
			return nil, nil, err
		}

		x = append(x, num1)
		y = append(y, num2)
	}

	return x, y, nil
}


func upHeap(idx int, x *[]int) {
	if idx == 1 {
		return
	}
	
	parent := idx / 2

	if (*x)[parent] > (*x)[idx] {
		swap(parent, idx, x)
		upHeap(parent, x)
	}
}

func downHeap(idx int, x *[]int) {
	if(idx == len((*x))) {
		return
	}

	left := 2*idx
	right := 2*idx + 1

	swapIdx := idx

	if(left < len((*x))) {
		if((*x)[swapIdx] > (*x)[left]) {
			swapIdx = left
		}
	}

	if(right < len((*x))) {
		if((*x)[swapIdx] > (*x)[right]) {
			swapIdx = right
		}
	}

	if(swapIdx != idx) {
		swap(swapIdx, idx, x)
		downHeap(swapIdx, x)
	}
}

func swap(parent, idx int, x *[]int) {
	tmp := (*x)[parent]

	(*x)[parent] = (*x)[idx]
	(*x)[idx] = tmp
}


func ComputeSimilarity(x, y []int) int {
	mapY := make(map[int]int)
	sim := 0

	for i := range len(x) {
		mapY[y[i]]++
	}


	for i := 0 ; i < len(x); i++ {
		val := x[i]
		sim += mapY[val] * val
	}
	
	return sim
}

func ComputeDistance(x,y []int) int {
	heapX := []int{-1}
	heapY := []int{-1}
	res := 0
	for i := 0; i < len(x); i++ {
		heapX = append(heapX, x[i])
		heapY = append(heapY, y[i])
		upHeap(i+1, &heapX)
		upHeap(i+1, &heapY)
	}

	for i:=1; i <= len(x); i++ {
		res += abs(heapX[1] - heapY[1])
		heapX[1] = MaxInt
		heapY[1] = MaxInt

		downHeap(1, &heapX)
		downHeap(1, &heapY)
	}

	return res
}