package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFilePaths(p string) (string, string) {
	d, f := filepath.Split(filepath.Clean(p))
	return d, f
}

func FilePathJoin(d, f string) string {
	filepath.Join()
	p := filepath.Join(d, f)
	return p
}

func FileStat(p string) bool {
	_, err := os.Stat(p)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func FileRemove(p string) bool {
	if err := os.Remove(p); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
