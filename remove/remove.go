package remove

import (
	"fmt"
	"sync"

	"github.com/teapod89/puffer/util/file"
)

func DoRemoveFiles(files []map[string]string) []map[string]string {
	var wg = &sync.WaitGroup{}

	var removeDetails []map[string]string
	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func(i int) {
			m := make(map[string]string)

			m = files[i]
			f := m["duplicate_filename"]
			d := m["duplicate_directory"]
			//パス名を生成
			p := file.FilePathJoin(d, f)
			// パスの存在確認
			m["filestat"] = "success"
			m["fileremove"] = "success"
			if !file.FileStat(p) {
				m["filestat"] = "error"
			}
			if !file.FileRemove(p) {
				m["fileremove"] = "error"
			}
			removeDetails = append(removeDetails, m)
			wg.Done()
		}(i)
	}
	wg.Wait()

	return removeDetails
}
