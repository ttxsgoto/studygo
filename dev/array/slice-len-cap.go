package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // 长度和容量
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // 6
	s = s[:0]     // [] 6
	printSlice(s)

	s = s[:4]     // [2,3,5,7]
	printSlice(s) // 6

	s = s[2:] // [5, 7] 4
	printSlice(s)

	var x []int
	fmt.Println(x, len(x), cap(x))
	if x == nil {
		fmt.Println("nil!")
	}

}
