package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	stringData := string(data)

	r, _ := regexp.Compile("mul\\(\\d{1,3},\\s?\\d{1,3}\\)")
	fmt.Println(r.MatchString(stringData))
	allMatches := r.FindAllString(stringData, -1)

	joinedMatches := strings.Join(allMatches[:], "")
	nr, _ := regexp.Compile("\\d{1,3}")

	fmt.Println(joinedMatches)
	allNums := nr.FindAllString(joinedMatches, -1)

	sum := 0
	operand1 := 0
	operand2 := 0
	for i, v := range allNums {
		num, err := strconv.Atoi(v)
		fmt.Println(operand1, operand2, sum)
		if err != nil {
			return
		}
		if i%2 == 1 {
			operand1 = num
			sum += operand1 * operand2
		} else {
			operand2 = num
		}
	}
	fmt.Println(allNums)
	fmt.Println(sum)

}
