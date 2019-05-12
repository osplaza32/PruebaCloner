package Utils

import "strings"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}