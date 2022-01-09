package template

import (
	"embed"
)

// FS exports current directory file system.
//
//go:embed *.tmpl
var FS embed.FS
