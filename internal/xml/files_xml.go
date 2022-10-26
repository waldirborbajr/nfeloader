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
				if err = MoveXML(config.AppPath, f.Name(), true); err != nil {
					customlog.HandleError("ZERO Size Error Moving", err)
				}
			}
		}
	}

	return files, nil
}

func moveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
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
		err = moveFile(f, errorPath)
	} else {
		err = moveFile(f, processedPath)
	}

	// err := os.Remove(f)
	if err != nil {
		customlog.HandleError("Renaming XML", err)
		return err
	}
	return nil
}
