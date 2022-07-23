package root

import (
	"lowframe/app"
	"lowframe/templates"
	"net/http"

	"github.com/unrolled/render"
)

// func redirectToUser() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		http.Redirect(w, r, "/user/", http.StatusMovedPermanently)
// 	}
// }

func indexHandler(layoutHTMLRender *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		layoutHTMLRender.HTML(w, http.StatusOK, "index", nil)
	}
}

var rootRList []app.R
var RootRouter *app.SubRoute

func init() {
	RootRouter = &app.SubRoute{
		PathPrefix: "",
		RList:      rootRList,
	}

	RootRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodGet,
			Path:       "/",
			Handler:    indexHandler(templates.LayoutHTMLRender),
		},
	)
}
