package main

import (
	"dkdk/config"
	"dkdk/database"
	"dkdk/router"
	"dkdk/task/offline_download"
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/docker_secret/cert"
)

func main() {
	database.Init()

	go func() {
		offline_download.Run()
	}()

	r := router.InitRouter()
	port := fmt.Sprintf(":%s", config.LocalListenPort)
	if len(config.CertificateName) > 0 {
		log.Logger.Info("self-signed certificate configured")
		ca, key := cert.CertificateFilePath(config.CertificateName)
		awesome_error.CheckFatal(r.RunTLS(port, ca, key))
	} else {
		awesome_error.CheckFatal(r.Run(port))
	}
}
