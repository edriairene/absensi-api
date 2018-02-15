// Code generated by go-bindata.
// sources:
// templates/footer.tpl
// templates/header.tpl
// pages/form.html
// pages/report.html
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _templatesFooterTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x50\x50\x50\xb0\xd1\x4f\xc9\x2c\xb3\xe3\xb2\xd1\x4f\xca\x4f\xa9\x04\xd1\x19\x25\xb9\x39\x76\x5c\x80\x00\x00\x00\xff\xff\x1f\x13\xcc\x82\x1b\x00\x00\x00")

func templatesFooterTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesFooterTpl,
		"templates/footer.tpl",
	)
}

func templatesFooterTpl() (*asset, error) {
	bytes, err := templatesFooterTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/footer.tpl", size: 27, mode: os.FileMode(420), modTime: time.Unix(1518580855, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeaderTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\x5d\x73\xea\x36\x10\x86\xef\xf3\x2b\x54\xdd\xe6\x60\x01\xe6\xb3\x83\x99\xf1\x49\x30\x05\x1a\x8a\x71\x08\x24\x77\x42\x96\x6d\x19\x7d\x38\x92\x30\xb8\xbf\xbe\x03\x84\x4c\xa6\x4d\x3a\x99\x73\x65\xef\xbb\xd2\xce\x3e\xfb\x8e\x76\xf0\x5b\xac\x88\xad\x0a\x0a\x32\x2b\xf8\xf0\x66\x70\xfa\x00\x8e\x65\xea\x41\x2a\xe1\x49\xa0\x38\x1e\xde\x00\x00\xc0\x40\x50\x8b\x01\xc9\xb0\x36\xd4\x7a\x70\xf5\x18\xd4\x7a\xf0\x63\x4a\x62\x41\x3d\x58\x32\x7a\x28\x94\xb6\x10\x10\x25\x2d\x95\xd6\x83\x07\x16\xdb\xcc\x8b\x69\xc9\x08\xad\x9d\x83\x1f\x80\x49\x66\x19\xe6\x35\x43\x30\xa7\x5e\xe3\x07\x30\x99\x66\x72\x57\xb3\xaa\x96\x30\xeb\x49\x75\x2d\x6d\x99\xe5\x74\xe8\xff\x8c\x46\xf3\x68\x32\x40\x97\xf0\x92\xe2\x4c\xee\x80\xa6\xdc\x83\xc6\x56\x9c\x9a\x8c\x52\x0b\x41\xa6\x69\xe2\x41\x64\x2c\xb6\x8c\x20\x62\x0c\x12\x98\x49\x87\x18\x03\x01\x7a\xbb\x69\x88\x66\x85\x05\x46\x13\x0f\x66\xd6\x16\xe6\x77\x84\x88\x8a\xa9\x93\xbf\xee\xa9\xae\x1c\xa2\x04\xba\xfc\xd6\x5c\xa7\xe9\x34\x1c\xc3\x99\x70\x04\x93\x4e\x6e\x20\x60\xd2\xd2\x54\x33\x5b\x79\xd0\x64\xd8\xed\xb5\x6a\xb3\xa9\xab\x9a\xf7\x33\x3b\xd9\x95\xcf\x93\x99\xbb\x1a\xcd\xff\x16\x0f\xdd\xd9\xdd\x6e\xa9\x91\x1e\xf5\x51\x58\xa4\x1d\xec\xbf\x8c\xa7\x87\xe0\xfe\xe1\x69\xee\xa3\x71\x31\x0e\x82\xbe\x9b\x6d\x8a\x71\x7b\xb6\x9b\x43\x40\xb4\x32\x46\x69\x96\x32\xe9\x41\x2c\x95\xac\x84\xda\x1b\x38\x1c\xa0\x4b\xaf\xdf\x40\xbe\x82\x08\x7c\x24\xb1\x74\xb6\x4a\x59\x63\x35\x2e\x4e\xc1\x09\xe8\x5d\x40\x2d\xa7\xee\xd4\xcf\xa3\x79\xd7\xce\x70\xe7\x19\xfd\x97\x6e\x2c\xdb\x6e\xaf\x75\x7c\x0d\x1b\x58\xad\x37\xfe\x6d\xbd\xdd\x5b\x6e\x16\xc7\x45\xda\x49\xaa\xd6\x64\x5d\x3e\xce\xb3\xfa\xa8\xd9\x71\x37\x22\x20\x53\x1e\xf9\x07\x36\x4e\x03\x7f\x8d\x62\x9f\x45\x9d\xe9\x46\x7c\x4d\xf7\x3f\x6e\xc4\x32\x37\x0e\xe1\x6a\x1f\x27\x1c\x6b\x7a\x26\xc0\x39\x3e\x22\xce\xb6\x06\x15\xaa\x28\xa8\x76\x72\x83\x1a\x4e\xa3\xe9\xf4\xd1\x5e\xc4\x57\xf1\x6b\x9b\xfc\x62\xbe\x4d\xb3\xfe\xcf\xdb\xe7\x46\x38\xb3\xa5\xbb\x94\xdd\xb5\x2b\xd2\xc5\x31\x5b\xf5\x67\x28\x22\xa1\xf1\x17\xdd\x6c\xc5\xb6\x1b\xb7\x9f\x77\x13\xbc\x0b\x16\x66\x57\x6e\xf6\xa6\x4c\x70\x7d\xdb\x0a\xbf\x6d\xd3\x67\x44\xdf\xb5\x25\xff\xb7\x2b\x9f\xb3\x4c\x5f\x96\x9d\xa8\xa0\x79\xd6\x5a\xd5\x9b\x71\x2f\xff\xcb\x76\xca\x3f\x47\x7f\x24\x14\x4d\xc3\x31\x5b\x2e\xa3\x30\x3c\x46\x49\xb0\x2e\x58\xe3\xe1\x75\xff\x14\xfb\x55\xbe\xc2\xba\x7d\xdb\xed\x2c\x9e\xee\xc4\x33\xff\x25\x96\xeb\xab\xca\x0d\x4a\x94\x16\xe7\xd6\x4e\xeb\xc3\x83\x96\x1e\x2d\xca\x71\x89\x2f\xc7\x3f\x56\x19\xa0\xcb\x12\x19\x6c\x55\x5c\xbd\x15\x8d\x59\x09\x08\xc7\xc6\x78\xf0\xb4\x26\x30\x93\x54\xc3\xe1\xcd\x3f\x01\x00\x00\xff\xff\xa0\xe6\x71\x6a\x90\x04\x00\x00")

func templatesHeaderTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeaderTpl,
		"templates/header.tpl",
	)
}

func templatesHeaderTpl() (*asset, error) {
	bytes, err := templatesHeaderTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/header.tpl", size: 1168, mode: os.FileMode(420), modTime: time.Unix(1518650320, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pagesFormHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x4f\x8f\x9b\x30\x10\xc5\xef\xf9\x14\x23\x9f\xda\x03\xeb\xa6\xea\xad\x06\x89\x2a\x2b\x2d\x6a\x9b\x5d\x25\x9c\x7a\x33\x61\x08\x96\x8c\x6d\xd9\x43\xda\xd5\x2a\xdf\xbd\xe2\x8f\x55\x92\x0a\xa9\xea\x89\x31\xef\xcd\x8f\x81\xc7\xbc\xbd\x01\x61\xe7\xb4\x24\x04\x16\xab\xc0\x5b\x94\x35\xfa\x07\x72\x9a\xc1\x03\x5c\xaf\x1b\x51\xab\x0b\x9c\xb4\x0c\x21\x65\xda\x9e\x95\x49\x9c\x3c\x23\xcb\x36\x00\x00\xa2\xdd\x46\xad\x56\xc1\x69\xf9\x9a\x7c\x02\xc2\x5f\x94\x9c\xd0\x10\xfa\xa9\xfe\xd9\x2a\x42\x96\xe5\x5f\x8e\x8f\xfb\x63\x21\x78\xbb\x9d\xbb\x17\xe8\xc6\xfa\x6e\x86\x8e\xd2\x70\x86\x0e\xa9\xb5\x75\xca\x5e\x9e\x8f\x25\xbb\x1d\xe2\xce\x3f\xf6\x68\x59\xa1\x86\xc6\xfa\x94\x19\xd5\x31\x30\xb2\xc3\xa9\xcc\xf6\xc5\x77\xc1\x47\xfd\xae\x47\x19\xd7\x13\xd0\xab\x1b\x9c\x7d\x57\xa1\x67\xa0\xea\xbf\x01\xb7\x5d\xad\x07\x7e\x77\xab\xea\x89\xac\x89\x43\x56\x64\x18\x5c\xa4\xee\x31\x65\xdb\x08\x72\x1e\x03\x1a\x4a\x26\x2b\xcb\x9e\xf2\x5d\x71\x10\x7c\x3a\xfe\x33\xee\xc3\x1a\xae\x2c\x76\xf9\x57\xf8\x3f\xe8\xc7\x35\x68\xf1\xa3\xd8\xaf\xd0\x5c\x04\x75\x18\xc2\xf0\x4f\x8c\x1f\xae\x96\x43\xd6\x82\xbb\x45\x98\x7c\x48\x6b\xce\x9c\xd7\xea\xf2\x27\xfe\x85\x49\x42\xeb\xb1\x49\x19\xf7\xe8\xac\x27\x36\x87\x32\x0f\x32\xb2\xa7\x3a\x89\x86\xc5\x7b\x64\xdf\x8a\xa7\xbc\x84\xc3\xe3\xcb\xf3\xa1\x14\x5c\xde\x3c\x2c\x5e\xc2\xc9\x2b\x47\x93\x74\x46\xda\x49\xc2\x77\xef\x3f\x6f\x04\x8f\xc2\xca\x4e\x34\xd6\xd2\x72\x27\x7e\x07\x00\x00\xff\xff\xec\xcd\x3b\x34\x3b\x03\x00\x00")

func pagesFormHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pagesFormHtml,
		"pages/form.html",
	)
}

func pagesFormHtml() (*asset, error) {
	bytes, err := pagesFormHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pages/form.html", size: 827, mode: os.FileMode(511), modTime: time.Unix(1518632024, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pagesReportHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\x5d\x8b\x9c\x30\x14\x7d\xdf\x5f\x71\xc9\xbb\x86\x85\x3e\x46\xa1\x85\x85\x4a\xe9\xec\x50\xe7\x0f\x5c\x27\x57\x13\x1a\xa3\xc4\xbb\xdb\x6e\xc5\xff\x5e\x1c\x63\x71\xd2\x6d\x9f\x3c\xe6\x9c\x73\x3f\x93\x79\x06\xa6\x7e\x74\xc8\x04\x62\x47\x93\x34\x84\x9a\x42\xce\xa3\x13\x90\xc3\xb2\x3c\x28\x6d\x5f\xe1\xea\x70\x9a\x0a\x11\x68\x1c\x02\x67\x23\x76\x24\xca\x07\x00\x00\x65\x1e\x77\x52\xdb\x69\x74\xf8\x96\x7d\x00\xa6\x9f\x9c\x5d\xc9\x33\x85\x0d\xff\x30\x96\x49\x94\x1f\x3f\xd5\x4f\xa7\xba\x52\xd2\x3c\x46\xf7\x21\x76\x3b\x84\x3e\x06\xbd\x51\xeb\x3f\xf4\xc4\x66\xd0\x85\x38\x3f\xd7\x17\x91\x54\x91\x18\x6e\x26\x87\x0d\x39\x68\x87\x50\x08\x8d\x6b\xca\x0b\xfa\xae\x43\xa7\xe4\x8d\x49\xd4\xd6\x8f\x2f\x0c\xfc\x36\x52\x94\x83\xd5\x3b\xf2\xd8\xef\xa7\x89\xcb\x04\x90\xc9\x51\xf3\xc2\x3c\xf8\xbd\xbe\x86\x7d\x5a\xeb\x26\xd8\xa3\xde\x1f\x96\xcf\x5f\x94\xdc\xf0\xa1\x7d\xb9\xb6\x17\xa7\x24\xb5\x7d\x8d\x90\xb1\x71\x74\x90\x71\x48\x4a\x61\x53\xc2\xa9\xfa\x0a\x4a\xb2\x79\x87\xfa\x8c\xda\x06\x79\xb1\x1a\xbf\xcb\xea\x97\xf5\xf7\x32\x25\x8f\xf1\xe6\x39\xa0\xef\x08\xf2\x65\xf9\x5f\x42\x5d\xce\x33\xe4\x27\xdb\xc3\xb2\x28\xc9\xfa\x7d\xfe\x1c\x68\x22\x7f\xa5\xbf\x44\x69\x4e\xf2\x3a\xe6\x53\x32\x76\xfb\xe7\xae\x1c\x5c\x08\x26\x50\x5b\x08\x29\xe2\x02\xf7\x11\xaf\x2b\xdc\x70\xb6\xcd\x59\x1c\xf7\x52\x56\x75\x05\xe7\x6f\x4f\xf1\x22\xe2\xdd\x84\xe3\xe7\x1f\x2f\xa3\x1d\x06\x3e\xbe\x8c\xdf\x01\x00\x00\xff\xff\xaf\x7a\xa2\x7c\x41\x03\x00\x00")

func pagesReportHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pagesReportHtml,
		"pages/report.html",
	)
}

func pagesReportHtml() (*asset, error) {
	bytes, err := pagesReportHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pages/report.html", size: 833, mode: os.FileMode(420), modTime: time.Unix(1518658157, 0)}
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
	"templates/footer.tpl": templatesFooterTpl,
	"templates/header.tpl": templatesHeaderTpl,
	"pages/form.html": pagesFormHtml,
	"pages/report.html": pagesReportHtml,
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
	"pages": &bintree{nil, map[string]*bintree{
		"form.html": &bintree{pagesFormHtml, map[string]*bintree{}},
		"report.html": &bintree{pagesReportHtml, map[string]*bintree{}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"footer.tpl": &bintree{templatesFooterTpl, map[string]*bintree{}},
		"header.tpl": &bintree{templatesHeaderTpl, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
