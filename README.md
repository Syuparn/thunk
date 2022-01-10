# thunk
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

`thunk` is a code generator to make interface's wrapper with methods evaluated lazily.

# About

`thunk` generates interface's wrapper for lazy evaluation.
Methods of the wrapper returns a *thunk* of the method call, instead of evaluating indivisually.

```go
// original interface `Hello`
hello := NewHello()
hello.Greet("Tom")

// generated wrapper `LazyHello`
lazyHello := NewLazyHello(hello)
greetThunk := lazyHello.Greet("Tom") // Hello.Greet is not evaluated yet!
greetThunk()                         // Hello.Greet is evaluated here
```

See `/testdata/src/hello` directory for details.

# Usage

(under construction...)

# Note

This project uses [skeleton](https://github.com/gostaticanalysis/skeleton) to generate the module.
