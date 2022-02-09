package database

import (
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/dkdk/test/test_config"
)

func init() {
	test_config.Init()
	cipher.Init()
	Init()
}
