package user

import (
	"lowframe/app"
	"lowframe/templates"
	"net/http"
)

var userRList []app.R
var UserRouter *app.SubRoute

func init() {
	UserRouter = &app.SubRoute{
		PathPrefix: "/user",
		RList:      userRList,
	}

	UserRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodGet,
			Path:       "/login",
			Handler:    userLoginFormHandler(templates.LayoutHTMLRender),
		},
	)

}
