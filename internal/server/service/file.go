package service

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type File struct {
	FilePath string
	buffer   *bytes.Buffer
}

func NewFile() *File {
	return &File{
		buffer: &bytes.Buffer{},
	}
}

func (f *File) SetFilePath(fileName, path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	f.FilePath = filepath.Join(path, fileName)
	return nil
}

func (f *File) Write(chunk []byte) error {
	_, err := f.buffer.Write(chunk)

	return err
}

func (f *File) WriteFile() error {
	if err := ioutil.WriteFile(f.FilePath, f.buffer.Bytes(), 0o644); err != nil {
		return err
	}

	return nil
}
