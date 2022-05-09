package api

import (
	"net"
	"net/http"

	"github.com/cesarvspr/avoxi-ip-service/app"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

//Init - starting endpoints
func Init(a *app.App, root *mux.Router) *API {
	api := &API{
		App:        a,
		BaseRoutes: &Routes{},
	}
	root.NotFoundHandler = http.HandlerFunc(NotFound)
	api.BaseRoutes.APIRoot = root.PathPrefix("/avoxi").Subrouter()
	listener, err := net.Listen("tcp", ":8013")
	if err != nil {
		panic(err)
	}

	// server listen and serve
	go a.Server.Server.Serve(listener)

	log.Info("[INFO] Listen server on:", listener.Addr().String())
	api.BaseRoutes.Root = root
	a.Server.ListenAddr = listener.Addr().(*net.TCPAddr)
	api.InitRoutes(api.BaseRoutes.Root)

	return api
}
