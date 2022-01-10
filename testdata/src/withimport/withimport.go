package withimport

import (
	"context"
	"io"
)

type A interface {
	Foo(ctx context.Context) io.Reader
	// Bar(m json.Marshaler) // TODO: add test for packages with `/`
}
