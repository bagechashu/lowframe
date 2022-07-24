package index

import (
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
