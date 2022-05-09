package app

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	geolite2 "github.com/oschwald/geoip2-golang"
)

type App struct {
	goroutineCount      int32
	goroutineExitSignal chan struct{}
	geoReader           *geolite2.Reader

	Server *Server
	Ips    map[string][]string
}

const (
	//READTIMEOUT - READTIMEOUT
	READTIMEOUT = 300
	//WRITETIMEOUT - WRITETIMEOUT
	WRITETIMEOUT = 300
)

//RecoveryLogger - struct
type RecoveryLogger struct {
}
type CorsWrapper struct {
	router *mux.Router
}

//Println - RecoveryLogger
func (rl *RecoveryLogger) Println(i ...interface{}) {
	log.Error("[CRASH APPLICATION ERROR]\n", i)
}

func New() *App {
	r := mux.NewRouter()
	var handler http.Handler = &CorsWrapper{r}
	app := &App{
		goroutineExitSignal: make(chan struct{}, 1),
		Server: &Server{
			Router: r,
			Server: &http.Server{
				Handler:      handlers.RecoveryHandler(handlers.RecoveryLogger(&RecoveryLogger{}), handlers.PrintRecoveryStack(true))(handler),
				ReadTimeout:  time.Duration(READTIMEOUT) * time.Second,
				WriteTimeout: time.Duration(WRITETIMEOUT) * time.Second,
			},
		},
		Ips:       make(map[string][]string),
		geoReader: nil,
	}
	return app
}

var allowedMethods = []string{
	"POST",
	"GET",
	"OPTIONS",
	"PUT",
	"PATCH",
	"DELETE",
}

func (cw *CorsWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if checkOrigin(r, "*") {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.Header().Set(
				"Access-Control-Allow-Methods",
				strings.Join(allowedMethods, ", "))

			w.Header().Set(
				"Access-Control-Allow-Headers",
				r.Header.Get("Access-Control-Request-Headers"))
		}
	}

	if r.Method == "OPTIONS" {
		return
	}

	cw.router.ServeHTTP(w, r)
}

func checkOrigin(r *http.Request, allowedOrigins string) bool {
	origin := r.Header.Get("Origin")
	if allowedOrigins == "*" {
		return true
	}
	for _, allowed := range strings.Split(allowedOrigins, " ") {
		if allowed == origin {
			return true
		}
	}
	return false
}
