// package main

// import "fmt"

// func max(num1, num2 int) int {
// 	var result int

// 	if num1 > num2 {
// 		result = num1
// 	} else {
// 		result = num2
// 	}
// 	return result
// }

// func main() {
// 	a := 100
// 	b := 200s
// 	var ret int

// 	ret = max(a, b)

// 	fmt.Println("max value: \n", ret)
// }

package main

import "fmt"

// 函数是指不属于任何结构体、类型的方法，也就是说，函数是没有接收者的
// 方法是有接收者的，方法要么是属于一个结构体，或者属于一个新定义的类型的，
// 方法在定义的时候 会在func和方法名之间添加一个参数，这个参数就是接收者
// 使用值类型接收者定义的方法，在调用的时候，使用的其实是值接收者的一个副本，
// 所以对该值的任何操作，不会影响原来的类型变量。

func add1(a *int) int { // 修改为指针类型的传递方式
	*a = *a + 1
	return *a
}

func main() {
	x := 10
	if x > 2 {
		fmt.Println("x is more than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	fmt.Println("Hello world!")
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is ", sum)

	interger := 4
	switch interger {
	case 4:
		fmt.Println("4")
		fallthrough
	case 5:
		fmt.Println("5")
		fallthrough
	case 6:
		fmt.Println("6")
	default:
		fmt.Printf("default")
	}

	fmt.Println("x = ", x)
	x1 := add1(&x) // 调用add1的时候，add1接收的参数其实是x的copy，而不是x本身, &x 传x的地址
	// 变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的内存
	fmt.Println("x + 1 = ", x1)
	fmt.Println("x = ", x)

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

}
