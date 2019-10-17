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

func IsInDir(path, fName string) (bool, error) {
	directory, err := GetFilesInDir(path)
	if err != nil {
		return false, err
	}

	for _, s := range directory {
		if s == fName {
			return true, nil
		}
	}

	return false, nil

}
