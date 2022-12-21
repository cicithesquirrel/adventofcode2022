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
	maxAmount := 0
	amounts := make([]int, 1)
	currentElf := 0
	for fileScanner.Scan() {
		s := fileScanner.Text()
		if s == "" {
			if amounts[currentElf] > maxAmount {
				maxAmount = amounts[currentElf]
			}
			currentElf++
			amounts = append(amounts, 0)
		} else {
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			amounts[currentElf] += v
		}
	}
	fmt.Println(maxAmount)
}
