package templates

import (
	"embed"

	"github.com/unrolled/render"
)

//go:embed templates
var TemplatesFS embed.FS

//go:embed assets
var AssetsFS embed.FS

var LayoutHTMLRender = render.New(render.Options{
	Directory: "templates",
	FileSystem: &render.EmbedFileSystem{
		FS: TemplatesFS,
	},
	DisableHTTPErrorRendering: true,
	Layout:                    "layout/layout",
	Extensions:                []string{".html"},
	IndentJSON:                true,
})

var RawRender = render.New(render.Options{
	Directory: "templates",
	FileSystem: &render.EmbedFileSystem{
		FS: TemplatesFS,
	},
	DisableHTTPErrorRendering: true,
	Extensions:                []string{".html"},
	IndentJSON:                true,
})
