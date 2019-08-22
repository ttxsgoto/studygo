package main

import "fmt"

// 哈希表是一种巧妙并且实用的数据结构   map[k]v, k都有相同的类型， 所有的value也有着相同的类型
// 内置的make函数可以创建一个map； ages := make(map[string]int)
// 映射将键映射到值 ; 映射的零值为 nil

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var m1 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	ages := map[string]int{
		"alice": 34,
		"boss":  35,
	}

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		43.23232, -982.902,
	}
	fmt.Println(m["Bell Labs"], m1, ages)
}
