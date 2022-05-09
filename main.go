package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cesarvspr/avoxi-ip-service/api"
	"github.com/cesarvspr/avoxi-ip-service/app"
	"github.com/labstack/gommon/log"
)

func main() {

	runServer()

}

func runServer() error {
	a := app.New()

	api.Init(a, a.Server.Router)
	a.LoadIso()
	a.StartCronJbos()
	c := make(chan os.Signal, 1)
	log.Info("[AVOXI IP SERVICE STARTED]")

	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c
	return nil
}
