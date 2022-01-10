package main

func main() {
	// original interface `Hello`
	hello := NewHello()
	hello.Greet("Tom")

	// generated wrapper `LazyHello`
	lazyHello := NewLazyHello(hello)
	greetThunk := lazyHello.Greet("Bob") // Hello.Greet is not evaluated yet!
	greetThunk()                         // Hello.Greet is evaluated here
}
