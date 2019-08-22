package main

import "fmt"

// 数组是长度固定的数据类型，必须存储一段相同类型的元素，而且这些元素是连续的, 可以很快的索引数组中的任何数据

// 同样类型的数组是可以相互赋值的，不同类型的不行，会编译错误
// 长度一样并且每个元素的类型也一样的数组才是同样类型的数组

func main() {
	var a [2]string //类型 [n]T 表示拥有 n 个 T 类型的值的数组
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	for i, v := range primes {
		fmt.Printf("%d %d\n", i, v)
		fmt.Println(v)
	}

	var s []int = primes[0:4] // 切片， 切片并不存储任何数据，它只是描述了底层数组中的一段
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)

	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s1 := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s1)

	q1 := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q1)

	type Currency int
	const (
		USD Currency = iota
		RMB
	)

	symbol := [...]string{USD: "$", RMB: "%"}
	fmt.Println(symbol[RMB])

	array := [5]*int{1: new(int), 3: new(int)}
	*array[1] = 1
	array[0] = new(int)
	*array[0] = 3
	fmt.Println(array[0], *array[0])

}
