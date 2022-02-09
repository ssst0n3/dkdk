package database

import (
	"dkdk/test/test_config"
	"github.com/ssst0n3/awesome_libs/cipher"
)

func init() {
	test_config.Init()
	cipher.Init()
	Init()
}
