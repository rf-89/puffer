package evaluate

func Duplicates(fileInfo []map[string]string) []map[string]string {
	var duplicates []map[string]string
	const hashKey = "hash"
	const fileKey = "filename"
	for _, ref := range fileInfo {
		for _, v := range fileInfo {
			if ref[fileKey] != v[fileKey] && ref[hashKey] == v[hashKey] {
				if mapArrayContains(duplicates, fileKey, v[fileKey]) ||
					mapArrayContains(duplicates, hashKey, v[hashKey]) {
					continue
				}
				m := map[string]string{}
				m[fileKey] = v[fileKey]
				m[hashKey] = v[hashKey]
				m["duplicate_filename"] = ref[fileKey]
				m["duplicate_hash"] = ref[hashKey]
				duplicates = append(duplicates, m)
				continue
			}
		}
	}
	return duplicates
}

//配列の中に特定の文字列が含まれるかを返す
func mapArrayContains(maps []map[string]string, key, str string) bool {
	for _, v := range maps {
		if v[key] == str {
			return true
		}
	}
	return false
}
