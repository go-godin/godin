package godin

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type baseWriter struct {
	Path string
	data []byte
}

type Writer struct {
	baseWriter
}

// NewFileWriter initializes a new Writer
func NewFileWriter(path string, data []byte) *Writer {
	return &Writer{
		baseWriter{
			Path: path,
			data: data,
		},
	}
}

// Write dumps the given data into a file and creates it if necessary.
// The overwrite flag can be set to overwrite any existing data.
func (f *Writer) Write(overwrite bool) error {
	if _, err := os.Stat(f.Path); err == nil {
		if !overwrite {
			return fmt.Errorf("%s file already exists, overwrite is disabled", f.Path)
		}
	} else {
		if _, err := os.Create(f.Path); err != nil {
			return fmt.Errorf("failed to create file %s: %s", f.Path, err)
		}
	}

	if err := ioutil.WriteFile(f.Path, f.data, 0644); err != nil {
		return errors.Wrap(err, "failed to write file")
	}
	return nil
}

type AppendWriter struct {
	baseWriter
}

// NewFileAppendWriter returns a new appending file-writer for Godin templates
func NewFileAppendWriter(path string, data []byte) *AppendWriter {
	return &AppendWriter{
		baseWriter{
			data: data,
			Path: path,
		},
	}
}

// Write will open the given file and try to append the given data to it
// The file is NOT created if it doesn't exist.
func (f *AppendWriter) Write() error {
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return errors.Wrap(err, "file to append cannot be opened")
	}

	if _, err := file.Write(f.data); err != nil {
		return err
	}
	return nil
}
