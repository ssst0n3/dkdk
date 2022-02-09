package offline_download

import (
	"fmt"
	"github.com/alexmullins/zip"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"io"
	"os"
)

func Archive(path, filename, password string) (outputPath string, err error) {
	log.Logger.Info()
	input, err := os.Open(path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	outputPath = fmt.Sprintf("%s.zip", path)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	zipWriter := zip.NewWriter(outputFile)
	defer zipWriter.Close()
	w, err := zipWriter.Encrypt(filename, password)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	_, err = io.Copy(w, input)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = zipWriter.Flush()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
