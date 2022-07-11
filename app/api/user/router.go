package user

import (
	"lowframe/app"
	"lowframe/templates"
	"net/http"
)

var ApiUserRList []app.R
var ApiUserRouter *app.SubRoute

func init() {
	ApiUserRouter = &app.SubRoute{
		PathPrefix: "/api/v1/user",
		RList:      ApiUserRList,
	}

	ApiUserRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodGet,
			Path:       "/unknown",
			Handler:    notImplementedHandler(),
		},
	)

	ApiUserRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodGet,
			Path:       "/getData",
			Handler:    getDataHandler(templates.TmplRender),
		},
	)
}
