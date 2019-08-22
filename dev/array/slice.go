package main

import "fmt"

// Slice（切片）代表变长的序列，序列中每个元素都有相同的类型
// Slice 由三部分构成： 指针，长度和容量
// 多个slice之间可以共享底层的数据，并且引用的数组部分区间可能重叠
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// slice 之间不能比较，不能使用==操作符， 一个零值的slice等于nil
	a := [...]int{0, 1, 2, 3, 4, 5} // 没有指定序列的长度，隐式地创建一个合适大小的数组，然后slice的指针指向底层的数组
	reverse(a[:])
	fmt.Println(a)

}
