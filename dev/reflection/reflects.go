package main

import "reflect"
import "fmt"

// reflect.TypeOf()   // 获取变量的类型，返回reflect.type类型
// reflect.ValueOf()  // 获取变量的值，返回reflect.value类型
// reflect.Value.Kind	// 获取变量的类别，返回一个常量
// reflect.Value.Interface() 	// 转换成interface{} 类型
type Student struct {
	Name  string
	Age   int
	Score float32
}

func test(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println(t)

	v := reflect.ValueOf(b)
	k := v.Kind() // 类别
	iv := v.Interface()
	fmt.Println(v, k, iv)
	// if ok {
	// fmt.Println(v, k, iv)
	// }
}

func testInt(b interface{}) {

	val := reflect.ValueOf(b)
	val.Elem().SetInt(100) // 得到内存地址 类似于*val操作
	// val.SetInt(100)
	c := val.Elem().Int()
	fmt.Println(val, c) // val 为内存地址， c 为获取到设置的值

}

func main() {
	// var a int = 200
	var a Student = Student{
		Name:  "ttxs",
		Age:   23,
		Score: 90,
	}
	test(a)
	var b int = 1
	testInt(&b)
	fmt.Println(b)

}
