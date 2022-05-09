package app

import (
	"time"

	"github.com/labstack/gommon/log"
	cron "github.com/robfig/cron/v3"
)

var NewStockFlag = 0

// RunServices - agendando serviços
func (a *App) StartCronJbos() {

	c := cron.New(cron.WithLocation(time.UTC))

	// every day at 00:30:00
	c.AddFunc("30 3 * * *", func() {
		currentTime := time.Now().UTC()
		log.Info("[CRON] daily loading of ISO from provider at", currentTime)
		a.LoadIso()
	})

	//this is for testing
	// c.AddFunc("@every 10s", func() {
	// 	//this is for testing
	// 	currentTime := time.Now().UTC()
	// 	log.Info("[CRON] daily loading of ISO from provider at", currentTime)
	// 	a.LoadIso()
	// })

	c.Start()

}
