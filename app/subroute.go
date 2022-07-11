package app

import (
	"net/http"
)

type R struct {
	HttpMethod string
	Path       string
	Handler    func(http.ResponseWriter, *http.Request)
}

type SubRoute struct {
	PathPrefix string
	RList      []R
}

func (router *SubRoute) AddRouter(r R) {
	router.RList = append(router.RList, r)
}
