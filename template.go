package thunk

import (
	"bytes"
	"fmt"

	"golang.org/x/xerrors"

	"github.com/syuparn/thunk/template"
)

func readTemplate(name string) (string, error) {
	filePath := fmt.Sprintf("%s.tmpl", name)

	fp, err := template.FS.Open(filePath)
	if err != nil {
		return "", xerrors.Errorf("failed to open template file %s: %w", filePath, err)
	}
	defer fp.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(fp)
	if err != nil {
		return "", xerrors.Errorf("failed to read template %s: %w", filePath, err)
	}

	return buf.String(), nil
}
