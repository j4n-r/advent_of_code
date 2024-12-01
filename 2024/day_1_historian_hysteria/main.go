package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	leftIds := []int{}
	rightIds := []int{}

	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		idStrings := strings.Split(line, "  ")

		idLeft, err := strconv.Atoi(strings.TrimSpace(idStrings[0]))
		idRight, err := strconv.Atoi(strings.TrimSpace(idStrings[1]))
		if err != nil {
			fmt.Println(err)
		}

		leftIds = append(leftIds, idLeft)
		rightIds = append(rightIds, idRight)
	}
	slices.Sort(leftIds)
	slices.Sort(rightIds)

	sum := 0
	for i := range leftIds {
		distance := leftIds[i] - rightIds[i]
		distance = int(math.Abs(float64(distance)))
		sum = sum + distance
		// fmt.Printf("%d \t %d \t distance: %d\n", leftIds[i], rightIds[i], distance)
	}

	fmt.Printf("sum: %d ", sum)
	partTwo(leftIds, rightIds)

}

func partTwo(leftIds []int, rightIds []int) {

	smScore := 0
	for _, vl := range leftIds {
		counter := 0
		for _, vr := range rightIds {
			if vl == vr {
				counter++
			}
		}
		smScore += vl * counter
	}
	fmt.Println("smScore: ", smScore)
}
