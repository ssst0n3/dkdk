package offline_download

import (
	"crypto/tls"
	"github.com/Code-Hex/pget"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"io"
	"net/http"
	"os"
)

func Download(url, path string) (err error) {
	log.Logger.Info()
	//err = DownloadByMultiProcess(url)
	//if err != nil {
	err = DownloadBySingleThread(url, path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
	//}
	//return
}

func DownloadByMultiProcess(url string) (err error) {
	log.Logger.Info()
	p := pget.New()
	p.URLs = append(p.URLs, url)
	err = p.Checking()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = p.Download()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func DownloadBySingleThread(url, path string) (err error) {
	log.Logger.Info()
	out, err := os.Create(path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	defer out.Close()

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}
	// Create Http Client
	client := &http.Client{Transport: transCfg}
	resp, err := client.Get(url)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
