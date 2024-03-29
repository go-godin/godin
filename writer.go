package godin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

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

type TemplateWriter struct {
	fs http.FileSystem
}

func NewTemplateWriter(fs http.FileSystem) *TemplateWriter {
	return &TemplateWriter{fs: fs}
}

func (tw *TemplateWriter) OverWrite(sourcePath, targetPath string) error {
	if err := tw.write(sourcePath, targetPath); err != nil {
		return errors.Wrap(err, "Overwrite")
	}
	return nil
}

func (tw *TemplateWriter) Write(sourcePath, targetPath string) error {
	if _, err := os.Stat(targetPath); err == nil {
		return fmt.Errorf("target template already exists, local version will not be overwritten")
	}
	if err := tw.write(sourcePath, targetPath); err != nil {
		return errors.Wrap(err, "Write")
	}
	return nil
}

func (tw *TemplateWriter) EnsurePath(path string) error {
	if _, err := os.Stat(filepath.Join(path)); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
	}
	return nil
}

func (tw *TemplateWriter) write(sourcePath, targetPath string) error {
	f, err := tw.fs.Open(sourcePath)
	if err != nil {
		return errors.Wrap(err, "unable to open template source")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "unable to read from template file")
	}

	if err := ioutil.WriteFile(targetPath, buf, 0644); err != nil {
		return errors.Wrap(err, "unable to write template into target")
	}
	return nil
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
