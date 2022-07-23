package user

import (
	"lowframe/app"
	"net/http"
)

var ApiUserRList []app.R
var ApiUserRouter *app.SubRoute

func init() {
	ApiUserRouter = &app.SubRoute{
		PathPrefix: "/api/v1",
		RList:      ApiUserRList,
	}

	ApiUserRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodGet,
			Path:       "/unknown",
			Handler:    notImplementedHandler(),
		},
	)

}
