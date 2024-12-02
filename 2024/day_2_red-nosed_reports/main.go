package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	//readFile, err := os.Open("sample.txt")
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
		if report[0] == report[1] {
			continue
		}
		if checkDecr(report) {
			save++
		} else if checkIncr(report) {
			save++
		} else {
			continue
		}
	}
	fmt.Println(save)
}

func checkIncr(report []int) bool {
	for i := 0; i < len(report); i++ {
		if i+1 >= len(report) {
			fmt.Println(report)
			break
		}
		if report[i] == report[i+1] {
			return false
		}
		if (report[i]-report[i+1]) < -3 || (report[i]-report[i+1]) > -1 {
			return false
		}
	}
	return true
}
func checkDecr(report []int) bool {
	for i := 0; i < len(report); i++ {
		if i+1 >= len(report) {
			fmt.Println(report)
			break
		}
		if report[i] == report[i+1] {
			return false
		}
		if (report[i]-report[i+1]) > 3 || (report[i]-report[i+1]) < 1 {
			return false
		}
	}
	return true
}
