package test_config

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/middleware"
	"os"
)

const (
	LocalTestServer = "127.0.0.1:14005"
)

func Init() {
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDriverName, "mysql"))
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDbHost, "127.0.0.1"))
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDbPort, "14306"))
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDbName, "dkdk"))
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDbUser, "dkdk"))
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDbPasswordFile, "/tmp/secret/MYSQL_PASSWORD_FILE"))
	awesome_error.CheckFatal(os.Setenv(consts.EnvDirSecret, "/tmp/secret"))
	middleware.InitJwtKey()
}