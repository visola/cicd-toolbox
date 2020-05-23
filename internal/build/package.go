package build

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

//CreatePackage create a zip package adding all content from a directory and also some aditional files.
func CreatePackage(outputFile string, contentDirectory string, additionalFiles ...string) error {
	zipFile, createErr := os.Create(outputFile)
	if createErr != nil {
		return createErr
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	addDirToZip(zipWriter, contentDirectory)
	for _, fileToAdd := range additionalFiles {
		addFileToZip(zipWriter, fileToAdd)
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, pathToFile string) error {
	fileStat, statErr := os.Stat(pathToFile)
	if statErr != nil {
		return statErr
	}

	return addFileToZipWithStat(zipWriter, pathToFile, fileStat)
}

func addFileToZipWithStat(zipWriter *zip.Writer, pathToFile string, fileStat os.FileInfo) error {
	fileToAdd, openErr := os.Open(pathToFile)
	if openErr != nil {
		return openErr
	}
	defer fileToAdd.Close()

	zipHeader, headerErr := zip.FileInfoHeader(fileStat)
	if headerErr != nil {
		return headerErr
	}

	zipHeader.Method = zip.Deflate

	fileWriter, createErr := zipWriter.CreateHeader(zipHeader)
	if createErr != nil {
		return createErr
	}

	_, writeErr := io.Copy(fileWriter, fileToAdd)
	return writeErr
}

func addDirToZip(zipWriter *zip.Writer, pathToDir string) error {
	return filepath.Walk(pathToDir, func(childPath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if childPath == pathToDir {
			return nil
		}

		if fileInfo.IsDir() {
			// recurse
			return nil
		}

		return addFileToZipWithStat(zipWriter, childPath, fileInfo)
	})
}
