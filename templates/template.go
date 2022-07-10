package templates

import "embed"

//go:embed templates
var TemplatesFS embed.FS

//go:embed assets
var AssetsFS embed.FS
