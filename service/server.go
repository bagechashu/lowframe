package service

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"lowframe/templates"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"

	"github.com/liushuochen/gotable"
)

type Server struct {
	Router  *mux.Router
	Negroni *negroni.Negroni
}

// NewServer configures and returns a Server.
func (s *Server) initServer() {
	s.Router = mux.NewRouter()
	s.Negroni = negroni.Classic()
}

func (s *Server) initRoutes() {
	tmplRender := render.New(render.Options{
		Directory: "templates",
		FileSystem: &render.EmbedFileSystem{
			FS: templates.TemplatesFS,
		},
		DisableHTTPErrorRendering: true,
		Layout:                    "layout/layout",
		Extensions:                []string{".html"},
		IndentJSON:                true,
	})

	s.Router.HandleFunc("/api/unknown", notImplementedHandler()).Methods("GET")
	s.Router.HandleFunc("/api/getData", getDataHandler(tmplRender)).Methods("GET")
	s.Router.HandleFunc("/", homeHandler(tmplRender)).Methods("GET")
	s.Router.HandleFunc("/userInfo", postUserInfoHandler(tmplRender)).Methods("POST")
	s.Router.HandleFunc("/userInfo", getUserInfoHandler(tmplRender)).Methods("GET")

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
	s.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.FS(assets)))).Methods(http.MethodGet)

}

func (s *Server) Run(addr ...string) {
	s.initServer()
	s.initRoutes()
	s.Negroni.UseHandler(s.Router)
	s.Negroni.Run(addr...)
}

func (s *Server) WalkRoutes() {
	s.initServer()
	s.initRoutes()

	table, err := gotable.Create(
		"ROUTE",
		"Path regexp",
		"Queries templates",
		"Queries regexps",
		"Methods",
	)
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	table.Align("ROUTE", 1)
	table.Align("Path regexp", 1)
	table.Align("Queries templates", 1)
	table.Align("Queries regexps", 1)
	table.Align("Methods", 1)

	err = s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		pathRegexp, err := route.GetPathRegexp()
		if err != nil {
			return err
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err != nil {
			return err
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err != nil {
			return err
		}
		methods, err := route.GetMethods()
		if err != nil {
			fmt.Println(err)
		}
		if methods == nil {
			methods = []string{"Null"}
		}

		table.AddRow([]string{
			pathTemplate,
			pathRegexp,
			strings.Join(queriesTemplates, ","),
			strings.Join(queriesRegexps, ","),
			strings.Join(methods, ","),
		})

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(table)
}
