package main

import (
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}