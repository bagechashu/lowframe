package user

import (
	"net/http"

	userModel "lowframe/model/user"

	"github.com/unrolled/render"
)

func getDataHandler(TmplRender *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		TmplRender.JSON(w, http.StatusOK, userModel.UList)
	}
}

// notImplemented replies to the request with an HTTP 501 Not Implemented.
func notImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}

// notImplementedHandler returns a simple request handler
// that replies to each request with a ``501 Not Implemented'' reply.
func notImplementedHandler() http.HandlerFunc { return http.HandlerFunc(notImplemented) }