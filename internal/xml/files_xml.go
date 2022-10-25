package xml

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/waldirborbajr/nfeloader/internal/config"
	"github.com/waldirborbajr/nfeloader/internal/customlog"
	"github.com/waldirborbajr/nfeloader/internal/entity"
)

// List all XMLs files and return a slice of files names
func ListXML(path string) ([]string, error) {
	var files []string

	xmlFiles, err := os.ReadDir(path)
	if err != nil {
		customlog.HandleError("Reading directory", err)

		return files, err
	}

	for _, f := range xmlFiles {
		ext := strings.ToUpper(filepath.Ext(f.Name()))

		if ext == ".XML" {
			fileInfo, err := f.Info()
			if err != nil {
				customlog.HandleError("file.Info()", err)
			}
			if fileInfo.Size() != 0 {
				files = append(files, f.Name())
			} else {
				err = MoveXML(config.AppPath, f.Name(), true)
			}

		}
	}

	return files, nil
}

// Read the content of a XML file and return a struct of content
func ReadXML(path string, file string) (*entity.NFeProc, error) {
	f := fmt.Sprintf(path + file)

	xmlFile, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("Processing %s", file)

	defer xmlFile.Close()

	nfeProc := &entity.NFeProc{}

	xmlContent, _ := io.ReadAll(xmlFile)

	// if err := xml.Unmarshal([]byte(xmlContent), &nfe); err != nil {
	if err := xml.Unmarshal(xmlContent, nfeProc); err != nil {
		customlog.HandleError("XML unmarshal", err)
		return nil, err
	}

	return nfeProc, nil
}

// Move XML file to processed folder
func MoveXML(path string, file string, hasError bool) error {
	var err error

	f := fmt.Sprintf(path + "/xmls/" + file)
	processedPath := fmt.Sprintf(path + "/xmlprocessed/" + file)
	errorPath := fmt.Sprintf(path + "/xmlerror/" + file)

	if hasError {
		err = os.Rename(f, errorPath)
	} else {
		err = os.Rename(f, processedPath)
	}

	// err := os.Remove(f)
	if err != nil {
		customlog.HandleError("Renaming XML", err)
		return err
	}
	return nil
}
