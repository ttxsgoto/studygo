package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	sum := 1
	// for i := 0; i < 10; i++ {
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum, sqrt(3), sqrt(-3))
}
