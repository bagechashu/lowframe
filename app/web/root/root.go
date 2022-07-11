package root

import (
	"lowframe/app"
	"lowframe/templates"
	"net/http"

	"github.com/unrolled/render"
)

func homeHandler(TmplRender *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		TmplRender.HTML(w, http.StatusOK, "user/index", nil)
	}
}

var rootRList []app.R
var RootRouter *app.SubRoute

func init() {
	RootRouter = &app.SubRoute{
		PathPrefix: "/",
		RList:      rootRList,
	}

	RootRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodGet,
			Path:       "/",
			Handler:    homeHandler(templates.TmplRender),
		},
	)
}
