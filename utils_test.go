package avroschema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNameAndOmit(t *testing.T) {
	s1 := "abcd"
	n1, opt1 := GetNameAndOmit(s1)
	assert.Equal(t, "abcd", n1)
	assert.False(t, opt1)

	s2 := "abcd,omitempty"
	n2, opt2 := GetNameAndOmit(s2)
	assert.Equal(t, "abcd", n2)
	assert.True(t, opt2)
}
