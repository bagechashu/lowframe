package user

import (
	"lowframe/app"
	"lowframe/templates"
	"net/http"
)

var apiUserRList []app.R
var ApiUserRouter *app.SubRoute

func init() {
	ApiUserRouter = &app.SubRoute{
		PathPrefix: "/api/v1/user",
		RList:      apiUserRList,
	}

	ApiUserRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodGet,
			Path:       "/getData",
			Handler:    getDataHandler(templates.RawRender),
		},
	)
	ApiUserRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodGet,
			Path:       "/info",
			Handler:    getUserInfoHandler(templates.RawRender),
		},
	)

	ApiUserRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodPost,
			Path:       "/login",
			Handler:    postUserLoginHandler(templates.RawRender),
		},
	)

}
