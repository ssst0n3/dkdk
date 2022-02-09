package config

import (
	"os"
	"strings"
)

const (
	EnvCert            = "CERT_NAME"
	EnvLocalListenPort = "LOCAL_LISTEN_PORT"
	EnvAllowOrigins    = "ALLOW_ORIGINS"
)

var (
	LocalListenPort = os.Getenv(EnvLocalListenPort)
	CertificateName = os.Getenv(EnvCert)
	AllowOrigins    []string
)

func init() {
	origins := os.Getenv(EnvAllowOrigins)
	if len(strings.TrimSpace(origins)) == 0 {
		AllowOrigins = nil
	} else {
		AllowOrigins = strings.Split(origins, ",")
	}
}
