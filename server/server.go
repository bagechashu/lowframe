package server

import (
	"fmt"
	"strings"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/liushuochen/gotable"
)

var (
	Mux     *mux.Router
	Negroni *negroni.Negroni
)

func init() {
	Mux = mux.NewRouter()
	Negroni = negroni.Classic()
	initAppRouter()
	initAssetsRouter()
}

func Run(addr ...string) {
	Negroni.UseHandler(Mux)
	Negroni.Run(addr...)
}

func WalkRoutes() {
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

	err = Mux.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
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
			fmt.Printf("Route \t%s\t donot define http method.\n", pathTemplate)
		}
		if methods == nil {
			// methods = []string{"Null"}
			return nil
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
