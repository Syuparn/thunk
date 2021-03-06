package thunk

import (
	"bytes"
	"fmt"
	"go/format"
	"go/types"
	"os"

	"golang.org/x/tools/imports"
	"golang.org/x/xerrors"

	log "github.com/sirupsen/logrus"

	"github.com/gostaticanalysis/analysisutil"
	"github.com/gostaticanalysis/codegen"
	"github.com/gostaticanalysis/knife"
)

const doc = "thunk is a code generator to make interface's wrapper with methods evaluated lazily."

var Generator = &codegen.Generator{
	Name: "thunk",
	Doc:  doc,
	Run:  run,
}

func run(pass *codegen.Pass) error {
	initDebugLog()

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
	log.Debugf("read template:\n%s\n---", tmpl)

	t, err := knife.NewTemplate(td).Funcs(extraFuncs).Parse(tmpl)
	if err != nil {
		return xerrors.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, ifaces); err != nil {
		return xerrors.Errorf("failed to execute template: %w", err)
	}
	log.Debugf("generated raw soruce code:\n%s\n---", string(buf.Bytes()))

	// format codes like `go fmt`
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return xerrors.Errorf("failed to format generated source: %w", err)
	}
	log.Debugf("go-fmted soruce code:\n%s\n---", string(src))

	// remove unused imports
	src, err = imports.Process(flagOutput, src, nil /* options */)
	if err != nil {
		return xerrors.Errorf("failed to remove unused imports: %w", err)
	}
	log.Debugf("output soruce code:\n%s\n---", string(src))

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
