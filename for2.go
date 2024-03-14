package main

import (
	"fmt"
)

func main() {
	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum = 0
	for _, value := range numbers {
		if value > 4 {
			break
		}
		sum += value
	}
	fmt.Println(sum)
}
