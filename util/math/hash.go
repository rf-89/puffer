package math

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

func Hash(f string) map[string]string {
	m := map[string]string{}
	m["filename"] = f

	fh, err := os.Open(f)
	if err != nil {
		//エラーの原因を詰めて返す。
		m["hash"] = err.Error()
		return m
	}
	defer fh.Close()

	h := sha256.New()
	if _, err := io.Copy(h, fh); err != nil {
		//エラーの原因を詰めて返す。
		m["hash"] = err.Error()
		return m
	}
	m["hash"] = bytes2str(h.Sum(nil))
	return m
}

func bytes2str(bytes ...[]byte) string {
	var s []string
	for _, b := range bytes {
		s = append(s, fmt.Sprintf("%02x", b))
	}
	return strings.Join(s, " ")
}
