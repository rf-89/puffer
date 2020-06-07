package evaluate

import (
	"sort"

	"github.com/teapod89/puffer/util/list"
)

func Duplicates(fileInfo []map[string]string) []map[string]string {
	var duplicates []map[string]string
	const hashKey = "hash"
	const fileKey = "filename"
	for _, ref := range fileInfo {
		for _, v := range fileInfo {
			if ref[fileKey] != v[fileKey] && ref[hashKey] == v[hashKey] {
				if list.MapArrayContains(duplicates, fileKey, v[fileKey]) ||
					list.MapArrayContains(duplicates, hashKey, v[hashKey]) {
					continue
				}

				//本来はfilename keyでソートすべきだが、ハッシュは同一で重複ファイルとみなすことができるためワークアラウンドとして本対応を行う。
				var sorted []string = []string{ref[fileKey], v[fileKey]}
				sort.Strings(sorted)

				m := map[string]string{}
				m[fileKey] = sorted[1]
				m[hashKey] = v[hashKey]
				m["duplicate_filename"] = sorted[0]
				m["duplicate_hash"] = ref[hashKey]
				duplicates = append(duplicates, m)
				continue
			}
		}
	}
	return duplicates
}
