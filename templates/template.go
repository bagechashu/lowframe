package templates

import (
	"embed"

	"github.com/unrolled/render"
)

//go:embed templates
var TemplatesFS embed.FS

//go:embed assets
var AssetsFS embed.FS

var TmplRender = render.New(render.Options{
	Directory: "templates",
	FileSystem: &render.EmbedFileSystem{
		FS: TemplatesFS,
	},
	DisableHTTPErrorRendering: true,
	Layout:                    "layout/layout",
	Extensions:                []string{".html"},
	IndentJSON:                true,
})
