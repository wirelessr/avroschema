package avroschema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNameAndOmit(t *testing.T) {
	var tdata = []struct {
		input    string
		name     string
		optional bool
		inline   bool
	}{
		{
			"abcd", "abcd", false, false,
		},
		{
			"abcd,omitempty", "abcd", true, false,
		},
		{
			"abcd,inline", "abcd", false, true,
		},
		{
			"abcd,inline,omitempty", "abcd", true, true,
		},
	}

	for _, tt := range tdata {
		t.Run(tt.input, func(t *testing.T) {
			tag := parseStructTag(tt.input)
			assert.Equal(t, tt.name, tag.Name)
			assert.Equal(t, tt.optional, tag.Optional)
			assert.Equal(t, tt.inline, tag.Inline)
		})
	}
}
