package crawler

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"
)

func DownloadAndWrite() (string, error) {
	// TODO: get api-key from env
	url := "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&license_key=4H4qKSqTQ3sOLRy2&suffix=tar.gz"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	f, err := ioutil.TempFile(os.TempDir(), "iso.tar.gz")

	if err != nil {
		log.Error(err)
	}
	n, err := f.Write(body)
	if err != nil {
		log.Error(err)
	}
	log.Info("Downloaded ", n, " bytes")
	return f.Name(), nil
}
