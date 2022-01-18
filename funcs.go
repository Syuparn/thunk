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

	// TODO: handle import alias
	// TODO: handle path slash

	return typ
}
