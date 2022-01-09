package withmethod

type WithMethod interface {
	Foo(n int) string
	Bar(f float64, bs []byte) (map[string]string, error)
}
