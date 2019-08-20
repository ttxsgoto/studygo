package main

import (
	"fmt"
	"math"
	"math/rand"
)

var c, python, java bool

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	ddd := "abc" // 只能在函数内使用， 不能在函数外定义
	fmt.Println("My number is", rand.Intn(10), math.Sqrt(7), math.Pi, add(3, 4))
	a, b := swap("ttxs", "goto")
	fmt.Println(a, b, ddd)
}
