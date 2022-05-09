package app

import (
	"net"
	"net/http"
)

func (a *App) LoginAPI(ip string, whitelistString []string, r *http.Request) bool {

	var clientIp net.IP

	if len(ip) < 1 {
		clientIpFromRequest := r.RemoteAddr // get client IP
		clientIp = net.ParseIP(clientIpFromRequest)
	} else {
		clientIp = net.ParseIP(ip)
	}

	if !a.IsAuthorized(clientIp, whitelistString) {
		return false
	} else {
		return true
	}
}
