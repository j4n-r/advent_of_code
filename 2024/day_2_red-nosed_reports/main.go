package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//readFile, err := os.Open("input.txt")
	readFile, err := os.Open("sample.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	save := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		report := []int{}
		reportStr := strings.Fields(line)
		for _, v := range reportStr {
			v, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}
			report = append(report, v)
		}
		for i, v := range report {

		}

		fmt.Println(report)
	}
	fmt.Println(save)
}
func checkValid(prev, curr, next int) bool {
	if prev == curr || curr == next {
		return false
	}
	if (prev-curr) > 3 && (curr-prev) > 3 {
		return false
	}
	if (prev-curr) < -3 && (curr-prev) < -3 {
		return false
	}
	return true

}
