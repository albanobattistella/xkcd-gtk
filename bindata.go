// Code generated by go-bindata.
// sources:
// data/about.ui
// data/transcript.ui
// data/viewer.ui
// DO NOT EDIT!

package main

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _dataAboutUi = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\xc1\x6e\xd4\x30\x10\x86\xef\x7d\x8a\xd1\x5c\x51\xe2\xa4\x15\x12\x42\x49\xaa\x6a\xbb\x44\x68\x4b\x41\xb0\x20\x6e\x91\xe3\x4c\x93\x21\x89\x1d\x6c\xa7\x6d\x1e\x89\xd7\xe0\xc9\xf0\xb6\xbb\x82\x03\x6a\x6e\x9e\xf1\xcc\x37\xbf\x7e\xfd\xd9\xe5\xe3\x38\xc0\x3d\x59\xc7\x46\xe7\x98\xc6\x09\x02\x69\x65\x1a\xd6\x6d\x8e\x5f\xf7\xef\xa2\x37\x78\x59\x9c\x65\xac\x3d\xd9\x3b\xa9\xa8\x38\x03\xc8\x2c\xfd\x9c\xd9\x92\x83\x81\xeb\x1c\x5b\xdf\xbf\xc2\xbf\x8c\x8b\x38\x3d\x47\xf1\x34\x67\xea\x1f\xa4\x3c\xa8\x41\x3a\x97\x63\xe9\xfb\xab\xda\xcc\xfe\x9a\xe5\x60\x5a\x04\x6e\x72\x94\x87\x46\xd4\x3c\x77\x0e\x3b\x61\x6b\xb2\x66\x22\xeb\x17\xd0\x72\xa4\x1c\xef\xd9\x71\x3d\x10\x16\x7b\x3b\x53\x26\x4e\xbf\xff\x1f\x0e\x65\x6b\xe5\x18\x1d\x2a\x2c\xbe\xef\x36\xd7\xf0\x8d\xe9\x81\xec\xda\xe2\x51\x3e\x16\x49\x9c\xae\xcd\x2a\x33\x8e\xa4\xbd\xc3\xe2\x0a\x1c\x8f\xd3\x40\xf0\x74\x29\xf4\x59\x81\x25\xd9\x90\x85\x3b\x63\xa1\xbc\xfd\xf8\x61\xbb\x46\x7b\xa0\xda\xb1\x0f\x6a\x3b\xef\x27\xf7\x56\x88\x96\x7d\x37\xd7\x71\xa0\x09\xdb\x1b\x72\xc1\x79\x27\x1e\x7b\xd5\x44\xc1\xe9\x75\x6d\xd3\x62\xb9\xed\x3c\x16\x9b\xd3\x13\x7e\xff\x82\xf3\x24\x7d\x0d\x9f\x17\xa9\x61\x77\x44\xae\x91\x06\x56\xa4\x1d\x45\x7e\x99\x82\xb8\x72\xbf\xab\x6e\xde\x6f\xb6\xb7\x5f\xb6\x55\xf9\xe9\xa6\xba\xa8\x92\x35\x80\x9c\x7d\x67\x6c\x70\xe9\x85\xb3\x99\x78\xce\x48\x88\x98\xf8\x27\x63\x7f\x02\x00\x00\xff\xff\x65\xf4\xef\x69\x97\x02\x00\x00"

func dataAboutUiBytes() ([]byte, error) {
	return bindataRead(
		_dataAboutUi,
		"data/about.ui",
	)
}

func dataAboutUi() (*asset, error) {
	bytes, err := dataAboutUiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/about.ui", size: 663, mode: os.FileMode(436), modTime: time.Unix(1449751697, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataTranscriptUi = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x94\xdd\x8e\xd3\x30\x10\x85\xef\xf7\x29\xac\xb9\x45\xde\xb4\x80\x10\x17\x49\x56\x42\x68\x79\x00\x0a\x5c\x3b\xf6\x6c\x32\xac\x63\x87\xb1\xd3\x94\xb7\x27\xfd\x23\xa1\x35\xea\xcf\xe5\xe4\x9c\x6f\x9c\x73\x1c\x25\x7f\xda\xb4\x56\xac\x91\x03\x79\x57\xc0\xf2\x71\x01\x02\x9d\xf6\x86\x5c\x5d\xc0\xb7\xd5\xb3\xfc\x08\x4f\xe5\x43\x4e\x2e\x22\xbf\x28\x8d\xe5\x83\x10\x39\xe3\xaf\x9e\x18\x83\xb0\x54\x15\x50\xc7\xd7\x37\x30\xed\x78\xf7\xb8\x7c\x0b\xd9\xce\xe7\xab\x9f\xa8\xa3\xd0\x56\x85\x50\xc0\x97\xf8\xfa\x99\x94\xf5\x35\x08\x32\x05\x44\x56\x2e\x68\xa6\x2e\x4a\xb3\x7f\xbc\x65\x46\xaa\x63\xdf\x21\xc7\xdf\xc2\xa9\x16\x0b\x58\x53\xa0\xca\x22\x94\x2b\xee\x31\xcf\x8e\x6a\xda\x1c\x29\xee\xad\xc7\xdd\x97\x80\x3e\xa0\x6c\x50\x19\x64\x59\x29\x86\x72\x79\x09\x30\xf8\xa2\x7a\x1b\x47\x88\xea\x26\x42\xf9\x7e\xb1\xb8\x16\x19\xc8\xc4\x06\xca\x0f\x09\x42\x37\x64\x8d\xd8\xb5\xec\x94\x95\xbb\x71\x8c\x5e\xf9\xcd\xa1\x95\x54\x9b\x9f\x66\xea\x8d\xbd\xa5\x80\xc6\xb7\xbe\x46\x87\xbe\x0f\xd7\x43\x95\xe7\x6d\x77\x87\x68\x67\xc1\xa6\x70\xd3\x9c\x8a\xf2\x55\xb3\xb7\x16\xcd\x0f\x72\xc6\x0f\x30\x37\xdf\x91\x2c\x7d\x6a\xfa\xe4\x15\x6e\xe2\x77\xc2\xe1\xf0\x51\x8e\x93\x5c\x6f\xc7\x53\xf2\xce\xd7\x48\x81\xba\xe7\xe0\x59\xfe\xe5\x9f\x95\x0d\x37\x2d\x40\x43\x51\xdd\x87\x0e\xac\x3a\xd9\x7a\x33\xb2\xc3\x78\x75\xff\x47\xf3\x6c\x5f\xd5\x49\xab\xd9\xf9\x65\x9e\x19\x4f\x4c\xff\x1a\x66\xe2\x24\xe4\xd9\xec\x07\xf3\x27\x00\x00\xff\xff\xd1\x52\xa5\x53\x94\x04\x00\x00"

func dataTranscriptUiBytes() ([]byte, error) {
	return bindataRead(
		_dataTranscriptUi,
		"data/transcript.ui",
	)
}

func dataTranscriptUi() (*asset, error) {
	bytes, err := dataTranscriptUiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/transcript.ui", size: 1172, mode: os.FileMode(436), modTime: time.Unix(1449751760, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataViewerUi = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\xdf\x6f\xdb\x36\x10\x7e\xef\x5f\x21\xf0\x75\x90\x9d\xb4\x40\xd1\x07\x5b\xc5\xd2\x60\xd9\xb0\xae\x28\x96\x6c\xdd\xdb\x40\x91\x67\x8b\x33\xc5\xd3\x28\xca\x8e\xf7\xd7\x8f\x94\x14\xc7\x8a\xa9\x5a\x96\xed\x26\x6e\xf2\x66\xca\xf7\x91\xbc\xef\xbe\x3b\xf1\x87\x46\xef\x6f\x53\x19\xcc\x41\xe7\x02\xd5\x98\x9c\x0f\xce\x48\x00\x8a\x21\x17\x6a\x3a\x26\x7f\xdc\xfc\x14\xbe\x23\xef\xa3\x57\x23\xa1\x0c\xe8\x09\x65\x10\xbd\x0a\x82\x91\x86\x7f\x0b\xa1\x21\x0f\xa4\x88\xc7\x64\x6a\x66\x3f\x90\xfb\x3e\xde\x0c\xce\x5f\x93\x61\x69\x87\xf1\x3f\xc0\x4c\xc0\x24\xcd\xf3\x31\xb9\x32\xb3\x2f\x42\x71\x5c\x90\x40\xf0\x31\x99\x0b\x58\x80\x0e\x17\xd5\x23\x67\x6f\x11\x99\xc6\x0c\xb4\x59\x06\x8a\xa6\xe0\x6c\x72\x11\x4b\x20\xd1\x8d\x2e\x60\x34\xbc\xfb\xd7\x6f\xcc\x61\x42\x0b\x69\xc2\x04\xc4\x34\x31\x24\x7a\x7b\x76\xd6\x15\xb2\x10\xdc\x24\x24\x7a\xe7\x41\xb0\x44\x48\x1e\x98\x65\x66\xcd\x8d\x30\x12\x62\xaa\xeb\xd9\xfa\x3c\xfc\x19\x28\x07\x7d\x61\x6d\x4a\x27\x93\xb2\xb9\xb2\xdf\xd1\x43\x1f\x20\x4f\x70\x11\x32\x89\x39\x84\x71\x61\x0c\xaa\xee\xd0\x72\xfa\x24\xfa\xeb\xd7\x0f\x97\xc1\x9f\x25\xfb\x5e\x54\xe9\xf1\x7d\xdb\xe7\xe5\x05\xde\x92\x75\x8b\x1e\x7e\xf9\x40\xa8\x05\x28\x43\x8d\x70\x5e\x25\xb6\xf5\x1f\xda\xa6\xec\x0a\x4f\x30\xc5\x29\x28\xc0\x22\xdf\x32\x6e\x6e\x96\x12\x9a\xcf\x9c\xe3\xce\xbf\xba\x2f\x29\xd4\x0c\x78\xa5\xe3\x35\x93\xa1\x07\xb9\x49\x58\x0b\x69\x55\xb8\x4a\x5d\x64\x1a\xe6\xa2\x9c\xe7\x03\x58\x4f\x26\x2b\xaf\xc4\x54\x51\x59\xc3\x98\x14\xcc\x79\x10\x24\x54\x71\x09\x7a\x4c\x3e\xd7\x63\x7e\xc0\x54\xb0\x87\x9e\xb5\xd3\xb2\x41\x8d\x48\xe9\x74\x25\x3e\x5f\x37\x3e\x92\xda\x89\xf2\x93\xf5\x8b\x1b\xc4\xc3\xce\x5e\x0c\xf9\xc0\x82\xa1\x0a\xdd\x4f\x12\x4d\x31\xbc\x8b\x4b\x98\x2f\xd3\x18\x2d\x85\x5f\xef\x6e\x34\xac\x26\xee\x23\xc1\xaf\x0a\x2f\xc0\x6b\xdc\x47\x56\x0a\x6e\xcd\xb7\x94\xd4\x27\x3b\xde\x8b\x9c\xda\xe5\xe4\xe2\xf1\xe8\x52\xda\x34\xdc\x30\xea\x54\xf3\xd7\x74\xa6\xad\x00\x30\x3d\xc8\x2b\x60\x8b\xc2\x7e\x2f\x47\xf2\x6a\xac\x43\x15\xff\x9a\xb6\xf6\xaa\xe5\x6d\x7a\xea\x9f\x69\xed\x3a\x4a\x81\x0b\x1a\x66\x92\x2e\xa5\xc8\xad\x9c\x92\x62\x32\x91\xd0\x41\x56\xdf\x5a\x21\xbf\x81\x2a\xd6\x55\x92\xda\xf6\x6a\x91\x72\xf8\xd5\x42\x61\x97\x40\x19\x66\x38\x77\x0b\xac\x5d\x80\x2b\x50\x39\xc1\xba\xb5\xc7\x5a\xe1\xbb\x50\x99\x7d\xae\xc2\x92\x90\xa3\x0a\xcb\xcd\x81\xb2\x99\xdd\x5e\x6c\x89\x91\x35\x0a\xdd\xaa\x9b\x44\xa0\xb8\x7f\x2a\xf6\xe9\xc3\xae\x0e\xa2\x5b\x0e\x86\x0a\x99\x3f\x69\xe9\xde\xcd\xf1\x45\xbd\xa5\x7a\x39\xb2\xc2\x8a\xd7\x84\xb5\x89\x80\x2e\x4b\xb8\x93\xd1\x71\x73\xfc\xc6\x9f\xcd\x68\x6c\x46\xe1\x9a\x69\x94\x12\xf8\x97\xf5\x7d\x76\xaf\x30\x74\x49\xa5\x2a\xe6\x65\x16\x31\xf7\xce\x0e\x85\x47\x04\x3d\x04\xb0\xf5\x0d\xd5\xca\xd0\xfa\x1f\x9b\xf3\xfd\x5c\xe7\x53\x23\xef\x57\x49\xd6\xe3\x4c\x62\x4b\x34\xae\xb4\xe0\x87\x3c\x0d\x48\xa9\x9e\x0a\x5b\xa2\xce\x37\x0e\x2e\x7c\xd6\x1a\x17\x61\x6e\xb5\x66\xa5\x46\xa2\xb7\x5d\x10\x0c\x65\x91\xaa\x7b\x50\xcb\x38\x1d\x94\xf1\x91\xc6\x20\x8f\x51\x4b\x6f\xa9\xb4\x4b\x48\x3b\xb5\xae\x00\x59\xcd\xe4\x53\x91\xc6\x7b\x15\x4e\x2e\xd2\xb0\xea\x6b\x7b\xd5\xec\x5f\x40\x0c\x66\x7f\x53\x63\x28\x4b\x48\xe4\x25\xdf\xeb\x22\x4c\xcc\x36\xd4\xa1\xde\x9e\x55\x60\x1b\x09\xa4\x4a\x6a\x8f\x11\xec\x1c\xa4\x1d\x9d\xf6\x16\x49\x2b\x17\x8f\x17\x9e\x16\xdd\x1e\x36\x3c\x4f\x28\xef\x6e\xdc\x29\xe4\x49\xa5\x5d\x77\x0f\x1f\x37\xed\xea\xf3\xdd\x67\x96\x75\xfd\xa2\xf3\xdc\xb2\xee\x92\x9a\xd3\x4a\xba\xd7\x27\x92\x74\xdc\x12\xfb\xec\x72\xae\x5f\x70\x9e\x5b\xce\x7d\x14\x6a\x76\x52\x39\xf7\xe6\x44\x72\xce\x5d\xce\x1d\x33\xd0\x47\xcd\x9d\x7e\x24\xf7\xce\x9d\x43\x6c\x8c\xd7\xcf\x49\x8f\xb1\x2b\x6e\x5e\x27\xef\xbd\x29\x6e\x5c\x23\xcf\xdd\x69\x10\xf3\x5f\x22\xef\xb8\x9b\xee\x72\x94\x88\x1c\xe4\xc5\xd1\xce\x0d\x4d\x79\xbf\x77\x85\x81\xc1\xc1\x60\xd0\xf3\xa8\xe4\x09\x79\xe2\xbe\x42\x08\x6e\x34\x55\x39\xd3\x22\x33\x7d\x2f\x89\xae\x13\x5c\xdc\xf7\xd2\x2c\x97\x07\xa2\xe3\x1a\x32\xaa\xa9\xc1\xa3\xec\x6a\x2b\xdd\x85\xb9\xa1\xda\x92\x72\xde\xf9\xed\x5a\xe3\x40\xf1\x1e\x28\x5b\x93\x76\x28\x46\x35\x28\x46\x2b\x88\x74\x07\xdc\x4e\x9f\x74\x9c\x92\x74\x7f\x8c\xb1\xe8\x2d\x58\xf7\x11\x4f\xd9\xc1\xa5\xa0\x12\xa7\x3b\x2a\xb6\x43\x49\x1f\x0d\xd7\xbe\xd8\xfa\x3f\x00\x00\xff\xff\x46\xac\x9e\x19\xe5\x25\x00\x00"

func dataViewerUiBytes() ([]byte, error) {
	return bindataRead(
		_dataViewerUi,
		"data/viewer.ui",
	)
}

func dataViewerUi() (*asset, error) {
	bytes, err := dataViewerUiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/viewer.ui", size: 9701, mode: os.FileMode(436), modTime: time.Unix(1449751718, 0)}
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
	"data/about.ui": dataAboutUi,
	"data/transcript.ui": dataTranscriptUi,
	"data/viewer.ui": dataViewerUi,
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
	"data": &bintree{nil, map[string]*bintree{
		"about.ui": &bintree{dataAboutUi, map[string]*bintree{}},
		"transcript.ui": &bintree{dataTranscriptUi, map[string]*bintree{}},
		"viewer.ui": &bintree{dataViewerUi, map[string]*bintree{}},
	}},
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

