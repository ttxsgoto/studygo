package main

import "fmt"

func main() {
	var i interface{}
	// 空接口没有任何方法， 所以所有类型都实现了空接口
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
