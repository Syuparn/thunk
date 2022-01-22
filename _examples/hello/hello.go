package main

import "fmt"

type Hello interface {
	Greet(person Person)
}

type myHello struct{}

func (h *myHello) Greet(person Person) {
	fmt.Printf("Hello, %s!\n", person)
}

func NewHello() Hello {
	return &myHello{}
}
