# thunk
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Test](https://github.com/Syuparn/thunk/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/Syuparn/thunk/actions/workflows/test.yml)

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
greetThunk := lazyHello.Greet("Bob") // Hello.Greet is not evaluated yet!
greetThunk()                         // Hello.Greet is evaluated here
```

See [/_examples/hello](https://github.com/Syuparn/thunk/tree/main/_examples/hello) for details.

# Install

```bash
$ go install github.com/syuparn/thunk/cmd/thunk@latest
```

# Usage

```bash
# generate wrapper code to zz_generated.thunk.go
$ thunk -o zz_generated.thunk.go example.com/your/package
```

Also, you can use `go generate` to generate thunks.

```go
//go:generate thunk -o zz_generated.thunk.go example.com/your/package
```

# Note

This project uses [skeleton](https://github.com/gostaticanalysis/skeleton) to generate the module.
