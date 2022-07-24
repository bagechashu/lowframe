package index

import (
	"lowframe/app"
	"lowframe/app/public"
	"net/http"
)

var apiIndexRList []app.R
var ApiIndexRouter *app.SubRoute

func init() {
	ApiIndexRouter = &app.SubRoute{
		PathPrefix: "/api/v1",
		RList:      apiIndexRList,
	}

	ApiIndexRouter.AddRouter(
		app.R{
			HttpMethod: http.MethodGet,
			Path:       "/unknown",
			Handler:    public.NotImplementedHandler(),
		},
	)

}
