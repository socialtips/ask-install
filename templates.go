// Code generated by go-bindata.
// sources:
// templates/Caddyfile
// templates/docker-compose.yml
// templates/setup.sh
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

var _templatesCaddyfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xcf\x6a\xe3\x30\x10\x87\xef\x7a\x8a\x1f\xf8\xba\x8e\x0f\xcd\xa9\xb7\x42\xbc\xbb\x85\x25\x81\x26\xbd\x57\x6b\x8f\x63\x11\x55\x0a\x9a\x71\xdd\x62\xfc\xee\x8b\xa5\x34\x6d\xf3\x87\xb2\x27\x33\x33\xd2\x7c\xf3\x69\x9c\x61\xb1\xc2\x72\xb5\x41\xb9\xb8\xdf\xe0\xe7\xfd\x9f\xf2\x07\xee\x1e\x37\xab\x5f\xe5\xb2\x7c\xb8\xdb\x94\x8b\x99\xca\x54\x86\x07\xea\x98\x20\x2d\x41\xf3\x2e\x37\x8e\x45\x5b\x0b\xdd\x08\x05\x50\x6d\xc4\xb8\x6d\xac\x56\xde\x35\x66\x8b\xc6\x58\x82\x76\x35\x02\xe5\xa1\x73\xa7\x17\x55\x06\xf1\xde\xa2\x37\xd2\xe2\x29\xe7\xa7\xdb\x08\xf9\xd2\x3b\x67\x95\x29\x35\x0c\x98\xfd\xf6\x2c\x4e\x3f\x13\xc6\x11\x83\x52\x80\xf5\x5b\xb0\xd4\xbe\x93\x29\x1a\x06\x98\x06\xb3\xb5\xd5\xd5\x6e\xe9\xc5\x34\xa6\xd2\x62\xbc\xe3\xd2\xe9\xbf\x96\x6a\xe4\xe3\xa8\x80\x0c\x3c\x9d\xa8\x27\x48\x8d\x7d\xf0\xaf\x6f\x0a\xe9\x8b\x62\xca\x15\x68\x45\xf6\xb7\x45\x11\xcf\x15\x18\x14\x80\x38\xa2\xef\x24\x1d\x89\x19\x09\xda\xf1\x5e\x07\x72\xa2\x80\x31\x4e\x90\x83\x2c\xd3\x91\xf4\x2d\x22\x05\xff\x49\x70\x35\xc6\x51\x45\x00\xd9\x5d\xeb\x83\x3b\x65\x1c\xd2\x47\xcc\x21\xbe\x9d\xcf\xe7\xf3\x33\xdc\xa1\x78\x91\x98\x2c\x3a\x69\xcf\x2c\x3a\x69\x3f\x2c\x62\x70\x6a\xd1\x49\x7b\xad\x67\x5a\x95\xf3\x82\xd9\x23\xd3\xfa\xe6\x63\x35\xa2\xc5\x54\xf1\xb7\x61\x05\x04\xef\x05\xc5\x8b\x0e\x45\xdf\xf7\x97\x1e\x78\x61\x38\xed\x56\x33\xd6\x37\xe8\x35\x83\xd2\xb6\x63\xf9\xac\x61\x76\xb9\xe5\xa7\x17\xad\xf4\x27\xcb\x77\xc1\x4a\xbf\xbd\xfb\x5d\x19\x7d\xba\xba\xa6\xf0\x42\x8c\xde\xd4\x5b\x12\xb4\xf2\x6c\x63\x81\x5e\x2b\xda\x0b\x8a\x94\xe7\x43\x97\x2f\x22\xdf\xaa\x24\x99\x2b\x80\xec\x3a\xe2\x28\x76\x61\x0d\xa3\xfa\x17\x00\x00\xff\xff\x70\xa5\x12\xc5\xf4\x03\x00\x00")

func templatesCaddyfileBytes() ([]byte, error) {
	return bindataRead(
		_templatesCaddyfile,
		"templates/Caddyfile",
	)
}

func templatesCaddyfile() (*asset, error) {
	bytes, err := templatesCaddyfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/Caddyfile", size: 1012, mode: os.FileMode(420), modTime: time.Unix(1493312981, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\x4d\x73\xe2\x38\x13\xbe\xfb\x57\xa8\xc2\x61\x2e\x2f\x38\x33\x99\x77\x6b\xd7\x55\x1c\x1c\xa3\x01\x2f\x8e\x4d\x59\x66\x32\x73\xf2\x08\x5b\x80\x17\x47\x62\x25\x39\xd9\x54\x8a\xff\xbe\x25\xd9\x80\x30\x1f\x43\xcd\x26\xa9\xa4\x6c\xf5\xd3\xad\x56\xf7\xd3\xad\x76\x07\x0c\x22\x10\x46\x09\x80\x03\x3f\x01\x5f\xfc\x00\xfe\x0f\xb8\xd3\x24\x1a\xc2\x10\xc6\x6e\x02\x07\x3d\xab\x63\x75\x40\x4c\x2a\x41\x80\x5c\x12\x80\xc5\xaa\x5b\x50\x21\x71\x59\x02\x3c\x97\x84\x03\x92\x17\xb2\xa0\x0b\x2d\xcd\x18\x9d\x17\x0b\x30\x2f\x4a\x02\x30\xcd\x01\x27\x5d\x5e\xd1\xb6\xa2\xd5\x01\x92\xb1\x12\xbc\x14\x72\x09\x7e\x74\xc5\x0f\x47\x6f\x72\x60\xbb\x2b\xac\x8e\x65\x3d\x13\x2e\x0a\x46\x1d\xf0\xe1\xd3\x07\xcb\xea\x9c\xfd\xb1\x3a\x00\x11\xfe\x5c\x64\x44\x5c\x42\x59\xa2\x01\x39\x96\x05\x40\x86\x5f\x1d\x0b\x00\x00\x8a\x27\xbc\x20\x0e\xc8\x18\xc7\xe5\x9a\xb3\xbf\x48\x26\x6d\x25\x7c\x7b\x03\x3d\x6f\x89\x29\x25\x25\xd8\x6c\x34\x94\x13\x21\x31\x97\x0e\x60\xb4\x3b\xc7\x45\x59\x71\xe2\x7c\xbc\xd5\x22\x42\x9f\x0b\xce\xe8\x13\xa1\xb2\x36\x0b\x40\x17\xdc\x0c\x5d\x7f\xd0\x57\x86\x86\x8c\x2d\x4a\xe2\x52\x5c\xbe\xca\x22\x13\xfe\x00\x6c\x36\x37\x7b\x9c\x3b\x4d\x46\xa9\x17\xf8\x30\x4c\x52\x7f\xd0\xcf\xf0\x6b\x5b\xa8\xfe\x45\xb1\x9f\x7c\xd7\xe6\x62\xc6\xe4\x34\x0e\xc0\x66\x63\xe3\x4a\x2e\xed\x8c\x51\x4a\x32\x69\x28\xc1\x60\x3c\x8a\xe2\x30\x9d\xc6\x41\x5b\x83\x94\xab\x25\xe3\xd4\xdc\x01\x8d\x4f\x01\xb1\x58\xe5\x6d\x3f\x60\xe8\xde\x07\x70\xd0\x97\xbc\x22\x86\x2c\x89\xa7\x28\xe9\xcf\x71\x29\x9a\x55\x4a\xe4\x0b\xe3\x2b\xb1\x0f\xc6\x0c\x67\xab\xae\x58\x92\x72\xae\xc2\xff\xf6\x06\x8a\x39\xe8\xa1\x12\x67\xab\x90\xc9\x62\x5e\x64\x58\x16\x8c\x0a\x48\xf1\xac\x24\x39\xe8\xea\x98\x0b\x25\x3f\x9f\x27\x2d\xee\xce\x58\x45\x33\xc2\x9d\x12\x4b\x22\xe4\xaf\x65\x6a\x12\xc5\x49\xff\xf7\x5b\xe3\x48\x28\x70\xbd\x7d\x58\xb4\xa3\x23\xc6\x56\x87\x89\xab\x41\xde\xc8\x0d\x43\x68\x00\xf7\xbc\x39\x36\x88\x60\x1c\xba\x0f\xb0\x8f\xc5\x6a\xc6\xe4\x91\xdc\xf7\xa2\x30\x85\x0f\xd1\x9f\x7e\xdf\xf9\xbb\x22\x42\x05\xc5\x31\x23\xed\xc6\x43\x98\x68\xbf\x96\x52\xae\x1d\x5b\xa7\xc9\x36\x10\x0f\xee\x64\xe2\x87\x43\xd4\xbf\x22\x13\x6a\x21\x27\x6b\x42\x73\x91\x32\xba\x87\x28\x9b\x3a\x4b\x5d\x40\x68\xae\xf8\x6f\xe9\x0a\xcd\xcf\xe7\x42\x4b\xdf\xa3\x68\x14\x1b\x47\x11\x4a\xfa\xb7\x3d\xfd\xeb\x1c\xa4\x45\x49\x83\x68\x38\xf4\xc3\x61\x1a\xc0\xaf\x30\xe8\x7f\x6c\x49\x1f\xa2\x70\x18\xa5\xd3\xd8\xef\x3f\x31\xba\x60\xf9\xcc\xb1\x6d\x7d\xda\xae\x7e\xb7\x15\x49\x5a\x2a\x35\xab\x53\x2f\x8a\x51\x3f\x89\xa7\xb0\x25\xd6\xcc\x9f\x4c\xef\x03\xdf\x4b\xc7\xb0\x2e\x41\xb7\x92\xcb\x49\x35\x2b\x8b\x6c\x4c\x5e\x5b\xd5\x8c\xc6\x69\x0c\x3d\x77\x92\x78\x23\x37\x45\xd0\x8b\x61\x52\xd7\x16\xc9\xf0\x5a\x66\x4b\x8c\x48\xc6\x89\xdc\x69\xfd\x52\x8a\x8c\x23\xa9\xe4\x34\x55\x7d\x3e\x3f\x5b\xc0\x7b\xa5\xa8\xc5\x40\x93\xc8\x77\xe9\xfd\xd4\x1b\x37\x87\x46\x77\xf7\x55\xb6\x32\x4e\xbb\xc3\xb8\x08\xee\xeb\xeb\x0e\xd2\x7c\xcd\x0a\xda\xc2\xb9\x8f\x28\x8d\xe1\xd0\x8f\xc2\x3a\xea\x8f\x28\x26\x8b\x82\xd1\x63\x94\xeb\x79\x10\x21\x95\x9f\xb4\x69\xba\xee\x23\x72\xb3\x8c\x08\x31\x26\xaf\x47\x1d\xf7\x40\xe3\x08\x7e\x08\xde\x25\xf3\x62\x16\xeb\x3e\xa2\x58\x82\x46\x70\x70\x70\xbc\x62\x0e\x28\x93\xa0\x37\x15\x04\xdd\x81\xcd\xa6\xd5\x68\x5f\x8a\x7c\x41\xa4\xb0\xdf\xde\x9a\x6a\xab\x0d\x9e\x50\xd4\xeb\xcf\xac\xac\x9e\x88\xc1\x96\xde\x2e\xb9\x76\x25\xb8\x2d\x78\x66\xe3\xf5\xda\x56\xf7\x1d\xe1\x5b\xeb\x5b\x93\x4d\x3d\xff\xb7\xd6\xa0\xba\x41\x25\x97\x17\x6e\x50\xf5\xd2\xd5\x98\xf7\x20\x9c\x17\xc5\x6e\xd0\xd4\xe1\x51\xa7\x36\x84\xfa\x1e\x4a\x27\x71\xf4\xed\x7b\xbb\x90\x4d\x54\x34\x86\x61\x0a\xbf\x4d\xfc\xf8\x7b\x9a\xf8\x0f\xb0\x7f\xf7\x1b\x58\xf2\xd3\xe0\x6d\x3b\x09\x8c\x76\xa2\xce\x75\xba\x9b\x98\x8e\xc6\xfe\x57\x37\x81\x87\x1d\x83\x17\xcf\x58\x92\x23\x86\x99\x6a\x57\xf7\x19\x43\x09\x41\x84\xfc\x28\x34\x9b\x0d\x22\x42\x4d\x4f\xa7\x48\x6a\x28\xba\x41\x10\x3d\xc2\x41\x33\x7c\x20\x35\x7a\x80\x16\x3b\x33\x5c\x96\x8a\x17\xa7\x0d\xc4\x51\x94\x9c\x1c\x1e\x2a\xb9\x3c\xad\x31\xf0\x91\xee\xb6\x5f\x62\xf7\x01\x0e\xa7\x6e\x3c\x68\x4a\x44\xb1\xfc\x91\xf1\x7c\xcd\x89\x10\x60\xb3\x51\x09\x54\x84\x2d\x05\x01\x9b\xcd\x17\x37\x40\xb0\x55\x21\x67\x08\xac\xf6\x3e\xcf\xdf\x5d\xee\x14\x8b\x2f\xcc\x8b\x00\x74\xc0\x00\x4b\x3c\xc3\x82\x88\xcb\x50\x35\xa9\xec\xdb\xf1\x41\x55\xd4\x2b\x77\xbd\x4f\x3f\x23\xfe\x51\x55\x9b\x16\xed\x1c\x4b\x6c\xe7\xb3\xab\x66\xab\xfd\x09\xdf\xc9\x13\xc3\xe0\x35\x8e\xe8\xf0\xff\x3c\xb6\x8f\x64\xa6\xe7\x76\xc2\x7f\x1e\xdd\x0c\xe7\xf9\xe1\xbc\x8e\x67\x05\x13\x6c\xae\x66\x75\x25\xba\xed\xfd\xd1\xfb\xff\xb9\xee\xe8\x29\x88\xfa\x2c\x71\x6c\x22\xb3\xfd\x6b\x03\x51\x53\xce\xc9\x46\xdb\xea\xad\xcf\x98\xdb\x2f\x2f\x2f\x07\x0d\xb5\xdd\x52\x95\x8a\xf6\xa8\x9b\x11\x2e\x85\x63\x73\xc6\xa4\xdd\xd3\x4b\x1a\xb0\x66\x6a\x79\x5f\x18\xdb\x21\x47\x95\xcf\x84\x71\x55\xa9\xe6\xf3\xcd\xa1\x8f\xda\x3f\x14\x98\xbb\xed\x2c\x7c\xfe\x7c\xa7\xfe\x6e\x4e\xfa\x75\x2a\x55\x7a\xb6\x3c\x53\x39\xd7\x5e\x06\x26\xbe\x19\x1a\xeb\x97\x26\x6a\x46\x50\x5e\xf7\x8e\x5d\x33\xfa\x37\x45\xa0\x50\x46\x0c\xb6\xd3\xe8\x05\xba\xec\xbe\x06\x41\xd8\x1c\xfa\xe2\x57\xa1\x19\x99\x26\x26\xf5\xe3\xf6\xf0\xbb\xb7\xfa\xc6\xbb\x6e\xeb\xaf\x35\x0b\x2f\xee\x6c\x30\xd5\xa4\x4c\x7d\x0d\xfe\x23\x09\xa7\xb8\x74\x80\xfe\xa8\x3a\x51\xd5\xc7\x88\xa3\x16\xd4\x86\xfc\x1b\x00\x00\xff\xff\xe5\x56\x7b\xba\xf0\x0f\x00\x00")

func templatesDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDockerComposeYml,
		"templates/docker-compose.yml",
	)
}

func templatesDockerComposeYml() (*asset, error) {
	bytes, err := templatesDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/docker-compose.yml", size: 4080, mode: os.FileMode(420), modTime: time.Unix(1493312981, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSetupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xc1\x6e\xc3\x20\x10\x44\xef\x7c\xc5\x94\x9c\x71\xee\xb9\xb6\xbd\x55\x51\xa4\xa8\x1f\x40\xf1\xa6\x46\xc5\x60\xb1\x38\x6e\x64\xf9\xdf\x2b\x70\xdc\x3a\xee\x09\xc4\xcc\x32\x6f\x76\xf7\xb4\xff\xb0\x7e\xcf\x8d\x10\x4c\x09\x8a\x84\x20\xd3\x04\xc8\x53\xef\x9c\xf5\x9f\x60\x8a\x57\x6b\x08\x4e\xdf\x28\xf2\x41\xde\x65\x29\x44\x1d\xcc\x17\x45\x65\x42\xdb\x05\x26\x74\xbd\x73\xcb\xf0\x5b\x31\x97\x27\xaa\xab\xd5\xcc\x7c\x79\x8e\xa4\x53\xfe\x3c\x35\xb4\x04\x30\x7c\x18\x0e\xbf\x16\x29\xc4\x0e\xe7\xa4\x63\x7a\x30\x55\xdb\xd4\xbe\x83\xaa\xff\x86\xe6\xf3\xbc\x7c\x69\x72\x50\x21\x18\x47\xd8\x0b\xaa\x93\x66\x1e\x42\xac\x31\x4d\x9b\x99\x63\x18\x66\x7b\xe6\xea\x99\xe2\xba\xea\x0e\x05\x99\x0a\x4b\x16\xff\x71\xc4\xde\x43\xa9\xd8\x42\xf7\xa9\x81\x71\xf6\x9e\x0d\x75\x81\x52\xd4\x6a\xeb\x20\xc7\x11\xd5\x6b\xb9\x4e\x93\x84\x52\xdd\x42\x53\x94\x15\x5b\x16\xbd\x6e\x69\x16\x5e\x2c\x77\x4e\xdf\x8e\xf9\x61\x9a\xe4\xb6\xed\x3b\x53\xc4\xa0\xd7\x6d\x1f\x0d\x27\x47\x3a\x23\x52\x1b\xae\xb9\x82\x65\x5c\xac\x23\x68\x86\x4d\x30\xc1\x27\x6d\x3d\x97\x6e\x9d\xd3\xd6\x27\xfa\x4e\x58\xd8\xe6\xdd\x91\x2f\x2b\xfb\x09\x00\x00\xff\xff\x1b\x30\xfc\xf6\x2e\x02\x00\x00")

func templatesSetupShBytes() ([]byte, error) {
	return bindataRead(
		_templatesSetupSh,
		"templates/setup.sh",
	)
}

func templatesSetupSh() (*asset, error) {
	bytes, err := templatesSetupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/setup.sh", size: 558, mode: os.FileMode(420), modTime: time.Unix(1493312981, 0)}
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
	"templates/Caddyfile": templatesCaddyfile,
	"templates/docker-compose.yml": templatesDockerComposeYml,
	"templates/setup.sh": templatesSetupSh,
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
	"templates": &bintree{nil, map[string]*bintree{
		"Caddyfile": &bintree{templatesCaddyfile, map[string]*bintree{}},
		"docker-compose.yml": &bintree{templatesDockerComposeYml, map[string]*bintree{}},
		"setup.sh": &bintree{templatesSetupSh, map[string]*bintree{}},
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

