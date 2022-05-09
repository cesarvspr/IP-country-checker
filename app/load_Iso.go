package app

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"strings"

	"github.com/cesarvspr/avoxi-ip-service/crawler"
	"github.com/labstack/gommon/log"
	geolite2 "github.com/oschwald/geoip2-golang"
)

func (a *App) LoadIso() {
	folderCreated := ExtractTarGz()
	log.Info("EXTRACT IS DONE: %v\n", folderCreated)

	geoLiteDB, err := geolite2.Open(folderCreated)
	if err != nil {
		log.Fatal(err)
	}
	a.geoReader = nil
	a.geoReader = geoLiteDB
	log.Info("LOAD DB DONE\n")
}

func ExtractTarGz() string {
	os.Remove(os.TempDir())
	fileName, err := crawler.DownloadAndWrite()

	if err != nil {
		log.Error("download the tar.gz name: %v\n", fileName)
	}

	reader, err := os.Open(fileName)
	if err != nil {
		log.Error(err)
		log.Fatal("open file failed")
	}

	uncompressedStream, err := gzip.NewReader(reader)
	if err != nil {
		log.Info(err)
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)
	var name string
	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err != nil {
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			if err := os.Mkdir(header.Name, 0755); err != nil {
				// handle already exists error
				if strings.HasSuffix(err.Error(), "file exists") {
					os.RemoveAll(header.Name)
					err = os.Mkdir(header.Name, 0755)
					if err != nil {
						log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
					}
				}
				log.Info("ExtractTarGz: Mkdir() failed but application has recovered")
			}
		case tar.TypeReg:
			if !strings.HasSuffix(header.Name, ".mmdb") {
				continue
			}
			name = header.Name
			os.Remove(header.Name)
			outFile, err := os.Create(header.Name)
			if err != nil {
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			outFile.Close()

		default:
			log.Fatalf(
				"ExtractTarGz: uknown type: %s in %s",
				header.Typeflag,
				header.Name)

		}
	}
	return name
}
