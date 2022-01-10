package main

import "fmt"

type Hello interface {
	Greet(person string)
}

type myHello struct{}

func (h *myHello) Greet(person string) {
	fmt.Printf("Hello, %s!\n", person)
}

func NewHello() Hello {
	return &myHello{}
}
