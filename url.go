package url

import (
	"github.com/golangplus/bytes"

	. "github.com/gohtml/elements"
	"github.com/gohtml/utils"
)

// Makes a URL.
// port is ignored if set to zero.
// fragment is optional. If specified as an empty string, a trailing '#' will be appended.
func U(host string, path string, q QUERY, fragment ...string) URL {
	return FullURL("", host, 0, path, q, fragment...)
}
func FullURL(scheme, host string, port int, path string, q QUERY, fragment ...string) URL {
	var b bytesp.Slice
	if len(host) > 0 {
		if len(scheme) > 0 {
			b.WriteString(scheme)
			b.WriteByte(':')
		}

		b.WriteString(`//`)
		b.WriteString(utils.EscapeHost(host))

		if port > 0 {
			b.WriteByte(':')
			b.WriteItoa(int64(port), 10)
		}

		if len(path) == 0 || path[0] != '/' {
			b.WriteByte('/')
		}
	}

	b.WriteString(path)

	if len(q) > 0 {
		b.WriteByte('?')
		b.WriteString(string(q))
	}

	if fragment != nil {
		b.WriteByte('#')
		b.WriteString(fragment[0])
	}

	return URL(b)
}

// Generates a URL QUERY.
// nameValue's should be paired otherwise the last one is ignored.
func Q(nameValue ...string) QUERY {
	var b bytesp.Slice

	for i := 1; i < len(nameValue); i += 2 {
		if i > 1 {
			b.WriteByte('&')
		}
		b.WriteString(utils.EscapeQuery(nameValue[i-1]))
		b.WriteByte('=')
		b.WriteString(utils.EscapeQuery(nameValue[i]))
	}
	return QUERY(b)
}
