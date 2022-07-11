package root

import (
	"lowframe/app"
	"net/http"
)

func redirectToUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/user/", http.StatusMovedPermanently)
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
			Handler:    redirectToUser(),
		},
	)
}
