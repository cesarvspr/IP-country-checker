package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cesarvspr/avoxi-ip-service/app"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

func (a *API) InitRoutes(r *mux.Router) {
	route := a.BaseRoutes.APIRoot.Handle
	public := a.Public

	route("/login", public(loginHandler)).Methods("GET")
	log.Info("[INFO] InitRoutes")

	// utils.PrintEndpoints(r)
}

func loginHandler(c *Context, w http.ResponseWriter, r *http.Request) {

	vals := r.URL.Query()
	ip := vals["ip"]

	whitelist := vals["whitelist"]
	if len(whitelist) < 1 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "empty whitelist: %v\n", vals["whitelist"])
		return // no whitelist on payload
	}

	// parse whitelist
	whitelist[0] = strings.Replace(whitelist[0], " ", "", -1)
	whitelist = strings.Split(whitelist[0], ",")

	// check if ip is in whitelist
	res := c.App.LoginAPI(ip[0], whitelist, r)
	if res {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusExpectationFailed)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "route not found. \n")

}

//API - struct
type API struct {
	App        *app.App
	BaseRoutes *Routes
}

//Routes - struct
type Routes struct {
	Root    *mux.Router
	APIRoot *mux.Router
}
