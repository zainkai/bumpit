package utils

import (
	"io/ioutil"
	"os"
)

type FileReader struct {
	path    string
	name    string
	content *[]byte
}

func New(path, name string) *FileReader {
	return &FileReader{
		path, name, nil}
}

func (this *FileReader) getFullPath() string {
	return this.path + "/" + this.name
}

func (this *FileReader) Exists() bool {
	_, err := os.Stat(this.getFullPath())
	return os.IsNotExist(err)
}

func (this *FileReader) Load() error {
	data, err := ioutil.ReadFile(this.getFullPath())
	if err != nil {
		return err
	}

	this.content = &data
	return nil
}
