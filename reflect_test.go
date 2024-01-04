package avroschema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimitiveType(t *testing.T) {
	type Entity struct {
		AStrField    string  `json:"a_str_field"`
		AIntField    int     `json:"a_int_field"`
		ABoolField   bool    `json:"a_bool_field"`
		AFloatField  float32 `json:"a_float_field"`
		ADoubleField float64 `json:"a_double_field"`
	}

	expected := `{
    "name": "Entity",
    "type": "record",
    "fields": [
      {"name": "a_str_field", "type": "string"},
      {"name": "a_int_field", "type": "int"},
      {"name": "a_bool_field", "type": "boolean"},
      {"name": "a_float_field", "type": "float"},
      {"name": "a_double_field", "type": "double"}
    ]
  }`

	e := Entity{}

	// test for instance
	r1, err1 := Reflect(e)
	assert.JSONEq(t, expected, r1)
	assert.Nil(t, err1)

	// test for pointer
	r2, err2 := Reflect(&e)
	assert.JSONEq(t, expected, r2)
	assert.Nil(t, err2)
}
