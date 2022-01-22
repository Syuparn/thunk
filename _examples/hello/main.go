package main

func main() {
	tom := Person("Tom")
	bob := Person("Bob")

	// original interface `Hello`
	hello := NewHello()
	hello.Greet(tom)

	// generated wrapper `LazyHello`
	lazyHello := NewLazyHello(hello)
	greetThunk := lazyHello.Greet(bob) // Hello.Greet is not evaluated yet!
	greetThunk()                       // Hello.Greet is evaluated here
}
