package templates

import "embed"

//go:embed *.tmpl *.yaml
var FS embed.FS
