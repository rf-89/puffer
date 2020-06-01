package main

import (
	"os"
	"path/filepath"
)

func getFiles(in string) ([]string, error) {
	var files []string
	err := filepath.Walk(in, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, p)
		}
		return nil
	})
	return files, err
}
