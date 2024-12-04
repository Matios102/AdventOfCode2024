package main

import (
	"os"
	"strings"
)

var XMAS_directions = [8][2]int{
	{-1, 0}, // Up
	{-1, 1}, // Up-Right
	{0, 1},  // Right
	{1, 1},  // Down-Right
	{1, 0},  // Down
	{1, -1}, // Down-Left
	{0, -1}, // Left
	{-1, -1}, // Up-Left
}

var X_MAS_directions = [][2][2]int{
	{{-1, -1}, {1, 1}}, // Top-left and Bottom-right
	{{-1, 1}, {1, -1}}, // Top-right and Bottom-left
}

func loadData(filepath string) ([]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n"), nil
}

func Task4(filepath string) (int, int, error) {
	matrix, err := loadData(filepath)
	if err != nil {
		return 0, 0, err
	}
	rows := len(matrix)
	cols := len(matrix[0])
	
	res_1 := 0
	res_2 := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 'X' {
				res_1 += countMAS(matrix, i, j, rows, cols)
			}

			if matrix[i][j] == 'A' {
				if checkX_MAS(matrix, i, j, rows, cols) {
					res_2++
				}
			}
		}
	}
	return res_1, res_2, nil
}

func countMAS(matrix []string, x, y, rows, cols int) int {
	count := 0
	for _, dir := range XMAS_directions {
		if checkMAS(matrix, x, y, dir, rows, cols) {
			count++
		}
	}
	return count
}

func checkX_MAS(matrix []string, x, y, rows, cols int) bool {
	for _, pair := range X_MAS_directions {

		corner1X, corner1Y := x+pair[0][0], y+pair[0][1] 
		corner2X, corner2Y := x+pair[1][0], y+pair[1][1] 

		if corner1X < 0 || corner1X >= rows || corner1Y < 0 || corner1Y >= cols ||
			corner2X < 0 || corner2X >= rows || corner2Y < 0 || corner2Y >= cols {
			return false
		}

		if (matrix[corner1X][corner1Y] == 'M' && matrix[corner2X][corner2Y] == 'S') ||
			(matrix[corner1X][corner1Y] == 'S' && matrix[corner2X][corner2Y] == 'M') {
			continue
		} else {
			return false
		}
	}
	return true
}

func checkMAS(matrix []string, x, y int, dir [2]int, rows, cols int) bool {
	target := "MAS"
	for k := 1; k <= len(target); k++ {
		nx, ny := x+k*dir[0], y+k*dir[1]
		if nx < 0 || nx >= rows || ny < 0 || ny >= cols || matrix[nx][ny] != target[k-1] {
			return false
		}
	}
	return true
}
