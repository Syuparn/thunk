package variadic

type Variadic interface {
	Foo(n int, options ...string) error
}
