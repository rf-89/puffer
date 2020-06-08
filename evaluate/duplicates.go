package evaluate

import (
	"sort"

	"github.com/teapod89/puffer/util/list"
)

func Duplicates(fileInfo []map[string]string) ([]map[string]string, int, int) {
	var duplicates []map[string]string
	const hashKey = "hash"
	const fileKey = "filename"
	var fnCount, dFnCount int

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
				fn := sorted[1]
				if len(fn) > fnCount {
					fnCount = len(fn)
				}
				dFn := sorted[0]
				if len(dFn) > dFnCount {
					dFnCount = len(dFn)
				}

				m := map[string]string{}
				m[fileKey] = fn
				m[hashKey] = v[hashKey]
				m["duplicate_filename"] = dFn
				m["duplicate_hash"] = ref[hashKey]
				duplicates = append(duplicates, m)
				continue
			}
		}
	}
	return duplicates, fnCount, dFnCount
}
