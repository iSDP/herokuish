package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _include_buildpack_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\x51\x6f\xe2\x46\x10\x7e\x36\xbf\x62\xe4\x43\xb9\xa4\x92\x43\xe9\xf5\x29\x17\x4e\x4d\x02\xd7\xa2\x26\x21\x25\xf0\x94\x46\x68\x59\x8f\x61\x95\x65\xd7\xdd\x5d\x87\x43\x69\xfe\x7b\x67\x6d\x83\xed\x40\x7a\x2f\x27\x21\x61\xef\xec\xcc\x7c\x33\xf3\xcd\xe7\x56\x8c\x5c\x32\x83\x60\x51\x22\x77\x18\xcf\x52\xe6\x96\xfb\xa7\x8a\xad\xb0\xd5\x9a\x67\x42\xc6\x29\xe3\x4f\x51\xfe\x74\x7c\x02\x2f\xad\x60\x7b\x37\x46\xcb\x7b\xe1\xa5\x37\x00\x53\xc0\xd2\x54\x0a\xce\x9c\xd0\x0a\x32\x2b\xd4\x02\x84\xb2\x8e\x49\x89\x31\xec\xe2\xd8\xb0\x15\xa0\xb2\x99\xc1\xc8\xe7\xb5\xad\xa0\x4a\x61\xd1\x65\x29\x7c\x81\x4e\x8c\xcf\x1d\x95\x49\xd9\x34\x7a\x60\xf0\x2f\x05\x8d\x51\xb9\xba\x89\xeb\x55\x2a\x24\xd6\x6c\xa9\xd1\x3c\xa1\xa3\xc8\x6d\x52\xb4\x95\xe1\xb5\x5e\x51\x89\xee\x50\x4d\xc3\xc2\x54\xc1\x86\xc4\xe8\x15\xfc\x2e\x1c\x4c\xc7\xd7\x54\x6c\x0c\x3a\xf5\x85\x32\x09\x94\x7d\x25\x9c\x13\x76\x19\x56\x61\x32\x23\x7b\x61\xbb\x1b\x96\x56\x7a\xfe\x25\x04\xdf\x52\x7a\x7a\xf9\x74\x16\xb5\x8f\xe7\xcc\xa2\x3f\x80\x76\xf7\xe4\x75\xaf\x29\x52\x73\x0a\xed\x98\x59\xa0\xcb\x07\x44\x7e\x3b\x30\xf9\x41\xa7\xed\xbd\xc9\x51\x24\xf0\xf0\x00\x91\x82\xb0\x5d\x24\x0b\xe1\xf1\xf1\x33\xb8\x25\xaa\x56\x10\x2c\x08\x32\x97\x5a\x21\x99\x09\x54\x48\x7f\xb5\xa8\xe4\x1e\xf0\x78\xff\x2c\xf7\x5a\x22\x7f\xd2\x99\x83\x28\xfa\x27\x13\xe8\xaa\xf8\x85\x53\xd4\x9c\x14\x4a\x8b\x8d\x7c\x51\x14\x63\x4a\xc0\xbb\xef\x65\x4e\x44\x2b\x30\x2b\x88\x4c\xd2\x34\x75\x4e\x17\x3e\x49\x63\x56\x52\x58\x77\x68\x50\xd7\x74\xfe\x1e\xcd\xa4\x85\xc8\x27\x6f\xf6\xed\x4d\xe0\x9c\x73\x45\xe4\x0f\x70\xb9\x1b\x37\x7e\x4b\x89\x6d\x39\x97\x69\x1a\xf4\xa6\x8d\x83\x8b\xbb\xbb\x59\x7f\x38\xa6\x51\x10\xd5\xb7\x55\x94\xb6\x3f\x46\x37\x83\x83\x86\xf1\xe0\xaf\xe9\xe0\x7e\x32\x1b\xf6\x7b\x61\x9e\x36\x6a\x8f\x2f\x6e\xfb\xa3\x9b\xea\xca\xfd\xe4\xe2\xea\xcf\x5e\xc8\x31\x66\x26\xea\xfe\x4a\x06\x9e\x52\x5b\xa0\x0a\xd7\x39\x0d\xb7\x85\x6c\xc3\x7b\xbc\x7d\xa3\xd3\x94\xca\x4e\x8d\x78\x26\xb6\x2f\x90\xb0\x66\x16\xcd\x4a\xd3\x78\xa2\xa5\xf6\xec\xf2\xc0\x40\xe9\xb9\x8e\x37\x14\x77\xa9\xd7\x0a\xa2\x71\x79\x70\xa6\xf4\xc2\x68\xda\xb9\xbf\x69\x72\x35\xf0\xe5\x7b\x2d\x5f\x79\xc2\x19\xb1\xa2\x44\xe0\x01\x4c\x2d\x26\x99\x24\xd1\x20\xfe\xab\x85\x85\x0e\x24\xc8\x1c\xf1\xb8\x6a\xda\x15\x2d\xcc\xec\x6a\x74\x7b\x3b\xb8\x9a\xcc\x26\xc3\x9b\xc1\x68\x3a\xe9\x85\x9f\x7e\x2e\x02\xe4\x1d\xb7\x0e\x53\x98\x53\xdb\xd7\xcc\xc4\xd6\xaf\x0c\x65\x10\x73\x21\x85\xdb\xec\xf8\x9d\x34\xda\x81\xea\xb9\x41\x73\xab\x33\xc3\x71\xef\x4a\x4e\xb2\x37\xf3\xf6\x32\x52\x0c\xbc\xb6\x39\x97\xd3\xe1\x75\xff\x8e\xc6\x30\x23\xb8\x8d\xc8\x4e\x38\x92\x96\xf0\x2b\x3a\xbe\xf4\x9a\xc6\x33\xeb\x48\x0a\x76\x11\xfd\x36\x78\x00\x75\x35\xdd\x5f\xd6\xc2\xcb\xdf\xdd\x31\xbe\xe1\xe1\xbb\x11\x0c\xbf\xde\xf7\x3e\x7e\xf8\x08\x06\x59\xec\x15\xa4\x14\x0f\x38\x3f\x3f\xdf\x83\x48\xd7\xf7\x94\xac\x5a\xb4\xad\x10\x94\x60\x8f\xea\x9b\x5a\x07\x5b\x2a\xd2\x71\xa6\x76\x14\x8a\xa1\x89\xac\x33\x17\x8a\x9c\x9d\x17\xdf\x1a\x23\x4e\xc2\xdd\xca\x17\x5a\x55\xad\x5e\xef\xf8\x6d\xf5\x3f\x9d\xd0\xb5\x44\x9b\x9a\x9e\x0a\xdf\xf5\x97\xca\xe9\xe1\xb7\xc7\xd7\xf0\x33\xc4\x9a\x6e\x7e\x0f\xdf\xce\xeb\x5d\x6c\x39\x5d\x83\xe0\xe8\x08\xde\x9d\x4b\xed\xce\x9c\x1a\xfe\x44\x2f\x31\xc9\x56\x21\x4b\x35\x66\x34\xc7\x74\x80\x19\xed\x06\x5a\xff\x15\x84\x02\x12\xc6\x55\x8f\xca\xbb\x53\xc5\xe6\xf4\xef\x74\x89\x0b\x58\x93\x48\xf8\x8d\xc6\xdd\xdd\x67\x6d\xf9\x85\xdb\xea\xd4\x64\xd4\x1f\x9d\x81\x43\xaf\x7d\x09\xc1\x11\x16\xe8\xa7\x90\xa3\xb5\xcc\x6c\x80\x3a\x4d\xe7\x6b\x04\xee\x3f\xcb\x72\xcd\x36\x16\x52\x66\x2d\xb4\x69\x29\xf2\x42\x60\x2d\xdc\xd2\x8b\xbb\xb0\x36\xc3\x6d\xc1\xd4\x69\x2f\x9a\x17\xd5\xbd\x93\x46\xc5\x8d\x31\x84\x07\x78\x52\xe2\x7c\x23\x56\xd0\x10\x0e\x7a\xdb\x46\xaf\xfa\xf3\x43\x22\x17\xb3\xfb\x6e\x28\x43\x27\xf4\xf1\xfd\x7f\x90\x5f\x1a\xd6\xce\xe9\xd6\xa9\xf5\xfa\x5f\x00\x00\x00\xff\xff\x40\xb8\x14\xca\x41\x09\x00\x00")

func include_buildpack_bash() ([]byte, error) {
	return bindata_read(
		_include_buildpack_bash,
		"include/buildpack.bash",
	)
}

var _include_cmd_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x41\x6f\x9c\x3c\x10\x3d\xc3\xaf\x98\xcf\x71\xa2\xe4\x80\xf8\x58\xf5\x44\xb4\x55\xa2\xb6\xb7\xb6\x97\x1c\xc3\x46\x72\xc0\x04\x14\xaf\x59\x61\xd8\xa6\x22\xfc\xf7\xce\x18\x9b\x40\xb2\xbd\xf4\x04\x9e\x19\xcf\x7b\xf3\xfc\xec\xb0\x90\xb9\x12\xad\x84\xe8\x16\xbe\xfc\xf8\x7a\x17\x86\xf9\xbe\x88\x54\x6d\xba\xcb\x2b\x18\xc2\xc0\xa7\x0b\x69\xf2\x2d\xfb\x8e\x71\x03\xe2\x28\x6a\x25\x1e\x95\x84\xbc\xd9\xef\x85\x2e\x0c\x7b\x2b\xd4\x66\xcb\x78\x82\x01\xdf\x27\x7a\x96\xbf\x0d\x30\xae\x0d\x83\x57\x30\xb2\x00\x66\x62\x5c\xa5\x71\xcc\xc2\xf1\x0d\xcf\xd6\xad\x41\xe7\x5e\x65\xd3\xc2\x33\xd4\x1a\xdb\x0c\xff\x11\xcd\xfb\x9b\xdd\xc8\xae\xa1\x68\xc2\x20\x90\x79\xd5\x60\xe2\x99\x48\x34\x5a\x22\xc8\x53\x2b\x0f\xc0\x1e\x08\xc4\x62\x36\x6d\xb7\x42\xd2\x0e\xe7\x5f\xda\x46\x47\x48\xdf\xf5\x94\x2f\x07\x5c\x9d\xd2\xeb\x9b\xcd\xa0\x62\x50\xf6\x3a\xef\xea\x46\x83\xa0\x95\xd3\x6d\x21\x5b\xa9\xed\xa8\x98\xc6\xef\xb0\x49\x23\x9e\x8c\x98\x56\x4d\x2e\x94\xd5\xc1\xa9\xa0\x89\x2e\xbf\x5c\x8c\x72\xb5\xe6\x5b\x6a\xb6\x54\x20\x62\x70\xf1\x19\xe2\x42\x1e\x63\xdd\x2b\x05\x17\x17\x93\xaa\xda\x8d\x15\x06\x76\x6e\x3a\x9e\x94\x0f\xc2\xc4\x67\xf8\x17\xc5\x23\xdb\x6d\x6d\xaf\xf5\x8c\xb3\x72\xef\x4e\xc8\x8d\xcb\x37\xd8\x55\x1e\x91\x31\x06\xa9\x10\x90\xd7\x4a\x91\x8c\x71\xfa\x66\x58\x67\x0d\xa2\x0d\xf0\x04\x32\x96\xf1\x9b\x0c\x85\x0f\x83\xd1\x39\x67\xc2\x83\xe9\xf8\x1d\xc5\xc4\x92\x4a\x66\x4e\x9e\xcc\x9b\x48\x98\xbc\x06\x53\xd5\x65\x07\x3e\x8c\x85\xab\xf8\xeb\x2b\x74\x6d\x2f\x7d\xda\x74\xa2\xeb\xcd\xf6\xff\x30\xa8\x4b\xf0\xaa\xce\x6e\xb5\x32\x3e\x70\x8c\x67\x7c\xa5\xe3\x35\x74\x95\xd4\x38\x04\x1f\x16\xfa\x61\x1d\xdb\x8d\xb8\xfb\x86\x74\x50\x06\x51\xa8\xed\xfd\x3d\x86\x28\x07\xbb\xdd\xbc\xd1\x9d\xd7\xcf\x06\x4c\x9f\x57\xde\x11\x29\xd8\x42\xca\x3b\x66\x1b\x3a\x5a\xe5\xbb\x10\xad\x8f\x4d\xf8\x65\xa9\x23\xd2\x75\xaa\xb8\xa2\xfd\x65\xed\x3c\x31\x5b\xe3\xf6\xc3\xb5\x4d\x6d\x21\xba\x0a\x31\xd7\xbe\x72\x8d\x9c\xb5\x82\x43\x5b\xeb\xae\x04\x06\x70\x1e\x6d\x3e\x19\x38\x37\x19\xda\xcc\x0d\xb5\x82\xff\x28\xc7\x44\xc7\xc2\x98\xfe\xf1\x14\x12\xd5\xcd\x50\x0b\xac\xf7\x68\xd3\xf6\xbf\x00\x62\x26\xf5\x15\x33\xe8\xe4\x70\xff\xf1\x6a\xbc\xd4\x1d\xf0\x49\xde\x90\x74\x72\x76\xaa\xa4\x3a\x9c\xba\xc4\x77\x55\xf3\xcb\x00\x65\x91\x39\x8e\xb1\x17\xf6\x22\xd3\x40\xa7\x6e\xb2\x68\x9f\xc8\x89\x64\x01\x7f\x6a\x14\x5a\x9c\x1b\x4d\xe6\x65\xa7\xb1\xad\x21\xe9\x07\x5f\x99\x33\x50\x02\x45\xc1\x1d\x53\xd9\xc2\xdb\x03\xb5\x89\xcf\x69\xd2\x98\x5e\xa9\x29\xa8\x31\x04\xeb\x35\xc4\x11\x5e\xa3\xc5\x76\xfb\xb8\x9c\x38\x19\xc7\x44\x47\x34\x97\x7b\x3b\x92\xd9\xb9\xee\x7e\xda\x97\x87\x44\x5a\x5c\x4b\x2f\x97\x55\x25\xfc\x13\x00\x00\xff\xff\xd3\x23\x50\x2b\x44\x06\x00\x00")

func include_cmd_bash() ([]byte, error) {
	return bindata_read(
		_include_cmd_bash,
		"include/cmd.bash",
	)
}

var _include_fn_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x51\x41\x6e\xc2\x30\x10\x3c\xc7\xaf\x18\xad\x22\x01\xaa\xa2\x08\xae\x34\x3d\x56\xea\x1b\x28\x07\xcb\xac\x89\xd5\xd4\x89\x6c\x03\xaa\x28\x7f\xef\x1a\xd2\x92\x43\x55\x55\xb9\x64\x77\x66\x67\x76\xd6\xca\xfa\x4a\x87\x7d\x9c\x2f\x70\x56\xc5\x8e\x4d\xa7\x03\x63\xc7\xd1\x34\xf4\xe2\xe3\xc0\x26\x41\xc3\x1e\xbc\x49\xae\xf7\xb3\x08\x21\x1f\xde\xd9\xa7\x48\xaa\xe8\x7a\xa3\xbb\xdc\xe9\x9c\xe7\xa6\x9c\xa7\x8f\x81\x51\x2e\xf1\x89\x7d\xe0\x01\xdf\x6a\x63\x59\x1d\x41\x53\x03\x12\xa0\x65\xbd\x43\xb5\x5c\xa8\x82\x4d\xdb\xa3\x62\x50\x79\x1e\x05\xeb\x1a\x35\xbd\x7a\xba\x64\xa2\x3e\xbd\xa1\x7a\x6e\x30\xab\x9b\xfa\x3c\x04\xe7\x13\xe8\x91\xca\x25\x3d\xd1\x65\x26\x78\x0a\xc8\x5c\xc8\xa7\x2e\x2a\xa7\xca\x16\xff\x4e\x95\xa1\xe0\x86\x5c\x51\x1e\xc8\x44\xf9\xe1\xa3\xe4\xa3\x5f\x82\x45\x33\x59\x9e\xd6\xb8\x6e\x5f\xe6\xfe\xe8\xee\xbc\xed\xff\x70\x8f\x13\x7b\xba\x73\xac\x6f\x72\x26\xc4\xb6\x3f\xc5\xfe\x10\x0c\x4b\xbd\xa2\xf1\x3a\x54\x5a\x8f\x72\x3e\xbe\x18\xa4\x5a\xfc\x40\xb8\x01\xd7\xcd\x26\x80\x2a\x9c\xc5\x66\x23\xa3\x77\x49\xc2\x76\xbb\x46\x6a\xd9\xab\xa2\xb8\x25\x13\x5d\x39\xa1\x76\x1d\x2a\x8f\x87\x95\xf4\x6f\xc3\xd6\x49\x9c\xaf\x00\x00\x00\xff\xff\x5c\xaf\x12\xe0\x23\x02\x00\x00")

func include_fn_bash() ([]byte, error) {
	return bindata_read(
		_include_fn_bash,
		"include/fn.bash",
	)
}

var _include_herokuish_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x54\xdf\x6f\xdc\x36\x0c\x7e\xb6\xff\x0a\x56\xbb\x36\x6b\x00\x2f\x58\xd7\xa7\x1c\x3c\xf4\x92\x5d\x97\x62\x5d\x7a\x68\xda\x0e\x43\x52\x04\x8a\x45\xfb\x84\xd8\x92\x20\xc9\x69\x0e\x41\xfe\xf7\x52\xfe\x75\xf6\x5d\xe2\x27\x8b\xfc\xf8\x89\xfc\x48\x2a\x96\x39\x5c\x5e\x02\x9b\x9d\x2c\x2e\xce\xae\xbf\x2d\x3f\x5f\x7c\x38\x7f\xff\x89\x41\x52\x7a\x60\x6f\x19\x7c\xff\x3e\x07\xbf\x46\x15\x47\x98\xad\x35\xb0\x17\x2f\xe0\x7f\x5d\x5b\x70\x1b\xe7\xb1\x82\x13\xee\xd6\x20\x1d\xe8\xda\x83\xce\x41\x70\x8f\xc7\x30\xe2\xfa\x74\xce\x46\x91\xab\x12\xb9\x43\xa8\x4d\x61\xb9\x40\xf0\xba\x8d\x7f\x0b\xda\x42\x61\x91\x82\xed\x6f\x01\x7f\x2f\x3d\xbc\x89\x73\x19\xc7\x64\x14\x5a\x95\x1b\xe0\xc6\x5c\x1b\xee\xd7\x29\x9b\x3d\x2c\x56\xab\xeb\xd5\xe2\xcb\xd9\x71\x72\x44\xe6\x47\xb6\x45\xa1\xba\x1b\x50\xcb\xf3\x6f\x3d\xca\x57\xe6\x88\x5c\x63\xe4\x4d\x2d\x4b\x31\x60\x4f\xbe\x7e\xf8\xf8\xd7\x18\xdd\xb8\xc7\xf8\x8c\x67\x6b\x1c\xf0\xa7\x8b\xd3\xb3\xe5\x18\xdf\xb8\xf7\xf8\x0d\xcf\x6e\xa7\x77\xac\x16\xa7\xff\xec\xdd\x13\x60\x8e\x82\x63\x81\x59\xc9\x2d\x09\xa4\x8c\x95\x77\xb2\xc4\x02\xc5\x75\xed\xd0\xa6\x4c\xe9\x1b\x2d\x36\xec\x69\x48\x61\x75\x6d\x02\xa6\xf9\x21\x22\x54\xae\xb6\x98\x84\xab\xdd\xaf\xaf\xe1\x21\x8e\xaa\x5b\x21\x2d\x24\x06\xae\xe2\x28\x62\xb3\x5e\x4e\xd6\x9d\x7b\xe1\xfa\xf3\x56\x9e\xde\xb2\x15\x60\x82\x19\x4a\x64\xf1\x63\x1c\x8f\x2e\xec\x13\x15\xe8\xb2\x94\x5d\xac\xf5\x0f\x07\xc1\x0d\x0e\xbd\x97\xaa\x70\xd4\x68\x2a\x41\xf9\x1c\xd8\xcb\xe4\x8f\x37\x0e\x7e\x81\x97\xee\x4a\x75\xec\x7d\x8f\xd3\x51\xaa\xc1\x6c\x4c\x29\x33\xee\xa5\x56\x2d\x9b\xa8\x89\xa3\x00\x5b\x2b\x2f\x2b\xec\x82\xfb\xd6\xa7\xa3\xba\xc8\xbc\x0a\x01\x34\x74\x39\xc9\xe6\x20\xa7\xa1\x13\x98\x4b\x15\xe2\x6f\xc2\x5c\x12\x58\x5a\xad\x2a\x54\xbe\x23\xda\xce\x45\x3a\x91\x24\x62\xff\x69\x7b\x1b\x02\x49\x55\xcc\xbc\xb6\x9b\x3e\x93\x06\xe6\xba\xf8\xed\x9c\xa4\x13\x01\x89\xb9\x17\xaf\x9d\x2c\x28\x75\x5b\xd5\xf8\xe2\x61\x58\xd2\x5d\xad\x61\xa8\x45\x2a\xe7\x79\x59\xa2\xd8\x4e\x9c\x6b\x5a\x71\x87\xd6\x11\xdf\x73\xcd\x80\xce\x0f\x5c\x09\x70\xb5\x31\xda\x7a\x22\xe9\xad\x52\xe5\x7a\xd8\xdc\xd9\x43\xb7\xcb\xc7\x89\xc0\xb0\x45\x44\xef\xa5\x2f\xb1\x25\x6f\x40\xb3\x83\x2b\xbc\xfc\xfd\xef\x24\x7c\x7f\x1e\xc0\xec\x30\x80\xa4\x12\xa4\x65\x8b\xfa\xb1\x26\xd5\x21\xec\x07\x94\x52\xe1\x1c\x84\xa6\x42\xfb\xd7\x27\x98\x18\xa4\x29\x24\xc9\xe1\xe8\xcd\x89\x26\xe4\x07\x0d\x8c\xac\x58\x3a\xdc\x75\x42\xf3\x1d\xf4\x5c\xe4\xa6\xf7\x23\xa2\x65\xc4\x90\xc9\x78\x5d\xda\x7c\x68\x0c\x6b\x29\x0a\x29\x28\x62\x6f\xdf\x18\xcc\xde\x85\xb0\x8a\x4b\x35\xc0\x21\x41\x0d\x46\x1a\xcc\xb9\x2c\xe7\x6d\xde\x5f\x3e\x2f\x4e\x97\xe1\x95\x84\x57\xaf\xa0\xc1\xdc\xc7\x71\x94\x55\x22\xc1\xfb\x20\x69\x33\xa4\x6e\x62\xe9\x24\x9e\xc0\x12\xe5\xb6\xfd\x03\xf6\x95\x86\x31\xf4\xa5\x6b\xee\xa4\xb5\x63\xaa\xc1\x9e\x34\x7f\xcf\xf8\x3a\x96\x67\xbc\xa5\x74\x7e\x2f\x17\x57\xd6\x05\xb0\x7f\xb9\xe2\x05\x86\xa7\x77\xd8\xb9\xe0\xd8\x49\x22\x98\x12\x59\x85\xff\x7d\x7b\x81\x0a\x2d\xbd\xeb\xfb\x9e\xf6\x7f\xef\x66\x63\x75\x16\xf6\xb3\x15\x61\xd5\x9d\x5c\x23\x07\x2d\x79\x48\x06\x32\x5d\x55\x74\xde\xc9\xa3\x8f\x4c\xa8\xda\x9d\x54\x06\x17\xde\x63\xf6\xb4\xc7\x70\x1b\x66\x8a\x9c\xe1\x25\x60\xb3\x8b\xe5\xc7\xf7\x8c\x1a\x00\x34\x48\x47\x0d\xe3\xeb\x28\x9a\x5e\x41\xa8\x77\x6c\x3e\x0f\x80\xc0\x3b\xf6\x87\xf3\xc8\xdd\xa8\x4d\xfe\x9d\x7e\x35\xce\x43\xb2\x37\x19\x51\xf1\x8c\xb5\x41\xb4\x53\x8e\x67\xf1\xe3\xcf\x00\x00\x00\xff\xff\x28\x16\x82\x35\x9d\x07\x00\x00")

func include_herokuish_bash() ([]byte, error) {
	return bindata_read(
		_include_herokuish_bash,
		"include/herokuish.bash",
	)
}

var _include_procfile_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x94\xdd\x4e\xe3\x3a\x10\xc7\xaf\x9b\xa7\x18\x45\x11\xb4\xe7\xe0\x06\x6e\x0f\xea\xd1\xae\xc4\x6a\xb9\x01\x56\xdc\xb2\xa8\x72\xe3\x69\x62\xd5\x89\xb3\xb6\xc3\xc7\xb2\xbc\xfb\xfa\x23\x69\x63\x1a\xae\x12\x8f\x67\xec\xdf\xfc\x67\x3c\x49\xab\x64\xb1\xe5\x02\x49\x4b\x95\xc6\xf9\x02\xde\x92\x19\xc3\x42\x50\x85\xc0\x50\x17\xab\xf4\x3b\x1a\x28\x64\x5d\xd3\x86\x81\x36\x8a\x37\x25\x6c\xa5\x02\x0a\x2e\x14\xb5\x06\xf3\xda\x22\x6c\x95\xac\xe1\x47\x7f\x58\x7a\x38\xc3\x6d\xae\xd2\xec\xc2\x9a\x0a\x6a\x20\xcd\x68\xdb\xae\x5b\x6a\xaa\x7c\xef\x0c\x7f\xe0\x95\xd6\x82\x94\xe8\xf6\x5d\x40\x9a\xbc\x27\x07\x32\x6d\xa8\x32\x53\x64\xf7\x5d\x13\x43\x0c\x98\x11\x0c\x98\x4a\xc9\xae\xac\x00\x5f\xb0\x98\x26\xdb\x5f\xe5\x5c\x20\x9b\xc7\xa2\x0c\x50\x8b\x88\xca\xb9\x7e\x06\x45\x35\x28\xcb\x21\xeb\x33\xe8\x9a\x56\xf1\x27\x1b\x50\x22\x83\x4e\xa3\x82\x67\x6e\x2a\xb8\x46\x25\x77\x1d\x11\x7c\x87\x80\xcd\xd3\x98\x21\x44\xf2\xdf\x48\x9c\xfb\x68\x43\xa3\x21\x95\xac\x71\x64\x12\x92\x32\x62\xe3\x3f\x9a\xec\xca\x2d\x92\x59\x74\x7d\xf6\x25\xca\xc0\x25\xa5\x43\x0a\x86\x1b\x2b\x54\x7a\xc5\x75\x21\x9f\xd0\xd7\x78\x2c\xac\xb6\x7c\x7c\x0b\x0f\x0f\x40\xb6\x56\x8d\x4d\xc7\x05\xfb\x58\xc4\xc7\xc7\x4b\xab\x34\x36\xc9\x6c\x26\x64\x41\x45\x08\xb4\x2b\xff\xb5\x42\xcf\x5d\xfd\xa7\x62\x87\xfa\xef\xf0\x55\xdb\xff\x17\xaa\x4a\x0d\x58\x54\x72\x61\xaf\x9d\xb9\x1f\x48\xf7\xbe\xbd\xd8\x3d\x17\x90\xff\x21\x7b\xf3\xbf\x79\x0e\xf9\x19\xbc\xbb\x10\x85\xa6\x53\x16\x64\xcb\xf7\xd8\x3a\xc6\x5e\x2a\x14\x48\xf5\x14\x36\xc3\x2d\xed\x84\x59\x0f\xf8\xd1\x7a\x2a\x8d\xe1\xac\x28\x8d\x21\xaa\x57\x31\x44\x1f\x27\x67\xd1\xd2\x2c\xba\xc1\x11\xc1\xc9\x09\xfc\xb4\xbb\x7d\xee\x57\x61\x3f\xae\x88\x7f\x83\x99\xb6\x57\x17\x06\xd9\xba\xa1\x35\x06\x31\xa2\xd3\x3e\x11\x25\x9c\x7b\x2b\x8f\x8e\xec\x1a\x16\xbf\xbd\xa1\xc1\x42\x9b\xf4\x62\x32\x0b\x6d\x6d\x3e\xfd\x48\x41\xc7\x84\xc0\x1b\xfb\x84\x84\x86\xbd\xcf\xe2\x12\x98\xf4\xf9\xbc\xb4\x52\xb9\x47\x8e\xab\x5e\xc7\xc1\x25\xcf\xd0\x0b\xc2\x64\x83\x9e\xf1\x08\xa2\x6f\xe9\x00\xa2\x2b\xd9\x1a\x57\xd5\xa6\x13\xa2\x14\x72\x93\xcc\xea\x1d\xe3\x0a\x48\x3b\x1e\x31\xcb\x3e\x68\x69\xb3\xf2\x70\xbe\x85\x1c\xdf\x84\x4b\xfe\xcf\x52\x57\x3d\xa9\x96\x9d\x2a\xdc\xc3\x1f\xa6\x99\xc7\xaa\xa8\xae\x80\xa8\x78\x38\xf5\x8f\x32\x70\xf5\x09\x5e\xdf\xdd\x7c\x5b\x1d\x38\xec\x01\xee\x29\xd7\x92\x01\xf1\xce\x23\xc6\xd4\xfe\x8f\x1f\xe9\xda\x79\xba\x61\x59\xc9\xe7\x06\xc8\xfd\xd4\xf6\x7f\xb1\xa9\xb4\xe3\xad\x4d\xc7\x67\x46\x84\xf1\x3c\x09\x9c\xa1\xd5\xdd\x9a\x33\xd7\xd3\xf3\xfb\xaf\xb7\x57\x77\x37\xff\x5e\x9c\x9f\x9f\x2f\x5c\x21\x0e\x0e\xae\xb3\x56\x69\x97\xbd\x05\x6f\xdb\x4b\xc9\x8c\x32\xe6\x2f\xb5\xe9\xfc\xea\xb8\x1d\xdb\x84\x94\xdc\x75\x45\xf0\x49\xfb\x3f\x17\x9a\x7a\x6f\x3f\xf8\x5c\x4b\x13\xa2\x2b\x14\x02\xf2\x0d\x6f\xf2\x8d\xd3\x33\x58\x19\xd7\x74\x23\xd0\x96\x99\x6a\xfd\x2c\x15\xeb\xed\xb6\x68\x05\x92\x0d\x65\xbe\xc1\x83\xad\x91\xa4\x50\x48\x0d\x06\x2d\x83\xb1\x8b\xee\x0f\xb6\x72\xca\x86\x85\xd4\x70\x7a\xda\x2f\x03\x7e\xf8\x3f\xaa\x8c\x33\x47\x99\xc4\xf3\xd4\xd7\x62\x15\x39\x1c\xd7\x25\xda\x7f\xff\x1b\x00\x00\xff\xff\x9d\x05\xaf\xe4\x6f\x07\x00\x00")

func include_procfile_bash() ([]byte, error) {
	return bindata_read(
		_include_procfile_bash,
		"include/procfile.bash",
	)
}

var _include_slug_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x53\x41\x8b\xdb\x3c\x10\x3d\xc7\xbf\x62\x3e\x11\xd8\xe4\xa0\xe4\x4b\xaf\xc5\x85\x92\x96\xb2\x50\xda\xd2\x4d\xa0\xb0\x5d\x82\x56\x9e\xd8\x02\xd9\x12\x92\xdc\x66\xdd\xf6\xbf\x77\x64\x3b\x89\xdd\xdd\x24\xa7\x78\x34\xf3\xe6\xcd\x9b\x79\x89\x43\x91\x99\x4a\x3f\x81\xd7\x75\xbe\xb3\x22\x14\x29\x5b\x86\xd2\x2e\xe3\xf7\x22\xe4\x0d\x4b\x92\xf8\x97\xab\xd2\x1a\x17\x66\x73\xf8\x95\x4c\x32\x94\x5a\x38\x84\x0c\xbd\x4c\xd9\x6d\xfb\x02\x02\xf2\x46\x59\x8b\x59\x0b\x05\x41\xb8\x47\xa1\x35\xec\x9d\x29\x61\xfb\xf5\x23\x18\x07\x77\x9b\x77\xb7\x9f\x80\x9d\x01\x6a\xa7\x53\x36\x5d\x51\x44\xed\xe1\xfe\x1e\xd8\x74\xa6\x3d\xf0\xb7\x30\x15\xd6\xb6\x6c\xe6\x0c\x1e\x1e\x5e\x43\x28\xb0\x4a\x26\x13\x87\xa1\x76\x15\xac\x92\x09\xea\x63\x05\x61\x8c\x72\x24\x05\x80\x13\x8a\x81\x65\x86\x3f\x96\x55\x4d\x2c\x38\xa7\x52\xf7\x04\xaf\xa8\xc0\x3b\xc9\xe0\x77\x24\x08\xfc\xd0\xac\x29\x72\x6c\xc6\x22\xae\xc7\x08\x22\xc2\xc5\x94\xbd\x4a\xfe\xf4\xa2\xe4\x58\xa1\x13\x01\x5f\x92\xe5\x43\xff\x76\x55\x18\xe2\x0c\x44\xd8\x61\x45\x02\x5a\x4b\xe8\xda\x48\xa1\xc1\xaa\xbc\xd9\x19\x1b\x94\xa9\x5a\x6d\x7e\x16\x4a\x16\x6d\x14\xde\x9c\xc7\x3a\xcd\x3c\x48\x4f\x19\xe7\xb5\x47\x2e\x4d\x69\x1d\x7a\xcf\xad\x33\xb9\x13\x65\x1a\x73\x3a\xf6\x7d\x8f\xc8\x46\xe5\x95\x71\x38\xec\x44\x9a\xf2\x3d\x0d\xfc\x58\x2b\x9d\xb5\x23\x2f\x17\xe7\xcc\x91\xd2\xcf\x00\xa8\xf7\x37\xb8\x54\xd9\xb5\x8e\x92\x4e\x07\x74\x61\xfa\x0c\x05\xbe\x13\x36\xe7\x78\x90\xba\xce\x30\xbd\x59\xe4\x2a\xdc\x74\xc1\xf5\x88\x18\xeb\x82\x32\xd2\x3d\x9d\x6f\x0c\x02\xfd\x26\x8b\xe1\x9c\x3b\xaf\x1a\xa4\x53\x9b\x65\x35\xf0\xbb\x02\xce\xf9\xb4\x66\x59\x07\x9a\x79\x35\x27\x8a\x41\x05\x8d\xc0\xd6\x24\x9e\xd2\xc7\x8d\xc5\x5a\x50\xbe\x2f\x8a\x5f\xec\x74\x01\x78\xb8\x64\x8b\xf7\xed\x0b\x1c\x4f\xe4\x9f\xed\x07\xd3\x9a\x62\xf6\x65\xbb\x99\xf7\xd6\xf8\xbc\xdd\x5c\xb1\xc6\x7f\xdd\x5a\x06\x73\xbe\xe8\x8b\xa8\xf1\x75\x67\xfc\x7f\xc5\x1c\xb4\x3e\x22\x04\x7c\x33\xee\xd4\x41\x8d\xcc\x31\x7c\xef\x2c\xf1\x37\x00\x00\xff\xff\x93\x63\xfe\x2d\x4b\x04\x00\x00")

func include_slug_bash() ([]byte, error) {
	return bindata_read(
		_include_slug_bash,
		"include/slug.bash",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"include/buildpack.bash": include_buildpack_bash,
	"include/cmd.bash": include_cmd_bash,
	"include/fn.bash": include_fn_bash,
	"include/herokuish.bash": include_herokuish_bash,
	"include/procfile.bash": include_procfile_bash,
	"include/slug.bash": include_slug_bash,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"include": &_bintree_t{nil, map[string]*_bintree_t{
		"buildpack.bash": &_bintree_t{include_buildpack_bash, map[string]*_bintree_t{
		}},
		"cmd.bash": &_bintree_t{include_cmd_bash, map[string]*_bintree_t{
		}},
		"fn.bash": &_bintree_t{include_fn_bash, map[string]*_bintree_t{
		}},
		"herokuish.bash": &_bintree_t{include_herokuish_bash, map[string]*_bintree_t{
		}},
		"procfile.bash": &_bintree_t{include_procfile_bash, map[string]*_bintree_t{
		}},
		"slug.bash": &_bintree_t{include_slug_bash, map[string]*_bintree_t{
		}},
	}},
}}
