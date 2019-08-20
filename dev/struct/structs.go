package main

import "fmt"

// 一个结构体（struct）就是一组字段（field）

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, 4}
	v.X = 8
	fmt.Println(v.X)
	fmt.Println(Vertex{1, 2})
	v1 := Vertex{1, 2}
	p := &v1
	p.X = 1e9
	fmt.Println(v1)
}
