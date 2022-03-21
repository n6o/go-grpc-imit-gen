package templates

import _ "embed"

//go:embed server.tmpl
var DefaultServerTemplate []byte
