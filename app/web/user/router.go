package user

import (
	"lowframe/app"
	"lowframe/templates"
	"net/http"
)

var UserRList []app.R
var UserRouter *app.SubRoute

func init() {
	UserRouter = &app.SubRoute{
		PathPrefix: "/user",
		RList:      UserRList,
	}

	UserRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodGet,
			Path:       "/info",
			Handler:    getUserInfoHandler(templates.TmplRender),
		},
	)

	UserRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodPost,
			Path:       "/info",
			Handler:    postUserInfoHandler(templates.TmplRender),
		},
	)

}
