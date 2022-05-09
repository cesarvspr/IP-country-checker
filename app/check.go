package app

import (
	"log"
	"strings"
)

func (a *App) IsAuthorized(ip []byte, whitelist []string) bool {
	geoLiteDB := a.geoReader

	record, err := geoLiteDB.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	isoCode := record.Country.IsoCode

	for _, item := range whitelist {
		item = strings.ToUpper(item)
		if item == isoCode {
			return true
		}
	}
	return false
}
