package main

import (
	"fmt"
	"os"
	"strings"
)

var count int = 0

func main() {
	readFile, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	input := string(readFile)
	matrix := initMatrix(input)

	for x := range matrix {
		for y := range matrix[x] {
			checkNN(matrix, x, y, 0)
		}
	}
	fmt.Println("count: ", count)

}

// forgot the direction, if the M is left to the X the A and S also have to be on the right side

func checkNN(matrix [][]string, x int, y int, charPos int) {
	row := len(matrix) - 1
	// fmt.Println("rowlen: ", row)
	col := len(matrix[0])
	// fmt.Println("collen: ", col)
	sstr := strings.Split("XMAS", "")

	fmt.Println(charPos)

	charPos++
	if charPos == 4 {
		count++
		return
	}

	inRange := func(x, y int) bool {
		// fmt.Printf("x: %d, y= %d\n", x, y)
		return x <= row-1 && y <= col-1 && x >= 0 && y >= 0
	}

	switch {
	case inRange(x, y-1) && matrix[x][y-1] == sstr[charPos]:
		checkNN(matrix, x, y-1, charPos)
	case inRange(x, y+1) && matrix[x][y+1] == sstr[charPos]:
		checkNN(matrix, x, y+1, charPos)
	case inRange(x+1, y) && matrix[x+1][y] == sstr[charPos]:
		checkNN(matrix, x+1, y, charPos)
	case inRange(x+1, y+1) && matrix[x+1][y+1] == sstr[charPos]:
		checkNN(matrix, x+1, y+1, charPos)
	case inRange(x+1, y-1) && matrix[x+1][y-1] == sstr[charPos]:
		checkNN(matrix, x+1, y-1, charPos)
	case inRange(x-1, y) && matrix[x-1][y] == sstr[charPos]:
		checkNN(matrix, x-1, y, charPos)
	case inRange(x-1, y+1) && matrix[x-1][y+1] == sstr[charPos]:
		checkNN(matrix, x-1, y+1, charPos)
	case inRange(x-1, y-1) && matrix[x-1][y-1] == sstr[charPos]:
		checkNN(matrix, x-1, y-1, charPos)
	default:
		return
	}
	return
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
