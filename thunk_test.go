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
		// TODO: with structs
		// TODO: without parameter name
		// TODO: variadic parameter
		// TODO: multiple interfaces
		// TODO: multiple files
		// TODO: smallcase (unexported) interfaces
		// TODO: imports
	}

	for _, tt := range tests {
		tt := tt // pin

		t.Run(tt.title, func(t *testing.T) {
			rs := codegentest.Run(t, codegentest.TestData(), Generator, tt.dirName)
			codegentest.Golden(t, rs, flagUpdate)
		})
	}
}
