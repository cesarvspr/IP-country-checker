package app

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router     *mux.Router
	Server     *http.Server
	ListenAddr *net.TCPAddr
}
