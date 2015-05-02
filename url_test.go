package url

import (
	"testing"

	. "github.com/gohtml/elements"
	"github.com/golangplus/testing/assert"
)

func TestQ(t *testing.T) {
	assert.Equal(t, "Q", Q("abc", "def"), QUERY("abc=def"))
	assert.Equal(t, "Q", Q("abc", "d=f"), QUERY("abc=d%3Df"))
	assert.Equal(t, "Q", Q("abc", "def", "123", "456"), QUERY("abc=def&123=456"))
}
