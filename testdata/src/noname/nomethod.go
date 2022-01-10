package noname

type A interface {
	Foo(string, int) error
	// NOTE: you don't concern about below because it is not allowed
	// Hoge(n int, float64) error // mixed named and unnamed parameters
}
