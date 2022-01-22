package thunk

import (
	"strings"
	"text/template"

	"github.com/gostaticanalysis/knife"
)

// extra funcs are added to knife default funcs to help generating source codes
var extraFuncs = template.FuncMap(map[string]interface{}{
	"prettyType": prettyType,
})

// prettyType converts type as string printed on the source code.
func prettyType(pkg *knife.Package, t *knife.Type) string {
	typ := t.String()

	// types defined in same package do not need prefix
	// NOTE: all appeared path pattern should be replaced for compound types (like maps and slices)
	typ = strings.Replace(typ, pkg.Path+".", "", -1)

	// convert type path into type name if it is defined in imported packages
	// this is neccessary to remove `/`
	for _, p := range pkg.Imports {
		// NOTE: pattern should include "." not to replace independent packages
		// i.e) if p.Path is simply replaced with p.Name,
		// `foo/bar/hoge/piyo.myStruct` might be replaced with `hoge/piyo.myStruct`,
		// as `foo/bar/hoge` is also imported.
		typ = strings.Replace(typ, p.Path+".", p.Name+".", -1)
	}

	// TODO: handle import alias

	return typ
}
