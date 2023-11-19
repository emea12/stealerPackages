package functions

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
) 

func ZipFiles(zipFileName string, fileNames ...string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		zipEntry, err := zipWriter.Create(filepath.Base(fileName))
		if err != nil {
			return err
		}

		_, err = io.Copy(zipEntry, file)
		if err != nil {
			return err
		}
	}

	return nil
}