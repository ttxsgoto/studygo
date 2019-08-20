package main

import "fmt"

// ch <- v    // 将 v 发送至信道 ch
// v := <-ch  // 从 ch 接收值并赋予 v;（“箭头”就是数据流的方向。）
// 和映射与切片一样，信道在使用前必须创建： ch := make(chan int)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
