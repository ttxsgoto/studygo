package main

import (
	"fmt"
	"math"
)

// 方法就是一类带特殊的 接收者 参数的函数
// 方法只是个带接收者参数的函数

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs1(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs1(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs2())
}
