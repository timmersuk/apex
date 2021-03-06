// Code generated by go-bindata.
// sources:
// bindata.go
// byline.js
// index.js
// shim.go
// DO NOT EDIT!

package shim

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(438), modTime: time.Unix(1508422851, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bylineJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x58\x6d\x4f\xe3\x48\x12\xfe\x8e\xc4\x7f\xa8\xe5\xa4\x49\xa2\x33\x06\x56\x3a\x69\x35\x88\x93\x3c\xc1\x80\xef\x82\x83\x1c\xb3\xdc\x68\x66\xb4\xea\xd8\x65\xdc\x8b\xd3\xed\xeb\x6e\x93\x8d\x16\xfe\xfb\xa9\xba\x9d\xd8\x09\x61\x66\xef\xc3\x4c\x12\x77\xd5\x53\x6f\x4f\x55\x97\x39\x39\x81\xb1\xac\x57\x8a\x3f\x96\x06\x86\xe3\x11\xfc\x7c\x7a\x76\x76\xfc\xf3\xe9\xd9\x3f\xe0\x5f\xb2\x14\x70\x83\x4b\x2d\xc5\xe1\xc1\xc9\x09\xfd\x83\x3b\x54\x0b\xae\x35\x97\x02\xb8\x86\x12\x15\xce\x57\xf0\xa8\x98\x30\x98\x7b\x50\x28\x44\x90\x05\x64\x25\x53\x8f\xe8\x81\x91\xc0\xc4\x0a\x6a\x54\x5a\x0a\x90\x73\xc3\xb8\xe0\xe2\x11\x18\x64\xb2\x5e\x59\x40\x59\x80\x29\xb9\x06\x2d\x0b\xb3\x64\x0a\x81\x89\x1c\x98\xd6\x32\xe3\xcc\x60\x0e\xb9\xcc\x9a\x05\x0a\xc3\x0c\xd9\x2c\x78\x85\x1a\x86\xa6\x44\x38\x9a\xb5\x1a\x47\x23\x32\x64\xc1\x72\x64\x15\x70\x01\x74\xbe\x3e\x86\x25\x37\xa5\x6c\x0c\x28\xd4\x46\xf1\x8c\x70\x3c\xe0\x22\xab\x9a\x9c\x7c\x59\x1f\x57\x7c\xc1\x5b\x2b\xa6\x44\x0b\x67\x93\xa2\x29\x8a\x46\xa3\x67\x7d\xf6\x60\x21\x73\x5e\xd0\x27\xda\x10\xeb\x66\x5e\x71\x5d\x7a\x90\x73\x82\x9f\x37\x06\x3d\xd0\xf4\x30\x43\x41\x5a\x4c\xe4\x27\x52\x59\x3c\x8d\x55\x45\x28\x1c\xb5\x8b\xbb\xf3\xd2\xca\x91\xa5\x9a\x12\x6c\xda\x94\x59\xdb\xcb\x52\x2e\xb6\x23\xe2\xda\xc2\x15\x8d\x12\x5c\x97\x68\xf5\x72\x09\x5a\x5a\xcb\xbf\x63\x66\xe8\x09\xa9\x14\xb2\xaa\xe4\x92\xc2\xcc\xa4\xc8\x39\x45\xa7\x3f\x6e\x8a\x99\x96\x08\x6c\x2e\x9f\xd1\x86\xe6\x28\x20\xa4\xe1\x99\xab\x82\xad\x4b\xdd\x15\xbc\x3d\xd2\x25\xab\x2a\x98\x63\x9b\x43\xcc\x81\x5b\x7e\x00\xeb\x45\xa7\xc8\x13\x6d\x98\x30\x9c\x55\x50\x4b\x65\x4d\xef\x46\xed\x77\xae\xdc\x84\x30\x9b\x5e\xa5\x0f\x41\x12\x42\x34\x83\xbb\x64\xfa\x6b\x74\x19\x5e\xc2\x51\x30\x83\x68\x76\xe4\xc1\x43\x94\xde\x4c\xef\x53\x78\x08\x92\x24\x88\xd3\xcf\x30\xbd\x82\x20\xfe\x0c\xff\x8e\xe2\x4b\x0f\xc2\xff\xdc\x25\xe1\x6c\x06\xd3\xc4\xc2\x45\xb7\x77\x93\x28\xbc\xf4\x20\x8a\xc7\x93\xfb\xcb\x28\xbe\x86\x4f\xf7\x29\xc4\xd3\x14\x26\xd1\x6d\x94\x86\x97\x90\x4e\xad\xd1\x16\x2e\x0a\x67\x04\x78\x1b\x26\xe3\x9b\x20\x4e\x83\x4f\xd1\x24\x4a\x3f\x7b\x16\xec\x2a\x4a\x63\xc2\xbe\x9a\x26\x10\xc0\x5d\x90\xa4\xd1\xf8\x7e\x12\x24\x70\x77\x9f\xdc\x4d\x67\x21\x04\xf1\x25\xc4\xd3\x38\x8a\xaf\x92\x28\xbe\x0e\x6f\xc3\x38\xf5\x21\x8a\x21\x9e\x42\xf8\x6b\x18\xa7\x30\xbb\x09\x26\x13\x32\x67\xf1\x82\xfb\xf4\x66\x9a\x90\xaf\x30\x9e\xde\x7d\x4e\xa2\xeb\x9b\x14\x6e\xa6\x93\xcb\x30\x99\xc1\xa7\x10\x26\x51\xf0\x69\x12\x3a\x73\xf1\x67\x18\x4f\x82\xe8\xd6\x83\xcb\xe0\x36\xb8\x0e\xad\xd6\x34\xbd\x09\x5d\x9c\x24\xea\x3c\x85\x87\x9b\x90\x1e\x93\xdd\x20\x86\x60\x9c\x46\xd3\x98\x42\x1a\x4f\xe3\x34\x09\xc6\xa9\x07\xe9\x34\x49\x37\xea\x0f\xd1\x2c\xf4\x20\x48\xa2\x59\x14\x5f\xbb\x30\x93\xe9\xad\x07\x94\xe2\xe9\x15\x89\x45\x31\xe9\xc6\xa1\x43\xa2\xf4\x6f\x57\x69\x9a\xd8\xdf\xf7\xb3\x70\x03\x0a\x97\x61\x30\x89\xe2\xeb\x99\x2b\x42\xbc\xa5\xe0\x1f\x1e\x1c\x1e\x3c\x33\x05\xda\x28\x64\x0b\xb8\x00\x85\xff\x6d\xb8\xc2\xe1\xc0\x3d\x19\x8c\xbc\xc3\x03\x00\x80\xc6\xf0\xaa\x7f\x4c\xbf\x07\xa3\x73\xd2\x3f\x39\x21\x26\x3f\x73\xc1\x51\x64\x08\xc1\x5d\x74\x78\xb0\x90\x79\x53\xa1\x8f\x7f\x10\xcd\x34\x5c\x40\xd1\x08\xdb\xe6\x43\x85\x2c\x9f\x59\x6c\x0f\x64\x6d\x29\x38\x82\x3f\xc9\x88\x42\xd3\x28\x01\xdb\xaa\x7e\xa6\x90\x19\x74\x1a\x7b\x95\xcf\x0f\x0f\x5e\xd7\x7e\xcc\x99\xe6\xd9\x3e\x0f\xb6\x60\xfe\x9a\x3b\xbc\x80\xde\x69\xfb\x70\xe3\xa5\xc3\x9b\x70\xf1\x03\xd7\x00\x5e\x01\x2b\x8d\x3b\xea\x02\x97\xd0\x53\xde\x56\xe8\x05\x94\x63\xad\x30\xb3\x53\xf7\xfd\xa8\x3a\xa0\xfd\x91\xb5\xbe\x67\x52\x68\x59\xa1\x5f\xc9\xc7\xe1\xe0\x21\x48\xe2\x28\xbe\xfe\x08\xf3\x55\xc5\x05\xfe\xed\x0d\x12\xd7\x7d\xe3\x34\x7b\x96\xdc\xcd\x18\x85\x0b\xf9\x8c\x39\x68\x29\xc5\xc0\xb9\xfc\xe3\xa4\x74\x65\x5a\x3b\xf8\xd7\x52\xd8\xd5\xe2\xa7\x3d\xc5\x30\xa5\x92\x4b\x9b\xcc\x50\x29\xa9\x86\x03\xfc\xa3\xc6\x8c\x1c\xee\x84\x07\x9b\xac\xbe\xc1\xf1\xe9\x2b\x9b\x57\xf8\x3e\x60\x27\x0c\x8b\x46\x1b\x17\xbf\x53\xea\x03\x53\x0f\x55\xc4\xf3\xef\x15\xb6\x67\xb8\xe6\x35\x0e\x2b\xbd\x95\xbd\x4a\x53\x8e\x5c\xe1\x6d\xf1\x1b\x4d\xb7\x04\x0d\x67\x42\x15\x32\x47\x78\x3e\xf5\xcf\x4e\xe1\xc8\x35\xa7\xfe\xf9\xc8\xb1\x82\xe4\xdf\x70\x63\x8b\x15\xdd\x8f\xed\x1a\xec\x71\xd5\xa5\xc2\x59\xf0\x53\xc5\x84\x2e\xa4\x5a\xf8\x19\xab\xaa\x21\x5d\x3f\x3b\xf4\x6e\x7f\xc0\xc5\xe6\xdb\xcb\x0b\xfc\xe9\x4a\x0d\x60\xa3\x40\x90\xf6\x06\xbc\xa5\x08\x8c\x04\x6d\x64\x6d\xc3\x92\x8d\xa9\x1b\x03\x85\x92\x0b\x98\x23\x05\x3b\x6f\x8a\x02\x15\xe6\xad\xee\xb2\xe4\x59\x09\x0a\x8f\x33\x29\x32\x66\x98\x60\x06\xb5\x55\x25\xd2\x6a\x0f\x7e\xa7\x9a\xac\xd7\x05\x81\x4b\xfb\xd8\x27\x6d\x72\xd5\xff\x6d\x5d\xab\x99\x61\x06\xfd\x9e\x1b\x17\x60\x54\x83\xe7\x9d\x24\x69\x7e\xb2\xd6\xe1\x02\xbe\x7c\xeb\x9d\x3c\x21\xd6\xe1\xa2\x36\x2b\x4a\x56\x2f\x50\x7f\xe7\xe0\xe5\x05\x0a\x56\xe9\x2d\x50\xa6\xcd\xb8\x6c\xc4\x53\x28\x72\xcc\x1f\xb8\x29\xc7\x09\xf5\x68\x2b\xd6\x46\x69\xd8\x13\xda\xa0\xb4\x6c\x54\x86\x03\x0d\x28\x32\x69\x37\x21\x5e\xc0\x12\x21\x97\x62\x60\xa0\x64\xcf\x08\x52\xe0\x06\x5e\x8a\xe1\x80\x88\x34\xf0\xba\xae\xd7\x2a\xdb\x90\xd9\xb2\xdd\x4a\xae\xf1\x36\x47\xd6\xee\xbc\x31\x84\x9e\x31\x42\xcf\xa5\x5b\x2e\x0a\xa9\x40\x56\xf9\xb1\x36\xab\x0a\x5b\x1a\xe8\xb5\x12\x21\x6a\x95\x01\x17\xb4\x47\x64\x28\x8b\x35\x51\x92\xdd\x4e\x72\xdd\xd4\xb3\x0d\x17\xa0\x55\xb6\x5b\x92\xf5\xe9\xf9\x5a\xeb\xd5\x7d\xb1\x1f\xaf\x23\xd7\x11\x74\xd9\xf8\x5c\x94\xa8\xb8\xd1\xc3\x8e\xb3\xde\x1b\x9a\xba\x1b\xa9\x93\xf0\x6b\x25\x8d\x34\xab\x1a\xfd\xdf\xcc\x5a\xa8\x3f\x25\x33\xaa\x8e\xb7\x49\xb8\x47\xb9\x5e\x07\x61\xe7\x6f\x46\x6c\x99\x73\xc1\xd4\x0a\xac\xb0\x06\xa6\xe1\x3e\xbd\x3a\xfe\x85\x64\x7a\xd1\x6d\xbe\xbe\xbc\xc0\xa0\x31\xc5\x2f\x83\xb6\xc4\x94\x35\x47\x2d\x9f\x6b\xf7\xc5\xd9\x1d\x6d\x95\xaa\x83\xba\x80\x81\x6b\x84\x41\x2f\x9d\x56\x03\x2e\xdc\xa7\x6f\xe4\xcc\x28\x2e\x1e\x87\xa3\x73\xdb\x64\xa6\xf8\x65\x2d\xd8\x73\x69\xe3\x46\x2f\xb3\xfd\xdb\xe8\x3d\xd0\x0d\x5f\xfa\xaa\xaf\x1d\xaf\x33\xc7\xe9\x37\xa1\xb7\x01\xdb\x69\xd8\x36\x8b\x43\xd6\x75\xc5\xcd\xf0\xe4\xab\xfa\x2a\x5e\xbe\xaa\x97\xaf\xe2\xe4\x71\xd4\x35\x80\xe3\xb7\x95\x81\x71\x32\xb9\x6a\xfb\x5e\xd7\x4c\xe8\x36\xe9\xeb\x3c\x7e\xa7\xaf\x3e\x7c\x70\xb2\x5f\x4e\xbf\xd9\x14\x7e\x15\x5d\xfa\xdc\x5c\xd0\x25\x2f\xcc\x70\x33\xb6\x77\x30\x37\x03\xc0\xaf\x50\x3c\x9a\x12\xfe\x09\xa7\xbd\x8b\x61\x5b\xe6\xcb\x7b\x4a\xc7\x70\xf6\x0d\xfe\x7e\xe1\x2c\x7e\x39\xfd\x76\xfe\x03\x07\xbe\x3b\x28\x5c\x3c\x2e\x85\x7d\x7c\x1b\x9e\x1a\xbc\x37\xbe\xde\xf8\xe6\xc6\xe7\xd0\x3a\x31\xea\x69\xd5\x8d\x2e\x5b\x42\x76\x1d\x70\xd6\x36\xc1\xe6\xc6\xde\xdf\x4e\x9d\x6e\xbf\x9f\x3a\x18\x9a\x8e\xbb\xed\xc4\xaa\x25\x5b\xe9\x76\xca\xbb\x39\xce\xb4\x81\x61\x2d\xb5\xe6\xf3\x6a\x05\x35\x53\xf4\x76\x32\xb2\xf9\x22\xa5\x65\xc9\x2b\xfc\x5e\x89\xc8\xcc\xa6\x4a\x6b\xe2\xed\xcb\x41\x3f\xf7\xd6\x19\xfd\xc4\x6b\x40\x9a\xdf\xae\x3a\x5d\x27\xee\x9d\xfb\x2f\x2f\x56\xac\x47\x0e\xe8\x75\x67\x37\x6b\x29\x2f\xc3\xf5\xed\x63\xd3\x81\x36\xf1\xdd\x90\x19\x8d\xb6\xa6\xa4\xbd\xe7\xd0\xbd\x21\x97\xfc\xb1\x3c\x5e\x32\x83\x0a\x16\x4c\x3d\xd1\x16\xa6\x90\x65\x25\xbd\xc9\xe7\x48\x49\x23\x78\xd4\xd0\x08\x5a\xc8\xdd\x6e\xf0\x87\x01\xc3\xb3\xa7\x0e\xd1\x6e\xf4\x58\x15\x6d\x1a\xce\xbb\x13\x8d\x26\x5a\x2c\x30\xa7\x17\xf9\xe1\xa6\x68\x5b\xee\x80\x55\x7d\x87\x1b\xbd\xa2\xf6\x50\x5f\xfb\x3f\xdc\x36\xb3\x7f\x9e\xd3\x7f\xa4\x3c\xfc\x11\xb9\x8a\xaa\xd1\x65\x9f\x57\x3d\x1e\xbd\xa1\xee\x9e\x99\xe4\xc1\xe9\x2e\x8b\xed\xcb\x3e\xc2\xfa\xa2\xfa\xf8\x91\x10\xde\xb3\xbf\xae\x5c\xdf\x05\x57\xc4\x2d\x33\xbd\xf5\x74\xfb\xae\xfb\xf0\x61\xe7\xf2\xfb\xe9\x62\xaf\xe6\xd6\x0b\x41\x1b\xce\x3e\x3b\xdd\x68\xde\xbe\xcf\xbb\x0d\xd4\x8e\xf5\x37\x8e\x6c\xcc\xd0\x9a\x61\xff\xa6\x53\xca\xa6\xca\x69\x8d\x25\xee\x2c\xa4\x36\x90\xc9\xc5\x82\xd6\x71\xa6\xd1\x03\xee\xa3\x0f\x4b\x1c\x28\x6c\xd7\x4f\x26\x1c\x6f\xed\xc2\x4f\xeb\x49\x7b\xe7\x6e\x79\x4f\x2e\xef\xb8\xf2\x7f\x85\xd7\x7f\xef\xf9\x5f\x00\x00\x00\xff\xff\x1d\x42\x7a\x74\xfe\x12\x00\x00")

func bylineJsBytes() ([]byte, error) {
	return bindataRead(
		_bylineJs,
		"byline.js",
	)
}

func bylineJs() (*asset, error) {
	bytes, err := bylineJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "byline.js", size: 4862, mode: os.FileMode(438), modTime: time.Unix(1507730537, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x4d\x6f\xe3\x36\x10\x3d\x37\x40\xfe\xc3\x6c\x80\x40\x72\xd6\x91\x93\x14\xb9\x38\xd0\xa1\x6d\xd2\x6e\x8a\x76\xb7\x40\x5a\xf4\x90\xa6\xbb\x14\x39\x96\xa7\x95\x49\x95\x1c\xf9\x03\x69\xfe\xfb\x62\x48\xc9\xb1\x17\x41\xb0\x17\x1b\x22\x1f\xdf\xcc\xbc\x79\x33\x87\x07\x4b\xe5\x41\xcf\xa9\x31\x50\x82\xc7\xff\x3a\xf2\x98\x67\xf1\xe0\x63\xeb\x9d\xc6\x10\xb2\xd1\x55\x82\x55\x9b\x86\x2c\xee\xe2\x8a\x49\x3a\x8b\x90\xc3\x83\xc9\xc9\xc9\xe1\x01\x9c\xc0\x35\x56\x5d\x0d\x68\x97\xb0\x54\xbe\x90\xa3\x89\x5c\x6b\x67\x03\x83\x89\x97\x25\xf4\xec\x05\xda\x65\x71\x7d\xf3\xfd\x1f\x3f\x7d\xbc\x7b\x77\xfb\xeb\x1e\xcf\x77\xb0\x50\x2d\xb8\x19\x04\xf6\x64\xeb\x9c\xcc\x08\xd8\x81\x56\x4d\x53\x29\xfd\x2f\xcc\x3a\xab\x99\x9c\x1d\x43\x17\xd0\xc0\xcc\x79\x58\xcd\xd1\xc6\xb7\x0b\x65\x37\xa0\x9d\xd5\x9d\xf7\x68\x39\xe6\x8c\x81\x03\x28\x8f\xe0\x3a\x0e\xac\xac\x21\x5b\x3f\xa7\x17\x95\xe8\xa9\x03\x94\xf0\xf8\xb4\x97\xcc\xef\x73\x84\x46\x05\x06\x32\xa0\x98\x95\x9e\xa3\x91\x6c\xd4\x40\x0d\x93\xe7\xcc\x5a\x45\x7e\x8f\x58\x5e\xde\x8a\xc6\xf9\xb5\x62\x2c\xac\x5b\xe5\x23\x98\xc0\xf9\xd9\xd9\xd9\x08\xfe\x87\xb3\xbd\x50\x16\xd7\x02\xae\xd1\xa2\x57\x8c\x01\xc8\x04\x58\xcd\x49\xcf\x61\x45\x4d\x03\xce\x36\x1b\xa8\x10\x3c\xb6\xa8\x18\x0d\xe0\x12\xfd\x06\x2e\xfe\xbe\xbc\x00\xa6\x05\x06\xa8\x90\x6c\xbd\x25\x30\xdb\x54\x06\xc5\xfa\x10\xf9\xe8\xf1\xf0\x00\x60\x32\x81\xdf\x3c\x2e\x45\xa6\xaa\xab\x25\x14\x7a\x04\xb2\x8c\x35\x7a\x68\x3d\x6a\x0a\xf2\x68\xe5\x55\x2b\x02\xba\xce\x1a\x70\x16\x66\x8d\x53\x2c\x81\x5a\x47\x96\xc1\x76\x8b\x0a\x7d\xe8\x19\xf3\x2e\x74\xaa\x69\x36\x03\xfe\xf2\xe2\xf4\xf2\x5b\xa8\x88\xc3\x48\x10\x22\x0a\x45\x41\x7a\x69\xde\xc2\xf9\xa0\x04\x00\xcd\x20\x97\xdb\xb2\xec\x95\x1b\x41\xcc\x14\xd2\x9b\xf3\x08\x7a\x92\x9f\xad\xb0\x64\xe2\xa1\x47\xee\xbc\x85\xbb\xad\x63\xae\x0e\x0f\x9e\x76\xc5\xfd\x21\x9a\xbd\x77\x5f\xf4\x4c\x45\x56\xf9\x0d\xdc\x4e\x3e\xec\x9b\x41\x30\x50\xa6\xe9\x28\x42\xab\x56\x56\x1c\xbf\x50\x64\xb3\x31\x3c\x42\x60\x43\x6e\x0a\xf7\x59\x4b\x2d\x66\x63\xd8\xff\x7f\x80\x27\x09\x2d\x1c\x45\x60\x83\xde\x17\xce\xe6\x99\x51\xac\xb2\xf1\xd6\xb9\xb9\xeb\xb8\xed\x38\xb5\x41\xe6\xc3\x35\x58\x34\xae\xce\xb3\xfb\xda\x3d\xc0\x29\x1c\x87\x6c\x0c\x3d\x48\x2a\x91\x9f\x6f\x7a\x5a\xe1\x43\xef\x9d\xdf\x25\x44\xef\xf7\xd9\x22\x22\xcf\xee\xc3\x9c\x16\x0f\x10\xbf\xa6\x89\x56\xa0\x51\xb3\xed\x28\xae\x89\xf3\xf3\x14\x47\x44\x78\x8e\xb2\x26\xde\x0d\xa2\x9d\xc1\x31\x04\xaa\xad\x6a\x5e\x8d\xb6\x26\x9e\x82\xa0\xcb\xe3\xd0\xe3\xcb\x18\x7b\x8f\xe1\xd5\x24\x86\xbe\xbd\xc7\x95\xac\x9a\x53\x83\x0d\x2d\x48\x5c\xff\xf3\xdd\x87\xf7\xd2\x05\xd7\xf1\x7e\xe3\x5c\xc7\x50\xf6\xeb\x2a\x1f\x3a\xe0\x3a\x8e\x7c\x82\x7e\xa9\x13\x02\x4e\xb5\x88\xf9\xe2\x96\x1a\x7d\xd1\x92\x54\x54\xab\x7c\x20\x5b\x4f\xe1\xd3\x71\xf8\x94\x8d\x21\x3e\x14\xe6\x64\xea\x45\xa8\x63\x3d\xec\x37\x83\x67\x17\x41\x16\x9e\xa4\x5b\xc8\x63\x4c\xb1\x92\x87\x41\x2b\xd6\x73\x88\x7d\x1b\xf0\x2f\x45\xed\x2c\xae\x5b\xd4\x52\xb7\x75\xf6\xf4\x9f\xe0\x6c\x8c\xbc\x9f\xc6\x55\x22\x48\x43\xd0\xcf\xc8\x50\x12\x6f\x5a\x74\x33\x49\xa6\x20\x03\x6f\xca\x12\xb2\xb4\x57\xb3\xaf\x0d\x1c\xd7\xff\x29\x18\x07\xd6\xb1\xec\xdc\x5e\xfd\xaf\xcb\x21\xad\xff\x38\x51\xc3\x96\xbd\x4f\xc9\x3c\xc4\x27\x06\x1b\x64\x7c\xf1\x72\x28\xe1\x8d\x7e\xde\x03\xaf\x37\x69\x27\x6b\xd3\xb5\x0d\x69\xc5\xb2\x2f\x43\xeb\x6c\xc0\x2f\x5b\xf7\x72\xba\xb9\xc4\x8f\x7e\x1e\x47\xd1\x96\xaa\xe9\x70\x3b\x84\xbb\xd6\x7c\xa7\xac\x69\x10\xe2\x06\x0d\x83\x15\x71\xdd\x3a\xcf\xa1\x98\xa7\xcb\x72\x67\x44\x05\x37\x06\xcd\xeb\x31\xe8\xaa\xaf\x48\xf3\xba\x18\x4a\xff\x53\x11\x87\x1f\x9d\xbf\x59\xb4\xbc\xb9\x11\xf4\x2f\xce\xb5\x42\xa1\x9a\x80\x57\xcf\x62\xc6\x65\x38\xac\xf2\x74\xbe\x55\x8f\xcc\x83\x48\x5d\xf5\xf2\x0d\x63\x40\xb6\x58\x79\x62\xcc\xa3\x1d\x93\x01\x68\xb6\xc9\x7b\x59\x8f\xc8\x1c\x4d\x81\xcc\xb8\xff\x8c\xb9\x1e\x4d\x53\x6d\xc3\xa1\x76\x96\x71\x2d\xc7\x9a\xd7\x51\xb2\xd1\xdb\xec\x2f\x9b\xf5\xbb\xf6\x73\x00\x00\x00\xff\xff\xea\x35\xd2\xf8\x55\x08\x00\x00")

func indexJsBytes() ([]byte, error) {
	return bindataRead(
		_indexJs,
		"index.js",
	)
}

func indexJs() (*asset, error) {
	bytes, err := indexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.js", size: 2133, mode: os.FileMode(438), modTime: time.Unix(1508422200, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _shimGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcc\xb1\x0d\xc2\x30\x10\x85\xe1\xde\x92\x77\x78\x0b\x24\xee\x99\x81\x82\x15\x9e\xf1\x71\x58\x21\x67\xeb\xec\x20\xb1\x3d\x45\x9a\x94\xff\x5f\x7c\x29\x69\xbb\xa9\x98\x38\xa7\x40\xdb\x92\xab\x15\x4e\x62\xe9\x9b\x62\xbc\xeb\x8e\x35\x86\x18\x52\xc2\x83\xcf\x8d\x2a\xe7\xec\xde\xbe\xb5\xc8\x00\xcf\x7e\x35\x87\x1f\x66\xd5\x14\xf4\x5c\xa7\xd3\x7f\xf8\xd0\xf4\xa0\xca\x40\x33\xdc\xb9\xe7\xc2\x35\x86\x7e\x81\x62\xf8\x07\x00\x00\xff\xff\x80\xb7\x62\x2d\x82\x00\x00\x00")

func shimGoBytes() ([]byte, error) {
	return bindataRead(
		_shimGo,
		"shim.go",
	)
}

func shimGo() (*asset, error) {
	bytes, err := shimGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "shim.go", size: 130, mode: os.FileMode(438), modTime: time.Unix(1507730537, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"bindata.go": bindataGo,
	"byline.js": bylineJs,
	"index.js": indexJs,
	"shim.go": shimGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
	"byline.js": &bintree{bylineJs, map[string]*bintree{}},
	"index.js": &bintree{indexJs, map[string]*bintree{}},
	"shim.go": &bintree{shimGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

