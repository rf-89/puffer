package fileinfo

import (
	"sync"

	"github.com/teapod89/puffer/util/file"
)

func GetDirFiles(files []map[string]string) []map[string]string {
	var wg = &sync.WaitGroup{}

	var dirfiles []map[string]string
	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func(i int) {
			m := make(map[string]string)
			d, f := file.GetFilePaths(files[i]["filename"])
			m["directory"] = d
			m["filename"] = f
			m["hash"] = files[i]["hash"]
			dirfiles = append(dirfiles, m)
			wg.Done()
		}(i)
	}
	wg.Wait()

	return dirfiles
}
