package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists_NoSite_ExpectFalse(t *testing.T) {
	e := Exists("")
	assert.False(t, e)
}

func TestExists_InexistentSite_ExpectFalse(t *testing.T) {
	e := Exists("NOT_EXISTENT_SITE")
	assert.False(t, e)
}
