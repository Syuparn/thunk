package withimport

import (
	"context"
	"encoding/json"
	"io"
)

type A interface {
	Foo(ctx context.Context) io.Reader
	// handle packages with `/` (`encoding/json`)
	Bar(m json.Marshaler)
	// compound types
	Baz(m map[json.Unmarshaler]io.Writer)
}
