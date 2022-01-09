package thunk

import (
	"bytes"
	"fmt"
	"go/format"
	"go/types"
	"os"

	"golang.org/x/xerrors"

	"github.com/gostaticanalysis/analysisutil"
	"github.com/gostaticanalysis/codegen"
	"github.com/gostaticanalysis/knife"
)

const doc = "thunk is a code generator to make interface's wrapper with methods evaluated lazily."

var (
	flagOutput string
)

func init() {
	Generator.Flags.StringVar(&flagOutput, "o", "", "output file name")
}

var Generator = &codegen.Generator{
	Name: "thunk",
	Doc:  doc,
	Run:  run,
}

func run(pass *codegen.Pass) error {
	ifaces := map[string]*knife.Interface{}

	s := pass.Pkg.Scope()
	for _, name := range s.Names() {
		obj := s.Lookup(name)
		if !obj.Exported() {
			continue
		}
		iface, _ := analysisutil.Under(obj.Type()).(*types.Interface)
		if iface != nil {
			ifaces[name] = knife.NewInterface(iface)
		}
	}

	td := &knife.TempalteData{
		Fset:      pass.Fset,
		Files:     pass.Files,
		TypesInfo: pass.TypesInfo,
		Pkg:       pass.Pkg,
	}

	tmpl, err := readTemplate("thunk")
	if err != nil {
		return xerrors.Errorf("unable to read template: %w", err)
	}

	t, err := knife.NewTemplate(td).Parse(tmpl)
	if err != nil {
		return xerrors.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, ifaces); err != nil {
		return xerrors.Errorf("failed to execute template: %w", err)
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return xerrors.Errorf("failed to format generated source: %w", err)
	}

	if flagOutput == "" {
		pass.Print(string(src))
		return nil
	}

	f, err := os.Create(flagOutput)
	if err != nil {
		return xerrors.Errorf("failed to create generated file: %w", err)
	}

	fmt.Fprint(f, string(src))

	if err := f.Close(); err != nil {
		return xerrors.Errorf("failed to close generated file: %w", err)
	}

	return nil
}
