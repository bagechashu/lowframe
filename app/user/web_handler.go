package user

import (
	"net/http"

	"github.com/unrolled/render"
)

func userLoginFormHandler(layoutHTMLRender *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		layoutHTMLRender.HTML(w, http.StatusOK, "user/login", nil)
	}
}
