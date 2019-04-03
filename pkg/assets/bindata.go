// Code generated by go-bindata.
// sources:
// assets/README.md
// assets/config/README.md
// assets/config/app.toml
// DO NOT EDIT!

package assets

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

var _assetsReadmeMd = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x56\x48\xce\xcf\x2b\x49\xcc\xcc\x2b\x56\x48\x2c\x2e\x4e\x2d\x29\x56\xc8\x4f\x53\x48\x2c\x28\xc8\xc9\x4c\x4e\x2c\xc9\xcc\xcf\x03\x04\x00\x00\xff\xff\x01\xe6\x8a\x6c\x21\x00\x00\x00"

func assetsReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_assetsReadmeMd,
		"assets/README.md",
	)
}

func assetsReadmeMd() (*asset, error) {
	bytes, err := assetsReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/README.md", size: 33, mode: os.FileMode(436), modTime: time.Unix(1545792672, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsConfigReadmeMd = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xf6\xcf\xf6\x4d\xcc\xc9\x51\x2f\x56\x48\xce\xcf\x4b\xcb\x4c\x2f\x2d\x4a\x55\x48\x49\x2d\x4e\x2e\xca\x4c\x4a\x05\x04\x00\x00\xff\xff\x24\x65\xc8\x97\x1d\x00\x00\x00"

func assetsConfigReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_assetsConfigReadmeMd,
		"assets/config/README.md",
	)
}

func assetsConfigReadmeMd() (*asset, error) {
	bytes, err := assetsConfigReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/config/README.md", size: 29, mode: os.FileMode(436), modTime: time.Unix(1545272885, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsConfigAppToml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x51\x3d\x6f\x1b\x3d\x0c\xde\xf5\x2b\x08\x65\x75\x0e\xe7\x18\x48\xde\xd7\xa8\x8b\xb6\x68\x33\x25\x43\x91\x6e\x86\x07\x5a\xc7\xf8\x08\xeb\x44\x45\xe4\xd9\xb8\x7f\x5f\xe8\x6c\x0f\x45\xbb\xf1\x43\x7a\xbe\xb8\xc5\x9c\x23\x07\x34\x96\xb4\x73\x09\x07\x82\x0d\x78\x39\x0e\x18\xa3\x77\x27\x2a\xca\x92\xea\xa8\x6d\x96\x4d\xeb\x1d\x8e\xd6\x4b\x51\xd8\xc0\xd6\xbf\x72\xe8\x91\x22\xbc\x30\x7c\xc2\xc8\xc3\xf4\xe5\xc0\x91\x6d\x6a\x12\xd9\x67\xbf\x73\x1d\x69\x28\x9c\xed\x8a\x70\x01\x05\x3c\x50\x32\xef\x22\x07\x4a\x3a\xb3\x7d\xcd\x18\x7a\xba\x7f\xa8\xf8\x6e\xab\x54\x4e\xb4\x73\x1a\x7a\x1a\xb0\xae\x7b\xb3\xec\x1d\x76\x5d\xa9\xdd\xf2\xe1\xa9\x69\x9b\xb6\x59\x7a\x97\xa5\x18\x6c\xe0\xbf\x76\xb9\x72\x81\x8a\x3d\x73\x9c\xf1\x6a\xdd\x64\x1a\xbc\x3b\xd2\x74\x1b\x1e\x69\xba\xcc\xdc\xb6\x43\xc3\x3d\x2a\xed\xdc\x1d\xfc\x60\xeb\xa9\x80\x7f\x9d\xde\x7e\xbe\xf8\x05\xf8\x2c\x6a\x87\x42\xb5\x03\x29\xe0\xdf\x3e\x22\x1b\xad\xfc\x02\x26\x19\x21\x60\x82\x20\x29\x51\x30\x30\x81\x5f\xfc\xfd\x1b\x9c\xd9\x7a\x98\xff\x43\x2e\x62\x12\x24\x3a\x9b\xf2\xcc\x7a\xfb\xed\x7a\x51\xfb\x43\xfe\x7a\xb5\x6a\x1f\xfd\x5f\x89\x8f\x4a\xb3\xcd\x22\x62\xde\x65\x54\x3d\x77\xb5\xaf\x75\xc1\xe1\x52\xde\xc1\x73\x95\x76\x95\xaa\x1e\x24\xc5\x69\x01\x74\xf5\xd2\xb1\xe2\x3e\x52\x75\x53\xe8\x63\xe4\x42\x17\x2b\x27\x2a\xfc\x3e\xdd\xbf\x8f\x95\x48\x35\xbe\x4a\x37\x73\xdf\xde\xdf\x70\xf5\x2a\x1a\x30\x75\xe0\x8d\xbb\xbd\x5f\xc0\xa8\x04\xb8\x57\x89\xa3\x11\x64\xb4\x1e\xce\x3d\xa5\x39\x14\x35\x2c\x06\xa8\x50\x4f\xc7\x81\xdc\xbc\xde\x80\x5f\x0f\x34\x48\x99\xd6\x35\xf4\x42\x1d\xeb\xce\x25\xb2\xb3\x94\x63\xdd\x5a\xf8\xd7\x59\xd7\x8f\xab\xa7\xff\xfd\xef\x00\x00\x00\xff\xff\x2b\xee\xfe\x42\x97\x02\x00\x00"

func assetsConfigAppTomlBytes() ([]byte, error) {
	return bindataRead(
		_assetsConfigAppToml,
		"assets/config/app.toml",
	)
}

func assetsConfigAppToml() (*asset, error) {
	bytes, err := assetsConfigAppTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/config/app.toml", size: 663, mode: os.FileMode(436), modTime: time.Unix(1550725432, 0)}
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
	"assets/README.md":        assetsReadmeMd,
	"assets/config/README.md": assetsConfigReadmeMd,
	"assets/config/app.toml":  assetsConfigAppToml,
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
	"assets": {nil, map[string]*bintree{
		"README.md": {assetsReadmeMd, map[string]*bintree{}},
		"config": {nil, map[string]*bintree{
			"README.md": {assetsConfigReadmeMd, map[string]*bintree{}},
			"app.toml":  {assetsConfigAppToml, map[string]*bintree{}},
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
