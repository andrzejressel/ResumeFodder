package templates

import (
	"embed"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//go:embed *.xml
var f embed.FS

func GetTemplates() (map[string]string, error) {
	templates := make(map[string]string)

	dirEntryList, err := f.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, dirEntry := range dirEntryList {
		fileName := dirEntry.Name()
		file, err := f.Open(fileName)
		if err != nil {
			return nil, err
		}
		content, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		templates[removeExtension(fileName)] = string(content)
	}

	return templates, nil
}

func removeExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
