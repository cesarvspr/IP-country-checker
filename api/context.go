package api

import (
	"net/http"

	"github.com/cesarvspr/avoxi-ip-service/app"
)

type Context struct {
	App  *app.App
	Path string
}

type handler struct {
	app        *app.App
	handleFunc func(*Context, http.ResponseWriter, *http.Request)
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{}
	c.App = h.app

	w.Header().Set("Content-Type", "application/json")

	h.handleFunc(c, w, r)

}

func (api *API) Public(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		app:        api.App,
		handleFunc: h,
	}
}
