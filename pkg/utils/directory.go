package utils

import (
	"io/ioutil"
)

func GetFilesInDir(path string) ([]string, error) {
	fileObjs, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	files := []string{}
	for _, f := range fileObjs {
		files = append(files, f.Name())
	}
	return files, nil
}
