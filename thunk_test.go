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
			title:   "hello world (used for readme introduction)",
			dirName: "hello",
		},
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
		{
			title:   "multiple source files",
			dirName: "multifile",
		},
		{
			title:   "method signature without parameter name",
			dirName: "noname",
		},
		{
			title:   "method signature with composite types",
			dirName: "composite",
		},
		{
			title:   "signature with imported types",
			dirName: "withimport",
		},
		/*
			// TODO: run test once the issue is resolved
			// https://github.com/golang/go/issues/46041
			{
				title:   "signature with 3rd-party imported types",
				dirName: "withthirdimport",
			},
		*/
		{
			title:   "method signature with types defined in the same package",
			dirName: "definedtype",
		},
		{
			title:   "method signature with types defined in the same package (module has slashes)",
			dirName: "definedtypeslash",
		},
		{
			title:   "main package",
			dirName: "inmain",
		},
		// TODO: smallcase (unexported) interfaces
		// TODO: import alias
	}

	for _, tt := range tests {
		tt := tt // pin

		t.Run(tt.title, func(t *testing.T) {
			rs := codegentest.Run(t, codegentest.TestData(), Generator, tt.dirName)
			codegentest.Golden(t, rs, flagUpdate)
		})
	}
}
