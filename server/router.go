package server

import (
	"io/fs"
	"log"
	"lowframe/app"
	apiUser "lowframe/app/api/user"
	"lowframe/app/web/root"
	"lowframe/app/web/user"
	"lowframe/templates"
	"net/http"

	"github.com/gorilla/mux"
)

func registerRoutes(mux *mux.Router, router *app.SubRoute) {
	subrouter := mux.PathPrefix(router.PathPrefix).Subrouter()

	for _, r := range router.RList {
		subrouter.HandleFunc(r.Path, r.Handler).Methods(r.HttpMethod)
	}
}

func initAppRouter() {
	registerRoutes(Mux, root.RootRouter)
	registerRoutes(Mux, user.UserRouter)
	registerRoutes(Mux, apiUser.ApiUserRouter)
}

func initAssetsRouter() {
	// // setup file server
	// webRoot := os.Getenv("WEBROOT")
	// if len(webRoot) == 0 {
	// 	if root, err := os.Getwd(); err != nil {
	// 		panic("Could not retrive working directory")
	// 	} else {
	// 		webRoot = root
	// 	}
	// }
	// s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))

	assets, err := fs.Sub(templates.AssetsFS, "assets")
	if err != nil {
		log.Fatal(err)
	}
	Mux.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.FS(assets)))).Methods(http.MethodGet)
}
