package thunk

import (
	"flag"
	"os"
	"testing"

	"github.com/gostaticanalysis/codegen/codegentest"
)

var flagUpdate bool

func TestMain(m *testing.M) {
	flag.BoolVar(&flagUpdate, "update", false, "update the golden files")
	flag.Parse()
	os.Exit(m.Run())
}

func TestGenerator(t *testing.T) {
	tests := []struct {
		title   string
		dirName string
	}{
		{
			title:   "interface without any methods",
			dirName: "nomethod",
		},
		{
			title:   "interface with some methods",
			dirName: "withmethod",
		},
		{
			title:   "method signature with a variadic parameter",
			dirName: "variadic",
		},
		{
			title:   "source file with multiple interfaces",
			dirName: "multiinterface",
		},
		/*
			// TODO: add test once package prefix can be trimmed
			{
				title:   "method signature with types defined in the same package",
				dirName: "definedtype",
			},
		*/
		// TODO: without parameter name
		// TODO: multiple files
		// TODO: smallcase (unexported) interfaces
		// TODO: imports
		// TODO: import alias
		// TODO: hello world for example
	}

	for _, tt := range tests {
		tt := tt // pin

		t.Run(tt.title, func(t *testing.T) {
			rs := codegentest.Run(t, codegentest.TestData(), Generator, tt.dirName)
			codegentest.Golden(t, rs, flagUpdate)
		})
	}
}
