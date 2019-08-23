package main

import "fmt"

func Test(a interface{}) {
	b, ok := a.(int) // 类型断言 a.(type)
	if ok == false {
		fmt.Println("ERROR")
		return
	}
	b += 3
	fmt.Println(b)
}

func main() {
	var a interface{}
	var b int

	a = b
	c := a.(int)
	// var d int
	Test(b)
	fmt.Printf("%d %T\n", c, a)

}
