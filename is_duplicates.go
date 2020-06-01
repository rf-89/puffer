package main

func isDuplicates(fileInfo []map[string]string) []map[string]string {
	var duplicates []map[string]string
	const hashKey = "hash"
	const fileKey = "filename"
	for _, ref := range fileInfo {
		for _, v := range fileInfo {
			if ref[fileKey] != v[fileKey] && ref[hashKey] == v[hashKey] {
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
