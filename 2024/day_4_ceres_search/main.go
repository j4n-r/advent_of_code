package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	input := string(readFile)
	matrix := initMatrix(input)

	for x := range matrix {
		for y := range matrix[x] {
			if matrix[x][y] == "X" {
				checkNN(matrix, x, y, 0)

			}
			fmt.Print(matrix[x][y])
		}
	}

}

func checkNN(matrix [][]string, x int, y int, charPos int) bool {
	row := len(matrix)
	col := len(matrix)
	sstr := strings.Split("XMAS", "")
	if matrix[x][y] != sstr[charPos] {
		return false
	}
	if charPos == 4 {
		return true
	}
	inRange := func(x, y int) bool {
		return x <= row-1 && y <= col-1 && x > 0 && y > 0
	}

	switch {
	case inRange(x, y-1):
		checkNN(matrix, x, y-1, charPos+1)
	case matrix[x][y+1]:
		checkNN(matrix, x, y+1, charPos+1)
	case matrix[x+1][y]:
		checkNN(matrix, x+1, y, charPos+1)
	case matrix[x+1][y+1]:
		checkNN(matrix, x+1, y+1, charPos+1)
	case matrix[x+1][y-1]:
		checkNN(matrix, x+1, y-1, charPos+1)
	case matrix[x-1][y]:
		checkNN(matrix, x-1, y, charPos+1)
	case matrix[x-1][y+1]:
		checkNN(matrix, x-1, y+1, charPos+1)
	case matrix[x-1][y-1]:
		checkNN(matrix, x-1, y-1, charPos+1)
	default:
		return false
	}
	return false
}

func initMatrix(input string) [][]string {
	lines := strings.Split(input, "\n")
	matrix := make([][]string, len(lines))
	for i, l := range lines {
		chars := strings.Split(l, "")
		matrix[i] = make([]string, len(chars))

		for j, c := range chars {
			matrix[i][j] = c
		}
	}
	return matrix
}
