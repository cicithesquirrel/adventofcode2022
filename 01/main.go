package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	maxAmounts := make([]int, 3)
	currentAmount := 0
	for fileScanner.Scan() {
		s := fileScanner.Text()
		if s == "" {
			maxAmounts = replaceSmallestAmount(maxAmounts, currentAmount)
			currentAmount = 0
		} else {
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			currentAmount += v
		}
	}
	fmt.Println(sum(maxAmounts))
}

func replaceSmallestAmount(amounts []int, amount int) []int {
	smallestIndex := 0
	smallest := amounts[0]
	for i := 1; i < len(amounts); i++ {
		if amounts[i] <= smallest {
			smallest = amounts[i]
			smallestIndex = i
		}
	}
	if smallest < amount {
		amounts[smallestIndex] = amount
	}
	return amounts
}

func sum(amounts []int) int {
	sum := 0
	for _, a := range amounts {
		sum += a
	}
	return sum
}
