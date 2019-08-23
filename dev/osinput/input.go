package main

import (
	"fmt"
)

type student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	var str = "ttxsgoto 28 90.0"
	var stu student

	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)

	fmt.Println(stu)
}
