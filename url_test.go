package url

import (
	"fmt"
	"testing"

	. "github.com/gohtml/elements"
	"github.com/golangplus/testing/assert"
)

func ExampleQ() {
	fmt.Println(Q("q", "golang", "page", "1"))

	// OUTPUT:
	// q=golang&page=1
}

func TestQ(t *testing.T) {
	assert.Equal(t, "Q", Q("abc", "def"), QUERY("abc=def"))
	assert.Equal(t, "Q", Q("abc", "d=f"), QUERY("abc=d%3Df"))
	assert.Equal(t, "Q", Q("abc", "def", "123", "456"), QUERY("abc=def&123=456"))
}

func ExampleU() {
	fmt.Println(U("http", "www.example.com", 1234, "/main", Q("q", "\"hello world"), "abc"))
	fmt.Println(U("http", "www.example.com", 1234, "main", Q("q&b", "hello+world"), "abc"))
	fmt.Println(U("", "www.example.com", 1234, "main", Q("q&b", "hello+world"), "abc"))
	fmt.Println(U("", "", 0, "main", Q("q&b", "hello+world"), "abc"))
	fmt.Println(U("", "", 0, "", Q("q&b", "hello+world"), "abc"))
	fmt.Println(U("", "", 0, "", "", "abc"))
	fmt.Println(U("", "www.example.com", 0, "main", Q("q&b", "hello+world"), "abc"))
	fmt.Println(U("http", "www.example.com", 0, "", Q("?", "?"), ""))
	fmt.Println(U("http", "www.example.com", 0, "", Q("?", "?")))

	// OUTPUT:
	// http://www.example.com:1234/main?q=%22hello+world#abc
	// http://www.example.com:1234/main?q%26b=hello%2Bworld#abc
	// //www.example.com:1234/main?q%26b=hello%2Bworld#abc
	// main?q%26b=hello%2Bworld#abc
	// ?q%26b=hello%2Bworld#abc
	// #abc
	// //www.example.com/main?q%26b=hello%2Bworld#abc
	// http://www.example.com/?%3F=%3F#
	// http://www.example.com/?%3F=%3F
}

func TestU(t *testing.T) {
	assert.Equal(t, "U", U("", "www:abc.com", 0, "", ""), URL("//wwwabc.com/"))
}
