package main

import (
	"fmt"
	"reflect"
)

// reflect.Value.NumField()获取结构体中字段的个数
// reflect.Value.Method(n).Call来调用结构体中的方法

type Student struct {
	Name  string
	Age   int
	Score float32
}

func (s Student) Set(name string, age int, sorce float32) {
	s.Name = name
	s.Age = age
	s.Score = sorce
}

func TestStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.NumField()   // 字段个数
	meth := val.NumMethod() // 方法个数
	fmt.Println(num, meth)

}

func main() {
	var a Student = Student{
		Name:  "ttxs",
		Age:   23,
		Score: 90.0,
	}
	// var a int = 23

	TestStruct(a)

}
