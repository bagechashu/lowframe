package index

import (
	"lowframe/app"
	"lowframe/templates"
	"net/http"
)

var indexRList []app.R
var IndexRouter *app.SubRoute

func init() {
	IndexRouter = &app.SubRoute{
		PathPrefix: "",
		RList:      indexRList,
	}

	IndexRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodGet,
			Path:       "/",
			Handler:    indexHandler(templates.LayoutHTMLRender),
		},
	)
}
