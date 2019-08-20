package main

import "fmt"

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
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		43.23232, -982.902,
	}
	fmt.Println(m["Bell Labs"], m1)
}
