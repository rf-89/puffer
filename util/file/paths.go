package file

import (
	"path/filepath"
)

func GetFilePaths(p string) (string, string) {
	d, f := filepath.Split(filepath.Clean(p))
	return d, f
}
