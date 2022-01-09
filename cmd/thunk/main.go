package main

import (
	"github.com/gostaticanalysis/codegen/singlegenerator"
	"github.com/syuparn/thunk"
)

func main() {
	singlegenerator.Main(thunk.Generator)
}
