// Code generated by vfsgen; DO NOT EDIT.

package godin

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Templates statically implements the virtual filesystem provided to vfsgen.
var Templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 8, 1, 11, 6, 21, 660925736, time.UTC),
		},
		"/transport": &vfsgen۰DirInfo{
			name:    "transport",
			modTime: time.Date(2019, 8, 1, 11, 6, 44, 716897420, time.UTC),
		},
		"/transport/grpc": &vfsgen۰DirInfo{
			name:    "grpc",
			modTime: time.Date(2019, 8, 23, 18, 7, 12, 539179795, time.UTC),
		},
		"/transport/grpc/client.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "client.go.tmpl",
			modTime:          time.Date(2019, 8, 23, 18, 7, 12, 535179892, time.UTC),
			uncompressedSize: 158,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x49\x2d\x2e\xe1\xe2\xca\xcc\x2d\xc8\x2f\x2a\x51\x50\x4a\xce\xcf\x2b\x49\xad\x28\x51\xe2\xe2\xaa\xae\xd6\x55\x50\x29\x28\xca\x2f\xc9\x4f\x2a\x4d\x53\xb0\xb2\x55\xd0\x0b\x80\x71\x6a\x6b\x21\xb2\xc9\x39\x99\xa9\x79\x25\x60\x39\xdf\xfc\x94\xd2\x9c\x54\xb8\x4c\x49\x41\x0e\x58\x38\x24\x35\xb7\x20\x27\xb1\x04\x2c\xc1\xa5\xaf\xaf\xe0\xe1\xea\xe3\xe3\xaf\xe0\xec\xe3\xe9\xea\x17\x02\xe2\x57\x57\xc3\x4c\xd1\x0b\x4a\x2d\x29\xaa\x74\xce\x2f\xcd\x2b\x01\x29\x06\x04\x00\x00\xff\xff\xe2\x0b\x91\x64\x9e\x00\x00\x00"),
		},
		"/transport/grpc/server.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "server.go.tmpl",
			modTime:          time.Date(2019, 8, 23, 18, 7, 5, 663338565, time.UTC),
			uncompressedSize: 2317,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x54\xcd\x6a\xdb\x40\x10\x3e\x47\x4f\x31\x98\x1e\xec\xe0\xca\xa5\xbd\x19\x1c\x08\x89\x9a\x06\x62\xc7\xb5\x9c\x53\x08\x61\x2b\x8d\x53\x11\x69\xa5\xcc\x8e\x4c\x82\xd0\x1b\xf5\x29\xfa\x64\x45\xbb\x6b\xfd\xf8\xaf\x50\x7a\x8b\x2f\x46\x3b\xf3\xed\x7c\xf3\xcd\x37\x9b\x89\xe0\x59\x3c\x21\x30\x2a\x76\x9c\x28\xc9\x52\x62\xe8\x05\xa9\x64\x7c\xe5\x9e\xe3\x14\xc5\x47\xf8\x90\x51\xca\xe9\x8f\x7c\x05\xe3\x09\xb8\xf3\xcd\x47\x59\x9a\xa8\x42\x5a\x23\xe9\xd8\x34\x0d\xf3\x18\xeb\x08\x67\xb1\x3e\x5e\x62\x92\xc5\x82\x75\xc0\x19\x8d\xa0\x28\x2c\xc8\xbd\xc4\x95\xc8\x63\x3e\x0f\x43\x42\xa5\xca\x72\x5c\x14\xb0\x15\x9b\x57\x8c\xca\xb2\xc2\xfd\xfe\x55\x85\x39\x8b\x5d\xff\x39\xca\xf4\x6d\x45\x01\x24\xe4\x13\x36\x24\x5d\x1f\x69\x1d\x05\xa8\xaa\x38\xbf\x65\x08\x45\x01\xee\x4c\x24\x58\x96\x10\x49\x46\x5a\x89\x00\xa1\x70\x00\x00\x2a\x9a\x06\xef\x2e\xe6\x17\x1a\x02\xf6\xa7\x3b\x20\x7c\xa9\x3a\xe8\x37\xb7\x5f\x21\x4f\x51\x29\xa1\x21\xf8\x92\xa3\xe2\xe5\x5b\x86\x83\x3d\x50\x75\x14\xaa\xb2\x54\x2a\xac\xb1\x1d\xb0\xa5\x74\x91\x26\x09\x4a\x56\xad\xab\xb5\x7a\xe0\x6e\x57\x43\x19\x6e\x1f\x45\x2b\xc0\x17\xe8\xc7\x28\x35\x17\xf7\x6b\x84\x71\x78\x13\x29\x1e\xc0\xa7\x6e\xae\x51\x07\xca\xb2\x1f\xf0\x2b\xd8\xd9\xbb\x17\xe6\x7f\x08\x64\xba\xac\x12\x2b\x3d\x36\xc9\x03\xe8\x23\x11\x20\x51\x4a\x83\x2e\x99\x58\xe1\x7f\xaa\x40\x56\x26\x1b\x52\x9b\xd0\x10\x0e\xd5\x6e\x84\x68\x7d\x1a\x3f\xda\x0f\xa7\x51\xb8\x99\x8d\x1d\x8c\x32\xde\x1d\x9d\xea\x1b\xfc\x6f\xe7\x0b\xef\x12\xa6\xb7\x97\xde\x8d\xef\x9c\x8e\xac\xaf\x8d\xb2\xee\xb5\x9a\xa6\x21\xc6\xc0\x94\xa3\x35\x68\xab\x53\x50\x3f\x05\x61\x08\x49\x95\xd3\xf5\xa1\x8e\x32\xe5\x01\xd7\x2e\xdc\x4c\x5c\x0f\x69\xc7\x86\xb6\xe0\x02\x33\x14\x8c\xa1\xae\xd8\x4a\xb1\x57\x70\xc4\x31\x36\x15\xee\x1f\xaa\x7a\x95\xbf\x76\xcc\xb2\x35\x9f\xfd\xf8\x83\xe8\xbf\x2b\xdc\x15\x7b\x9f\x98\xd0\xf7\x66\x77\x53\x7f\xd0\x88\xba\x3d\x10\x4f\xe6\x89\xaa\x5f\x12\x59\x91\x1a\x4f\xda\x1c\x6b\xc5\xcd\x99\x49\x29\x4b\x40\x99\x27\x48\x82\xa3\x54\xd6\xb2\xd7\xc1\x48\xf2\x97\xcf\x4e\x90\x4a\xc5\xd0\x77\x4e\x5a\xbb\xd6\x28\x7f\xd2\x02\x3c\xee\xea\x32\xd1\xca\xdc\x52\x88\xb4\x47\x87\x81\xe9\x18\xae\x67\x4b\x98\x9c\x81\xbf\x5c\x5c\xcf\xae\x40\x77\xb9\x16\xd4\xe6\xf2\xa8\xff\x27\x90\x88\xec\x5e\xd3\x7a\x50\x4c\x91\x7c\x2a\x8e\xd0\xaa\xcb\x8e\xa1\xd7\x72\x53\x6f\x68\x30\xf5\x2c\x0c\x07\x5b\x7c\x72\xa6\xd9\xec\xe5\xb0\x16\x71\xbe\x21\x61\xca\x3f\x68\x2e\x3b\xaf\x63\xd7\x98\x9d\xe2\xe3\x8e\x20\xc3\xe3\xbb\xd7\xf8\x61\xe1\x7d\xbf\xf3\xfc\xa5\x7f\xd8\x02\xdd\x9d\x6c\xed\x9d\x7d\x77\xeb\xcd\x7b\x17\xdb\x75\x74\xb9\x16\x9e\x3f\xbf\x9d\xf9\xde\xbf\xa9\x69\xdf\xd8\x77\x2f\xe7\x9f\x00\x00\x00\xff\xff\xc8\x8a\x0e\xe7\x0d\x09\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/transport"].(os.FileInfo),
	}
	fs["/transport"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/transport/grpc"].(os.FileInfo),
	}
	fs["/transport/grpc"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/transport/grpc/client.go.tmpl"].(os.FileInfo),
		fs["/transport/grpc/server.go.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}