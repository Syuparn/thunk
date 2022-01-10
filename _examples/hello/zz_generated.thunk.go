// Code generated by thunk; DO NOT EDIT.
package main

type LazyHello struct {
	inner Hello
}

func (l *LazyHello) Greet(person string) func() {
	return func() {
		l.inner.Greet(person)
	}
}

func NewLazyHello(inner Hello) *LazyHello {
	return &LazyHello{inner: inner}
}