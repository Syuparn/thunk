package withthirdimport

import (
	"golang.org/x/xerrors"

	"github.com/gostaticanalysis/knife"
)

type A interface {
	Foo(node knife.ASTNode) xerrors.Formatter
}
