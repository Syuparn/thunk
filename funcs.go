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

	// trim prefix if type is defined in pkg
	if strings.HasPrefix(typ, pkg.Path) {
		return strings.TrimPrefix(typ, pkg.Path+".")
	}

	// convert type path into type name if it is defined in imported packages
	// NOTE: this is neccessary to remove `/`
	for _, p := range pkg.Imports {
		if strings.HasPrefix(typ, p.Path) {
			return strings.Replace(typ, p.Path, p.Name, 1)
		}
	}

	// TODO: handle import alias

	return typ
}
