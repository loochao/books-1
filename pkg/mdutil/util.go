package mdutil

import "strings"

const (
	// KVRecordSeparator is a (hopefully) unique string that separates records in Key/Value file
	KVRecordSeparator = "|======"
)

// https://stackoverflow.com/questions/695438/safe-characters-for-friendly-url
func charIsURLSafe(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= '0' && c <= '9' {
		return true
	}
	switch c {
	case '-', '.', '_', '~':
		return true
	}
	return false
}

func shortenConsequitve(s string, c string) string {
	s2 := c + c
	for strings.Contains(s, s2) {
		s = strings.Replace(s, s2, c, -1)
	}
	// strip c from the beginning
	for len(s) > 0 && s[0:1] == c {
		s = s[1:]
	}
	return s
}

// MakeURLSafe converts arbitrary string into a string that can be used as a file name or url
func MakeURLSafe(s string) string {
	n := len(s)
	d := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		c := s[i]
		if charIsURLSafe(c) {
			if c == '.' {
				c = '-'
			}
			d = append(d, c)
		} else {
			if c == ' ' {
				d = append(d, '-')
			}
		}
	}
	s = string(d)
	s = strings.ToLower(s)
	s = shortenConsequitve(s, "-")
	return s
}
