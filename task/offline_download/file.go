package offline_download

import (
	"crypto/sha256"
	"fmt"
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io"
	"os"
)

func Info(path string) (dgs digest.Digest, size int64, err error) {
	f, err := os.Open(path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	defer f.Close()

	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	dgs = digest.Digest(fmt.Sprintf("sha256:%x", h.Sum(nil)))
	fi, err := os.Stat(path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	size = fi.Size()
	return
}
