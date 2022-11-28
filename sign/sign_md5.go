package sign

import "crypto/md5"

func Md5Sign(_, body string) []byte {
	m := md5.New()
	_, _ = m.Write([]byte(body))
	return m.Sum(nil)
}
