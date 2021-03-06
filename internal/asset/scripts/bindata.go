// Code generated for package scripts by go-bindata DO NOT EDIT. (@generated)
// sources:
// scripts/sql/record.sql
package scripts

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _scriptsSqlRecordSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd2\xc1\x6b\xd3\x50\x1c\x07\xf0\x7b\xff\x8a\xdf\xad\x09\x78\xe8\x26\x13\x41\x7a\x48\xd3\xe7\x0c\x26\x2f\x25\x7d\x39\xec\x94\x84\xe4\xa9\x0f\xba\xb4\x64\x2f\x83\xdd\x3c\x28\x96\xa9\x38\x70\xba\xc8\x60\x8a\x0c\x06\xca\x6a\x85\x59\xa4\xed\xfe\x9b\xbd\x26\x39\xf9\x2f\x48\x16\x5d\x8b\x4e\x65\x78\x0a\x81\xe4\xf3\xfd\xbd\xdf\xfb\xaa\x16\x52\x08\x02\xa2\x34\x74\x04\x6e\x44\xfd\x6e\x14\x6c\xb8\x20\x55\x00\x5c\x16\xb8\xc0\x42\x2e\x2d\xd5\x64\x88\xc3\x0d\x76\x3f\xa4\x01\x60\x93\x00\xb6\x75\x1d\x14\x9b\x98\x8e\x86\x55\x0b\x19\x08\x13\x50\x4d\xe3\xfc\x59\xcd\x9e\x7c\x10\xef\x0f\xce\xbe\x4e\xf2\xdd\x41\xf5\x5a\xe1\x84\xde\x3a\x75\x61\xd3\x8b\xfc\x07\x5e\x24\x2d\xaf\xac\xc8\xa0\x9a\xba\x5e\xe4\xc6\xfc\xde\x4d\x27\x0e\x99\xdf\x0d\xa8\xe3\xb3\xb9\x3e\xe7\x06\x43\x71\xfa\x4a\xec\x3c\x4f\x8f\x86\x25\xc7\xb7\x7a\x0b\xdc\xf5\xe5\x2b\x69\x5f\x1e\xcd\xc6\x3b\xe9\xe7\x89\x38\x78\x5a\x6a\x9b\x5e\x27\xfe\xef\xe9\x1e\x4e\x7f\x8c\xc6\x3b\x97\xac\xac\x89\x6e\x2b\xb6\x4e\xa0\x7a\xa3\x56\xab\xce\xff\x4d\xa7\x2f\xc5\x71\x92\xee\xbe\x15\xc7\xc9\x6c\x6f\x94\xef\x9d\x94\x48\x2f\x62\xdd\x88\xf1\xad\xbf\x49\x4b\x8b\xd0\xd9\x34\x11\x8f\xfb\xe9\xf8\x48\xca\x46\xdb\x62\xf8\xe2\xe2\x3d\x1b\x6d\xe7\x1f\x13\xb9\x64\xfd\x88\x7a\x9c\x06\x8e\xc7\x5d\x08\x3c\x4e\x39\x5b\xa7\x17\xa0\x6a\x5b\x16\xc2\xc4\x21\x9a\x81\xda\x44\x31\x5a\xbf\x1d\xb1\xbf\x2f\x26\xe3\xc5\x31\xe3\x5e\x70\x15\xcf\xc4\x60\xb7\x9a\xc5\x56\xff\x99\x35\xdb\x3f\x99\xbd\x1e\x2e\x66\x05\xb4\x43\xff\x98\x75\xd9\x8d\x64\xa7\x9f\x44\xff\x5d\xfe\xe6\xb0\x54\xbe\x4d\x9f\xe5\x93\x24\x1b\x1c\x16\xdf\x9e\x93\x2d\x4b\x33\x14\x6b\x0d\xee\xa2\x35\x90\x8a\xa6\xcb\x60\xb7\x35\xbc\x0a\x0d\x62\x21\x54\x91\x01\xe1\x55\x0d\xa3\xba\x16\x86\xdd\x66\x63\x7e\xae\x3b\x8a\xd5\x46\xa4\x5e\x34\xe3\x67\x4d\xea\xbf\xd4\xe4\x56\xe5\x7b\x00\x00\x00\xff\xff\x5b\x9d\xd8\x1c\x55\x03\x00\x00")

func scriptsSqlRecordSqlBytes() ([]byte, error) {
	return bindataRead(
		_scriptsSqlRecordSql,
		"scripts/sql/record.sql",
	)
}

func scriptsSqlRecordSql() (*asset, error) {
	bytes, err := scriptsSqlRecordSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/sql/record.sql", size: 853, mode: os.FileMode(420), modTime: time.Unix(1574732821, 0)}
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
	"scripts/sql/record.sql": scriptsSqlRecordSql,
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
	"scripts": &bintree{nil, map[string]*bintree{
		"sql": &bintree{nil, map[string]*bintree{
			"record.sql": &bintree{scriptsSqlRecordSql, map[string]*bintree{}},
		}},
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
