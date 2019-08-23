package main

type Link struct {
	data interface {
	}
	next *Link
}


func(p *Link) AddNode(data interface()) {
	node :=&Link(
		data: data,
		next: nil
	)
}


