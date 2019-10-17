package utils

import (
	"bufio"
	"io/ioutil"
	"os"
)

type Editor struct {
	Path string

	FileIOReadFile  *os.File
	FileIOWriteFile *os.File

	ExistingFile *bufio.Scanner
	TmpFile      *bufio.Writer
}

func NewEditor(path string) (*Editor, error) {
	scannerHandle, err1 := os.Open(path)
	if err1 != nil {
		return nil, err1
	}

	writerHandle, err2 := ioutil.TempFile(".", "bumpit.*.tmp")
	if err2 != nil {
		return nil, err2
	}

	existingFile := bufio.NewScanner(scannerHandle)
	tmpFile := bufio.NewWriter(writerHandle)

	editor := Editor{
		Path:            path,
		FileIOReadFile:  scannerHandle,
		FileIOWriteFile: writerHandle,

		ExistingFile: existingFile,
		TmpFile:      tmpFile,
	}

	return &editor, nil
}

func (e *Editor) ReadLine(readLine func(int, string)) {
	lineNum := 0
	for e.ExistingFile.Scan() {
		readLine(lineNum, e.ExistingFile.Text())
		lineNum++
	}
}

func (e *Editor) EditLine(editLine func(int, string) string) {
	e.ReadLine(func(lineNo int, line string) {
		newLine := editLine(lineNo, line)
		e.TmpFile.WriteString(newLine)
	})

	e.TmpFile.Flush()
}

func (e *Editor) CommitFile() error {
	TFile := e.FileIOReadFile.Name() + ".tmp"

	if err := os.Rename(e.FileIOReadFile.Name(), TFile); err != nil {
		return err
	}
	if err := os.Rename(e.FileIOWriteFile.Name(), e.Path); err != nil {
		return err
	}
	if err := os.Remove(TFile); err != nil {
		return err
	}
	if err := e.FileIOReadFile.Close(); err != nil {
		return err
	}
	if err := e.FileIOWriteFile.Close(); err != nil {
		return err
	}

	return nil
}
