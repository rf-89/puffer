package list

//配列の中に特定の文字列が含まれるかを返す
func MapArrayContains(maps []map[string]string, key, str string) bool {
	for _, v := range maps {
		if v[key] == str {
			return true
		}
	}
	return false
}
