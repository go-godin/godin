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
			modTime: time.Date(2019, 8, 27, 18, 58, 25, 470560916, time.UTC),
		},
		"/endpoint": &vfsgen۰DirInfo{
			name:    "endpoint",
			modTime: time.Date(2019, 8, 27, 18, 58, 34, 894542425, time.UTC),
		},
		"/endpoint/endpoints.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "endpoints.go.tmpl",
			modTime:          time.Date(2019, 8, 27, 18, 58, 34, 886542442, time.UTC),
			uncompressedSize: 1987,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\xcd\x6e\xa3\x30\x10\xbe\xf3\x14\xb3\x68\x0f\x20\x25\xe6\x1e\x29\xa7\x34\xbb\xaa\xb4\x6d\xa3\x26\x2f\x40\xcd\x90\x7a\x4b\x6c\x6a\x9b\xb4\x15\xe2\xdd\x57\x63\x0c\x38\x7f\xbd\xee\x81\x24\xf6\xcc\x7c\xf3\xcd\xcc\x37\x24\xcb\x60\xa5\x0a\x84\x3d\x4a\xd4\xb9\xc5\x02\x5e\xbe\x60\xaf\x0a\x21\x19\xdc\x3d\xc1\xe3\xd3\x0e\xd6\x77\xf7\x3b\x16\xd5\x39\x7f\xcb\xf7\x08\x28\x8b\x5a\x09\x69\xa3\xa8\x6d\xe7\xf0\xb3\xd6\xea\x2f\x72\x0b\x8b\x25\xb0\x8d\xff\xdd\x75\xa3\xcd\xaa\x97\xa6\x1c\x8c\xfd\x61\xb0\x1a\xd4\x47\xd4\xce\xf6\xa0\x8a\xa6\xc2\xd1\x62\xeb\xca\x5d\xef\xf0\x50\x57\xb9\x75\x86\x48\x1c\x6a\xa5\x2d\x24\x11\x00\x40\xcc\x95\xb4\xf8\x69\xe3\xfe\xb4\x17\xf6\xb5\x79\x61\x5c\x1d\xb2\xbd\x9a\xbf\x09\x9b\xd1\x33\x30\x8d\x23\xe7\xc5\x20\x6e\xdb\x91\xf0\x94\x33\x13\xd2\xa2\x96\x79\x95\x85\xe6\x2d\xea\xa3\xe0\xc8\x1e\xf3\x03\x39\xc5\x51\x4a\x05\x83\xce\xe5\x1e\xa7\xca\x06\x37\x43\x14\xed\x57\x8d\xb0\x45\x0b\xc6\xea\x86\x5b\x68\x5d\xda\x31\x88\x3d\x6f\x56\xce\xaf\xbf\x75\xc8\x5d\xb7\xf6\x24\xc7\xbe\xb2\xe1\xc6\xfb\xcd\xc9\x42\x61\x5d\x14\x95\x8d\xe4\xf0\x88\x1f\x83\xcb\x16\x6d\x62\x8e\x1c\x3c\x8b\xd4\x65\xff\x2e\xed\x31\xd7\x64\xa9\xd4\x07\x6a\x18\x6a\xbb\x95\xda\x7d\x7a\xac\xb3\x88\x25\x3c\xe4\x6f\x78\x51\x04\x91\x49\x5d\x58\x77\xce\xde\x9d\x35\xda\x46\x4b\x62\x19\x82\xcf\xaf\x31\xbd\xda\xa4\xc5\x19\x95\xae\x9b\x9d\xe0\xf8\x5c\x3d\x81\x2e\xba\xec\x81\x93\x97\xc6\x77\x92\x57\x32\x4d\xf1\x37\xda\x07\x34\x26\x77\xae\xf8\xde\xa0\xb1\xbb\xaf\x1a\xd3\x20\xc4\x7c\x1b\x62\x6a\x25\x0d\x8e\x31\x51\x96\x5d\x6f\x10\x70\x25\x7b\x75\x18\xb0\xaf\x08\xe3\x7d\xa9\xb4\xbb\x18\x23\xe0\x79\xb3\xea\xe7\x7d\xb3\xd3\xd3\xd8\x2f\x06\xe8\x87\xe7\xfb\x4d\x30\x09\xb7\x9f\xe0\xb7\x86\xad\xfa\xef\x19\xe8\xbe\x58\x70\x2b\x50\xe6\x1c\xdb\x2e\x85\x44\xfb\x7a\xc2\xeb\x19\xa0\xd6\xf4\x28\x9d\x06\xd2\xf0\xbd\xf4\x38\x2c\x69\x5b\x6a\xaf\x27\x9b\x06\xc3\xa1\x16\xb2\x5f\x02\xab\x82\x6c\xa6\xf3\x80\x8b\x25\x98\x23\x67\x63\x7d\x44\x73\x06\x1e\x65\x72\xdf\x68\x2c\xc5\x27\x16\x10\xd3\x7d\x1c\x42\x13\x59\x82\xf1\x29\x7a\x98\x89\xe0\x79\xf6\xad\xeb\xfe\xbd\x14\x36\x14\xcf\x5a\xeb\x05\xf1\x99\x6e\xba\x20\x81\x6b\x22\xe5\x99\x81\x14\x55\x20\xaf\x49\xdd\xff\x55\x6a\x93\x68\xc4\xa1\xae\xf0\x80\xd2\xcb\xcb\x0b\x64\x9a\x23\x08\x09\x4a\x17\xa8\x9d\xe0\xb8\x92\x47\x94\x02\xa5\x85\x9c\x73\x34\x06\xac\x72\x71\x83\x9e\x0c\x23\xf8\xdd\xab\x30\x20\x0c\x34\x06\xcb\xa6\x22\x08\x5e\x51\x90\x61\xbd\x40\x13\x43\x1b\x9d\xc2\xc9\x10\x2f\xb5\x46\x6f\xd7\x71\xaa\x7f\x84\xb1\xc1\x1f\x44\x47\xb2\xeb\x1d\xcc\x0d\x87\x2b\x02\xec\x47\x32\xc8\x88\x5d\x2e\x49\x28\x26\x2f\x8c\x53\x1a\x93\x1a\xe8\xad\xe6\x55\x25\x4a\x87\xf9\x63\x49\xd3\x3e\xd1\xba\x13\x82\x14\x95\x4b\x1a\xbc\xe7\xc6\x85\x71\xab\x60\x6a\xbf\x07\xe6\x64\x0f\x7c\xf8\x49\x95\x17\xe2\xee\x71\x58\xec\x0a\x1e\x8f\x6b\xad\xcf\xf4\x16\x1e\xfe\x05\x00\x00\xff\xff\xca\xf0\xb0\xde\xc3\x07\x00\x00"),
		},
		"/endpoint/request_response.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "request_response.go.tmpl",
			modTime:          time.Date(2019, 8, 27, 18, 54, 54, 302961924, time.UTC),
			uncompressedSize: 919,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x4f\x8f\xd3\x30\x10\xc5\xcf\xf5\xa7\x18\xad\x38\xb4\x12\x38\x77\x24\x4e\xdb\x2e\xe2\xb0\x7f\xc4\xf6\x0b\x78\xe3\x97\x60\x48\x6c\x77\xec\x2c\xaa\xa2\x7e\x77\x64\xc7\x4d\x0b\x11\xac\xb8\x39\xf3\xfc\x66\x7e\xcf\x93\xaa\xa2\x5b\xa7\x41\x2d\x2c\x58\x45\x68\x7a\x39\x52\xeb\xb4\xb1\x92\xb6\x8f\xf4\xf0\xb8\xa7\xdd\xf6\xcb\x5e\x0a\xaf\xea\x1f\xaa\x05\xc1\x6a\xef\x8c\x8d\x42\x8c\xe3\x07\x7a\xe7\xd9\x7d\x47\x1d\xe9\xe3\x27\x92\x4f\xe5\x7c\x3a\xcd\x5a\x74\x2f\x43\x73\x16\xa7\x8f\xb3\x1a\xc0\xaf\xe0\xac\xdd\x3b\x3d\x74\x98\x95\xe8\xbb\x5c\xde\xa3\xf7\x9d\x8a\x59\x10\xa6\xf7\x8e\x23\xad\xc5\x4a\xd2\xcd\x38\xce\x93\x2f\xe6\xca\xd8\x08\xb6\xaa\xab\xae\xe5\x67\xf0\xab\xa9\x21\x1f\x54\x9f\x2e\xdd\x88\x8d\x48\xe8\xc4\xca\xb6\xb8\x30\x9e\xef\x85\x89\xa2\xc8\xf2\xeb\xd3\x6d\x98\xb9\x18\x87\xc4\xb5\xbe\x98\x3e\x23\xde\x23\x04\x95\xaf\xe2\x30\x20\xc4\xfd\xd1\x63\x73\x65\x09\xff\xb4\x04\xef\x6c\xc0\xec\x11\x55\x45\xe3\x98\x59\x4f\xa7\xd2\x90\x7e\xb2\xf2\x81\xe2\x37\x10\x97\x4a\x63\xd0\xe9\x40\xae\xc9\xd5\xd9\x40\xbb\xb2\x1b\x29\xe2\xd1\x63\xd9\x29\x44\x1e\xea\x48\xa3\x58\xa5\x17\x62\x1c\xe4\x5d\xea\xf4\x9c\xcb\x5b\xd4\x9d\xe2\xab\x27\x49\x40\x0b\xa4\x09\xf8\x37\xa6\x52\x7a\x0b\xea\x3d\x29\xad\x8d\x6d\x49\x59\x02\xb3\xe3\xc9\x91\xfa\x9b\xde\x77\xe8\x61\x63\x92\x93\xfb\x4e\x99\x0e\x4c\x79\xa1\x8d\xaa\xb1\x0c\x54\x86\xfe\x99\x28\xbc\x99\x68\xb5\x63\x9e\xe6\x8b\x69\x47\xb0\xfa\xbc\xae\x72\xfc\xff\xdf\xa3\x19\x6c\x4d\x6b\x5e\x02\x6e\xa6\x28\x7a\xbd\x29\x99\x47\x62\xc4\x81\x2d\xb1\x4c\x20\x7f\x43\xf8\x15\x00\x00\xff\xff\x37\x88\x07\x82\x97\x03\x00\x00"),
		},
		"/protobuf_example.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "protobuf_example.go.tmpl",
			modTime:          time.Date(2019, 8, 27, 18, 58, 25, 466560925, time.UTC),
			uncompressedSize: 2277,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x54\x4f\x6f\x9b\x4e\x10\x3d\x87\x4f\x31\xb2\x7e\x07\x88\xfc\xc3\x55\x7b\xb3\x44\xa4\x2a\xa1\x69\xa4\xd8\x49\xc1\x39\x45\x51\xb4\x85\x89\x4b\x0b\x0b\xde\x1d\xac\x44\x88\xef\x5e\xb1\xbb\xfc\xf3\xbf\x4a\x55\x6f\xf1\x05\xc3\xec\xdb\xf7\xf6\xcd\x9b\x2d\x58\xf4\x8b\xad\x11\x08\x25\x59\x56\x92\x15\xb9\x20\x98\x44\x39\x27\x7c\xa5\x89\x65\x55\xd5\xff\xf0\x5f\x21\xf2\x9f\x18\x11\xcc\x3d\x70\xef\xcd\xff\xba\xee\x6a\x94\x7f\x2f\x5f\xda\xa2\x7e\x69\xab\x12\xc5\x16\x85\xaa\x2d\xf2\xb8\x4c\xb1\xab\x50\x91\xaa\xcf\x2b\xcc\x8a\x94\x91\x2a\x58\x55\x05\x82\xf1\x35\xf6\xdb\xba\x21\x8a\x6d\x12\xa1\x6c\xea\xf4\x56\x20\x54\x15\xb8\x4b\x96\x61\x5d\x43\xc2\x09\xc5\x0b\x8b\x10\x2a\x0b\x00\xa0\xd9\x58\xe3\xdd\xe0\xfe\x52\x41\xc0\xfc\x14\xa7\xc0\x4d\xc3\x69\xf7\xbb\x5f\x23\x2d\x50\x4a\xa6\x20\xb8\x29\x51\xd2\xea\xad\x40\xe7\x00\x54\x9e\x84\xca\x22\xe7\x12\x3b\xec\x08\x6c\x24\x5d\xe6\x59\x86\x9c\xe4\x60\xeb\xd9\x4c\x1d\x67\x97\x0d\x79\xbc\xfb\x29\x79\x01\xdc\x80\x9d\x22\x57\x5a\xdc\x2f\x09\xa6\xf1\x6d\x22\xa9\xb7\xca\x81\x0f\x63\x94\xf6\x09\xea\xda\x8e\xe8\x15\x4c\x4f\xdd\x4b\xfd\x9c\x82\xd0\xe7\x6d\x16\x36\xce\xb4\x8b\x1d\xb0\x51\x08\x40\x21\x72\xe1\x8c\x65\xa5\x12\xff\x11\x83\x30\x86\x99\x92\x6c\x4b\x53\x38\xc6\xdd\x5b\x32\x78\xd5\x59\x32\x2f\x56\xef\x75\xdf\x25\xd3\x22\xa9\x73\x37\x3b\x57\x3b\x84\x5f\x3f\x07\xfe\x15\x2c\xee\xae\xfc\xdb\xd0\x3a\x9f\x99\x4c\x6a\x8f\xdd\x1b\xb9\xc8\x63\x4c\x81\x44\xa9\x8e\x6b\x9a\x64\x14\x82\xfc\xc1\x04\xc6\x90\x35\x6b\xc6\x89\x54\x55\x12\x65\x44\x5d\x1e\xdb\xde\xab\x76\xed\x05\xd2\x10\x06\x58\x20\x23\x8c\x15\xe3\x60\x89\xd9\x82\x12\x4a\xb1\x67\x78\x7c\x6a\xf8\x9a\xa4\xed\xc5\x66\xa7\x3f\x87\xf1\x47\xd1\x7f\x76\x78\x6c\xf6\x21\x33\xc1\xf6\x97\x0f\x8b\xd0\xe9\x4d\xdd\x6d\x88\xcf\xcb\x4c\x76\xb7\x00\x6f\x44\xcd\xbd\xa1\xc6\xce\x71\xfd\x4d\x2f\xa9\x6b\x40\x5e\x66\x28\x18\x25\x39\xef\x6c\xef\x8a\x09\xa7\x4f\x1f\xad\x28\xe7\x92\xc0\xb6\xce\x06\x53\xd7\x3b\x7f\x36\x00\x3c\xef\xfb\xe2\x29\x67\xee\x44\x8c\xe2\x80\x0f\x8e\x3e\x31\xdc\x2c\x57\xe0\x5d\x40\xb8\x0a\x6e\x96\xd7\xa0\x4e\xb9\x65\x62\xa8\xe5\x59\x3d\x3d\xc8\x58\xf1\xa8\x64\x3d\x49\x12\x09\x5f\x57\x27\x64\x75\xb4\x73\x98\x0c\xd2\x34\x99\x6a\x4c\xd7\x0b\xad\xc1\x90\x7b\x17\x4a\xcd\x41\x0d\x5b\x96\x96\xad\x08\x4d\xff\xa4\xb4\xec\xdd\x93\xe3\x60\x8e\xc8\xe7\x23\x43\xa6\xa7\x67\xaf\xcf\x43\xe0\x7f\x7b\xf0\xc3\x55\x78\x3c\x02\xe3\x99\x1c\xcc\x9d\xb9\x81\xbb\xc9\x7b\x17\xd3\x75\x72\xb8\x02\x3f\xbc\xbf\x5b\x86\xfe\xdf\xb9\x69\xee\xd8\x77\x6f\xe7\xef\x00\x00\x00\xff\xff\x72\xaa\xb1\xe7\xe5\x08\x00\x00"),
		},
		"/service": &vfsgen۰DirInfo{
			name:    "service",
			modTime: time.Date(2019, 8, 27, 18, 59, 5, 170482727, time.UTC),
		},
		"/service/interface.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "interface.go.tmpl",
			modTime:          time.Date(2019, 8, 27, 18, 59, 5, 166482734, time.UTC),
			uncompressedSize: 642,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x51\x6b\xdb\x30\x14\x85\xdf\xfd\x2b\x2e\xa5\x50\x1b\x5a\xf9\x7d\xb0\x87\x91\x6e\xa3\xb0\xb6\x61\xf3\x1f\x50\xa5\x63\x47\x9b\x2d\x29\xd2\x4d\xd6\x20\xfc\xdf\x87\x65\xc7\x19\x81\xed\xc1\xd8\x57\xe7\x1c\xdf\x4f\xf7\xd6\x35\x6d\x9c\x06\x75\xb0\x08\x92\xa1\xe9\xed\x44\x9d\xd3\xc6\x0a\x7a\x7c\xa5\x97\xd7\x86\x3e\x3f\x3e\x35\xa2\x48\xe9\x81\x6e\x7d\x70\x3f\xa1\x98\x3e\x7c\x24\xb1\x5d\xbe\xc7\x71\xd5\xd8\xbd\x1d\xda\xb3\x38\x17\x67\x35\x22\x1c\x11\xb2\xf6\xec\xf4\xa1\xc7\xaa\xb0\xef\xf3\x71\x83\xc1\xf7\x92\xb3\x50\x78\xa9\x7e\xc9\x0e\x94\xd2\xda\x54\xfc\x40\x38\x1a\x05\xf1\x22\x87\xd9\x64\x06\xef\x02\x53\x59\x10\x11\xdd\x28\x67\x19\xef\x7c\x53\x54\x45\x91\x12\x05\x69\x3b\x5c\xa8\xce\xe9\x38\x25\xeb\x9a\x96\x92\x4c\x24\xde\x81\x06\x69\x2c\x19\xcb\x08\xad\x54\xa0\xdf\x3b\xa3\x76\x93\xa6\x11\xcc\x11\x9a\xda\xe0\x86\x6c\x5c\x6f\xf9\x69\xfb\x44\xd1\x43\x99\xd6\x28\xc9\xc6\x59\x2a\xef\x52\xca\x70\xe3\x78\x57\x15\x7c\xf2\xb8\x74\x59\xff\x9c\x32\xec\xca\x27\xbe\x6f\x37\x19\x69\x3e\x7d\xa0\xdb\x80\xfd\x34\x8e\xf2\x42\xfe\x15\xfc\x8c\x18\x65\xb6\x63\x7f\x40\xe4\xe6\xe4\x51\x5d\xc5\xe2\x7f\x63\xd1\x3b\x1b\x71\x95\x5b\x68\x4b\xc5\xef\xb4\x8c\x4f\x6c\xe6\xf7\x7d\x1e\x7d\xc0\x5e\x7c\x31\xe8\xf5\x37\x13\xf9\xaf\x15\x8f\x63\x45\xe5\x6c\x88\xff\x30\xdc\x13\x42\x98\x1e\x17\xaa\x15\x13\x56\x4f\xdd\xe7\xcd\x2f\xc5\x9f\x00\x00\x00\xff\xff\xe8\x23\xee\xad\x82\x02\x00\x00"),
		},
		"/service/models.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "models.go.tmpl",
			modTime:          time.Date(2019, 8, 27, 18, 54, 5, 463050331, time.UTC),
			uncompressedSize: 1039,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xcf\x8a\xdb\x3c\x14\xc5\xd7\xa3\xa7\x38\x0c\x1f\xe4\x6b\x98\xda\xd0\xee\x02\x9e\x55\xdb\x21\x8b\xa4\xa5\x09\x85\x52\x86\xa2\xb1\xae\x6d\x75\x64\xc9\x95\xe4\xa4\xc1\xf8\xdd\x8b\xe4\x3f\x49\x86\xce\xce\xd2\xbd\xba\xe7\x77\xcf\x71\xd7\xbd\xc5\x7f\x8d\x35\xbf\x28\xf7\x58\x65\x48\xbe\x8c\xdf\x7d\xcf\xa6\x9a\x37\x4f\x6d\x31\x15\x87\xc3\x54\x75\x64\x0f\x64\x63\x6d\x63\x44\xab\x68\xae\xf8\x46\xc5\xeb\x3d\xd5\x8d\xe2\x3e\x16\x58\xc3\xf3\x67\x5e\x12\xba\x6e\x16\x4d\x76\x64\x0f\x32\xa7\x64\xcb\xeb\xa1\x29\x5d\xb2\x7d\x45\xa8\x8d\x20\xe5\xc0\x2d\xe1\x89\x3b\x12\x30\x45\x01\x5f\x11\x5c\x43\xb9\x2c\x24\x09\xd4\xe4\x1c\x2f\xc9\x41\x6a\x4c\x9c\x09\xbe\x9b\x16\x39\xd7\x68\x1d\x85\xfe\x9a\x71\x07\x1e\x67\xa0\x30\x36\x8e\x10\xa6\xe6\x52\x2f\x1c\x2a\x29\x04\xe9\x59\xcb\x85\xea\x09\x47\xa9\x14\x6a\xe3\x3c\x94\x7c\x26\x75\x42\x6e\xac\xa5\xb8\xc5\x51\xfa\x2a\x34\x31\x57\x71\x1b\x10\xe2\xcb\x84\x3d\x18\x21\xf5\xf0\x50\x1b\x0f\x4b\x5a\x50\xd0\x92\x0e\x85\x54\x84\x63\x45\x1a\x39\x57\x4a\xea\x12\x8b\x32\x76\x97\xa4\xc9\x72\x4f\x8b\x3b\x18\xad\x4e\x90\x05\xa4\x87\x30\xe4\xf4\xc2\x83\xfe\x48\xe7\x13\x86\x65\xca\x58\xd7\xc1\x72\x5d\xd2\x39\x8e\x64\x33\xed\x1e\x0d\x0f\x6f\xe9\x37\x92\xb5\xdb\x04\x20\x78\xdb\x46\x33\xfd\xa9\x09\x6e\x47\x73\xfb\x1e\xce\xdb\x36\xf7\xe8\xd8\x4d\xd7\x21\xf9\x24\x49\x89\x5d\xbc\xfa\x40\xb9\x0a\x4e\x9f\xe3\xee\x7b\x36\x24\x49\x5a\x4c\xa1\x8e\x9f\xf1\xfb\x25\xcf\x47\xdd\xd6\x6e\x4e\x5f\x87\x30\x57\x19\xbc\xf4\x8a\x30\xca\xb3\x34\x0d\xc9\x0f\x77\x43\x4b\xdf\x83\x74\x5b\x07\x1b\xa4\xd1\x13\xee\xb9\x28\xb5\x7f\xff\x8e\xe5\x46\x3b\x8f\xff\x2f\x74\x07\xf6\xa8\x77\x73\xd1\xff\x73\x1e\x3f\xfd\x4e\xc8\xc2\xbc\xe4\xb3\x0d\x71\x5c\xaf\xf1\x26\x2c\x92\x2e\xb1\xde\xee\x91\xdd\x63\xb7\xff\xba\xde\x3e\x60\x99\xf6\x3d\x3b\x70\x7b\x89\x11\x67\x65\xa8\x79\xf3\x23\x02\x3d\x3a\x6f\xa5\x2e\xa3\x8d\xaf\x10\xcd\x8a\x2b\xdc\x86\xe3\x88\x73\x7b\x37\xbc\x19\x11\xfa\x11\x61\xd4\xce\xee\x23\xcc\xbf\x10\xbe\x71\xd5\x4e\x0c\x83\xfa\x63\x44\x79\x95\xe1\x4a\x75\x75\x65\xc2\x4b\x86\x8b\xc3\xdf\x00\x00\x00\xff\xff\xf3\x2e\xa3\xe5\x0f\x04\x00\x00"),
		},
		"/transport": &vfsgen۰DirInfo{
			name:    "transport",
			modTime: time.Date(2019, 8, 1, 11, 6, 44, 716897420, time.UTC),
		},
		"/transport/grpc": &vfsgen۰DirInfo{
			name:    "grpc",
			modTime: time.Date(2019, 8, 27, 18, 45, 14, 527830158, time.UTC),
		},
		"/transport/grpc/client.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "client.go.tmpl",
			modTime:          time.Date(2019, 8, 23, 18, 7, 12, 535179892, time.UTC),
			uncompressedSize: 158,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x49\x2d\x2e\xe1\xe2\xca\xcc\x2d\xc8\x2f\x2a\x51\x50\x4a\xce\xcf\x2b\x49\xad\x28\x51\xe2\xe2\xaa\xae\xd6\x55\x50\x29\x28\xca\x2f\xc9\x4f\x2a\x4d\x53\xb0\xb2\x55\xd0\x0b\x80\x71\x6a\x6b\x21\xb2\xc9\x39\x99\xa9\x79\x25\x60\x39\xdf\xfc\x94\xd2\x9c\x54\xb8\x4c\x49\x41\x0e\x58\x38\x24\x35\xb7\x20\x27\xb1\x04\x2c\xc1\xa5\xaf\xaf\xe0\xe1\xea\xe3\xe3\xaf\xe0\xec\xe3\xe9\xea\x17\x02\xe2\x57\x57\xc3\x4c\xd1\x0b\x4a\x2d\x29\xaa\x74\xce\x2f\xcd\x2b\x01\x29\x06\x04\x00\x00\xff\xff\xe2\x0b\x91\x64\x9e\x00\x00\x00"),
		},
		"/transport/grpc/server.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "server.go.tmpl",
			modTime:          time.Date(2019, 8, 27, 18, 45, 14, 523830163, time.UTC),
			uncompressedSize: 2377,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\xc1\x6e\xe3\x36\x10\x3d\x4b\x5f\x31\x35\x8a\xc2\x5a\x38\xd4\x3d\x40\x4e\x4e\xb0\xed\x61\xe3\xc0\xf1\xbd\xa0\xa5\xb1\xc2\xae\x44\x32\x24\x95\xdd\x40\xf0\xbf\x17\x43\x52\xb4\x62\xd9\x45\x7a\xd9\x83\x61\x92\x33\xc3\x79\x33\xef\x71\x54\x96\xb0\x56\x35\x42\x83\x12\x0d\x77\x58\xc3\xfe\x1d\x1a\x55\x0b\xc9\xe0\x7e\x03\x8f\x9b\x1d\x3c\xdc\xff\xb5\x63\xb9\xe6\xd5\x77\xde\x20\x34\x46\x57\x79\x3e\x0c\x37\xf0\xbb\x36\xea\x1f\xac\x1c\xdc\xde\x01\x7b\x8a\xeb\xe3\x31\xd9\x9c\xda\xf7\x87\xd1\x18\x36\xa3\xd5\xa2\x79\x43\xe3\x6d\xdf\x54\xdd\xb7\x98\x2c\x4e\xb7\xfe\x78\x87\x9d\x6e\xb9\xf3\x86\x5c\x74\x5a\x19\x07\xcb\x3c\x5b\x54\x4a\x3a\xfc\xe9\x16\x79\xb6\x40\x63\x94\xb1\xb4\x6a\x84\x7b\xe9\xf7\xac\x52\x5d\xd9\xa8\x9b\xef\xc2\x95\xf4\x73\x86\x4b\x4b\x81\x25\x61\x5e\xe4\x79\xa6\xf7\xb0\x18\x86\x04\x3c\xe1\x3a\x81\x20\xaf\x0f\x2e\xc9\x52\x0a\xe9\xd0\x48\xde\x96\x28\x6b\xad\x84\x74\x8b\xbc\xa0\x46\x80\xe1\xb2\xc1\x53\xc5\xec\x19\xcd\x9b\xa8\xd0\x12\x74\xf7\xae\x11\x86\x01\xd8\x23\xef\xf0\x78\x7c\x0e\x75\x5b\x67\xfa\xca\xc1\x90\x67\x29\x9c\x6d\x9f\xd6\x3e\x22\xcb\x46\x77\x38\x1e\xff\xe4\xb2\x6e\xd1\xf8\xa6\xb3\xb8\xa1\xa0\x1b\x40\x59\x93\x37\x75\x2d\xb4\x53\x54\xe8\x83\xa8\x79\x31\x3a\x3f\xf4\xb2\x82\x47\xfc\x31\xb9\x31\x20\x58\x8e\x35\x58\x18\x57\xec\x19\xdd\x0a\x94\x76\x42\x49\x0b\x8c\x31\x9f\x33\xb8\x6f\xfc\x69\x01\x7a\xcf\x66\x57\x51\x15\x06\x5d\x6f\x24\xfc\x31\x33\x0e\xa1\x9c\x4f\x94\x78\x1b\x6a\x7c\xc4\x1f\x11\x61\x9e\x65\x59\x42\x39\x4d\xfb\x10\x0f\x57\xe4\x71\x8f\x95\xaa\x71\x62\xdd\xe2\x6b\x8f\x36\x18\x1f\xe4\xcc\x68\xb5\x92\x16\xbd\x35\x96\xca\x18\xa3\x6d\xb1\xf2\xb0\x52\x63\x33\xea\x6d\x3e\xc7\x1e\x54\x6a\xf0\x95\x1a\xbd\x3c\x91\xfe\x15\xdd\x37\xb4\x96\x7b\xdf\x80\x61\xf7\xae\xb1\x48\xc2\x36\x68\xff\x33\x24\x20\x4b\x31\x79\x59\xc2\x04\x39\x08\x0b\xee\x05\x81\x57\xae\xe7\x2d\xbc\x44\x5d\x88\x4e\xb7\xd8\xa1\x74\x9c\x6a\x01\x75\xf0\x4e\xdb\xa7\x35\xd4\x78\x10\x12\x6b\x10\xd2\x1f\xa5\xa4\x41\x13\x4b\x6b\xde\xe0\xcb\xb9\x74\x46\xd6\x8a\x69\xe6\x65\xe5\x7e\x42\x7c\x75\x6c\x1d\xfe\x57\x60\x42\x89\xf0\x25\x48\x82\x1a\x32\x06\x14\xb0\x3c\x9d\xda\xf1\x74\x05\xfe\xb9\x16\x24\x97\xbf\x29\x3e\x32\x41\xc7\xd4\x17\x6b\xde\xd8\x5c\x15\x41\x7f\x5f\xb7\x4f\x6b\x82\x91\xd2\x16\x79\x26\x0e\x3e\xf2\xb7\x3b\x90\xa2\xa5\x4b\x01\x00\xca\x12\x76\x9b\xfb\xcd\x2d\xa0\x67\x3e\xa4\x0c\xa6\x28\x51\x29\x5a\x9f\x92\xf8\x1d\x65\x3b\x62\x61\x97\x70\x17\x2b\x8a\xc9\x03\x89\x51\x1b\x17\x64\xf1\x0b\x55\x71\x4d\xf2\x24\x91\x58\x8a\xd8\xb7\x48\x62\xa8\xc9\x55\xc8\x06\xb8\x04\x21\x2b\xd5\xd1\x3a\x4d\xe6\x33\xde\xa0\x0b\x69\x29\x87\x90\x4e\x01\x87\x5a\x75\x5c\xc8\x1b\xab\xb1\x12\x07\x51\xc5\xb9\x15\x35\x74\x0d\xc7\x65\xc1\xe8\x7d\x82\x49\x73\xf4\xc0\x2b\x1c\x48\x2a\x93\xdd\x54\x22\xe2\x30\x89\xb8\x4b\x24\x67\x67\x34\x2a\x63\x69\x64\x2c\x17\xde\x3e\x83\xb2\x28\x22\xcf\x9e\x97\x74\xe1\x48\x74\x1c\xca\xdb\x24\xaa\x28\xaf\x24\xca\x73\x97\x50\xb2\x59\x1a\x7c\xbd\x24\xc1\x73\x74\x1f\x45\x16\xaf\x0e\x6a\xa2\x26\x5f\x1d\x4f\x17\x98\xf4\x82\x26\xf6\x54\xef\x1a\x45\x8b\x33\x9d\x46\x6e\xec\x84\xbd\x44\x74\x24\x36\xf2\x36\xa6\x4d\xa5\x85\xa4\xd7\xde\xf9\x08\xe9\x93\xac\xa5\x80\x4f\x93\x76\x86\x23\x71\xe6\x1f\xc6\xe9\x6d\xa6\xef\xd4\x30\xa4\xba\x8f\xc7\x82\x3e\xea\xdb\xf3\x61\x32\xbb\x34\xd4\x4c\xc4\xd9\xff\x4b\xdc\xf4\xfa\xd9\x24\x98\x2c\xff\x0d\x00\x00\xff\xff\x85\xf9\x71\x1f\x49\x09\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/endpoint"].(os.FileInfo),
		fs["/protobuf_example.go.tmpl"].(os.FileInfo),
		fs["/service"].(os.FileInfo),
		fs["/transport"].(os.FileInfo),
	}
	fs["/endpoint"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/endpoint/endpoints.go.tmpl"].(os.FileInfo),
		fs["/endpoint/request_response.go.tmpl"].(os.FileInfo),
	}
	fs["/service"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/service/interface.go.tmpl"].(os.FileInfo),
		fs["/service/models.go.tmpl"].(os.FileInfo),
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
